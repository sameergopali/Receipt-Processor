package main

import (
	"cmd/main.go/internal/handler"
	"cmd/main.go/internal/repository"
	"cmd/main.go/internal/service"
	"log"
	"net/http"

	_ "cmd/main.go/docs"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample Swagger server.
// @host localhost:8080
// @BasePath /

func main() {
	loadConfig()

	receiptRepo := repository.NewRepository()
	receiptService := service.NewReceiptService(receiptRepo)
	receiptHandler := handler.NewReceiptHandler(receiptService)

	router := mux.NewRouter()
	setupRoutes(router, receiptHandler)

	addr := viper.GetString("server.address") + ":" + viper.GetString("server.port")
	log.Println("Server started on", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func setupRoutes(router *mux.Router, handler *handler.ReceiptHandler) {
	router.HandleFunc("/receipts/process", handler.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handler.GetPointsById).Methods("GET")
	// Swagger documentation route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func loadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
}
