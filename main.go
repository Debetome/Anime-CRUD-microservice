package main

import (
	"microservice/handlers"
	"microservice/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"os"
	"os/signal"
	"context"
	"log"
)

const CONNECTION_STRING = "mongodb://pepeluis:12345678@localhost:27017/?authSource=People"

func main() {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTION_STRING))
	if err != nil {
		logger.Fatal(err)
	}

	defer client.Disconnect(ctx)
	db := database.NewDatabase(client, logger)
	handler := handlers.NewAnimes(db, logger)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getOneRouter := sm.Methods(http.MethodGet).Subrouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	updateRouter := sm.Methods(http.MethodPut).Subrouter()
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	
	getRouter.HandleFunc("/", handler.GetAnimes)
	getOneRouter.HandleFunc("/{id}", handler.GetAnime)
	postRouter.HandleFunc("/", handler.PostAnime)
	updateRouter.HandleFunc("/{id}", handler.UpdateAnime)
	deleteRouter.HandleFunc("/{id}", handler.DeleteAnime)

	//postRouter.Use(handler.MiddlewareValidateAnime)
	//updateRouter.Use(handler.MiddlewareValidateAnime)

	server := &http.Server{
		Addr: ":8000",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port 8000")

		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received", sig, "signal")

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}