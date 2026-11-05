package main

import (
	repository_postgres_rep "Threat-intelligence-service/internal/features/Threat/repository/postgres"
	threatservice "Threat-intelligence-service/internal/features/Threat/service/ThreatService"
	transport_threat_handlers "Threat-intelligence-service/internal/features/Threat/transport/http/handlers"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Threat-service start server")

	err := godotenv.Load()
	if err != nil {
		log.Println("Предупреждение: файл .env не найден")
	}

	ctx := context.Background()

	router := mux.NewRouter()

	conn, err := repository_postgres_rep.CreateConnection(ctx)
	if err != nil {
		fmt.Println("Database connection error")
		panic(err)
	}

	ThreatRep := repository_postgres_rep.NewThreatRepository(conn)
	ThreatService := threatservice.NewThreatService(ThreatRep)
	ThreatHandler := transport_threat_handlers.NewThreatHandler(ThreatService)

	router.Path("/info").Methods("POST").HandlerFunc(ThreatHandler.CreateThreate)

	if err := http.ListenAndServe("localhost:5050", router); err != nil {
		panic(fmt.Errorf("Failed to start the server: %w", err))
	}

}
