package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/handler"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/repository"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/service"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

// gin-swagger middleware
// swagger embed files

// @title BiblioMusic App API
// @version 1.0
// @description API сервер для приложения BiblioMusic

// @host localhost:8000
// @BasePath /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	log.Info("Loading environment variables...")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	log.Info("Creating DB connection...")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.WithField("err:", err.Error()).Error("Couldn't create DB connection!")
		return
	}
	log.Debug("DB connected successfully")

	log.Info("Creating repositories...")
	repos := repository.NewRepository(db)
	log.Debug("Repositories created successfully")

	log.Info("Creating services...")
	serv := service.NewService(repos)
	log.Debug("Services created successfully")

	log.Info("Creating handlers...")
	handl := handler.NewHandler(serv)
	log.Debug("Handlers created successfully")

	log.Info("Creating server...")
	server := new(domain.Server)
	log.Debug("Server created successfully")

	go func() {
		log.Info("Starting the HTTP server...")
		if err := server.Run(os.Getenv("APP_PORT"), handl.InitRoutes(), db); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("error occured while running http server: %s", err.Error())
			}
			log.Info("Server stopped running")
		}
	}()

	log.Info("BiblioMusic-app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("Shutting down the server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}
	log.Info("Server gracefully stopped")
}
