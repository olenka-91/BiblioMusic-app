package domain

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импортируем драйвер для PostgreSQL (или другого)
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler, db *sqlx.DB) error {
	if err := s.runMigrations(db); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) runMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}

	// Создаем экземпляр миграции
	m, err := migrate.NewWithDatabaseInstance(
		"file://./schema", // Путь к миграциям на диске
		"postgres",        // Тип базы данных (например, postgres)
		driver,            // Драйвер базы данных
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Выполняем миграции
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Миграции не требуются: нет новых изменений.")
		} else {
			log.Fatalf("Ошибка при применении миграций: %v", err)
		}
	} else {
		log.Println("Миграции успешно применены.")
	}

	return nil
}
