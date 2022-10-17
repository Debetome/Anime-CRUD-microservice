package main

import (
	"microservice/handlers"
	"microservice/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
	"os"
	"os/signal"
	"errors"
	"context"
	"log"
)

func parse_connection_string() (connection string, err error) {
	const BASE_STRING = "mongodb://%s:%s@%s:%s/?authSource=%s"

	user, ok_user := os.LookupEnv("MONGO_USER")
	passw, ok_passw := os.LookupEnv("MONGO_PASSWORD")
	host, ok_host := os.LookupEnv("MONGO_HOST")
	port, ok_port := os.LookupEnv("MONGO_PORT")
	dbname, ok_dbname := os.LookupEnv("MONGO_DBNAME")

	if (!ok_port) { port = "27017" }
	if (!ok_host) {
		err = errors.New("'MONGO_HOST' environment variable not declared ...")
		return
	}

	if (!ok_user || !ok_passw) {
		err = errors.New("'MONGO_USER' or 'MONGO_PASSWORD' environment variable not declared ...")
		return
	}

	if (!ok_dbname) { dbname = "AnimeDB" }

	connection = fmt.Sprintf(BASE_STRING, user, passw, host, port, dbname)
	return
}


func main() {
	logger := log.New(os.Stdout, "anime-api ", log.LstdFlags)

	connection_string, err := parse_connection_string()
	if err != nil {
		logger.Println(err)
		os.Exit(-1)
	}

	logger.Printf("Mongo connection: %s", connection_string)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection_string))
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
	getRouter.HandleFunc("/get-animes", handler.GetAnimes)
	getOneRouter.HandleFunc("/get-anime/{id}", handler.GetAnime)
	postRouter.HandleFunc("/new-anime", handler.PostAnime)
	updateRouter.HandleFunc("/update-anime/{id}", handler.UpdateAnime)
	deleteRouter.HandleFunc("/delete-anime/{id}", handler.DeleteAnime)

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