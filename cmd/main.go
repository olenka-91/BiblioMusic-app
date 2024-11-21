package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("BD_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.WithField("err:", err.Error()).Error("Couldn't create DB connection!")
	}

	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	handl := handler.NewHandler(serv)

	server := new(domain.Server)

	go func() {
		if err := server.Run(os.Getenv("APP_PORT"), handl.InitRoutes(), db); err != nil {
			log.Fatalf("error occured while running http server %s", err.Error())
		}
	}()

	log.Println("BiblioMusic-app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}
}
