package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	// "github.com/indenigrate/List_Jobs/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
}

func main() {
	//INITIALISATION
	//loading .env (environment variables)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//retrieving PORT variable data
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	//initiating router
	router := chi.NewRouter()

	// initiating cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//initialise storage

	// fmt.Printf("%+v\n", store)
	apiCfg := apiConfig{}

	//handle requests
	router.Get("/healthz", handlerReadiness)
	router.Get("/{input}", apiCfg.handlerRequest)
	router.Post("/chat", apiCfg.handlerChatRequest)
	//initiate server properties
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v\n", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server %v\n", err)
	}
}
