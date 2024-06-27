package handler

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ReceiptHandler struct {
	service *service.ReceiptService
}

func NewReceiptHandler(service *service.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{service: service}
}

// ProcessReceipt processes a receipt and returns the receipt ID.
// @Summary Process a receipt
// @Description Process the receipt and return the ID
// @Tags receipts
// @Accept json
// @Produce json
// @Param receipt body models.Receipt true "Receipt"
// @Success 200 {object} models.ProcessReceiptResult
// @Failure 400 {string} string "The receipt is invalid"
// @Router /receipts/process [post]
func (h *ReceiptHandler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	log.Println("ReceiptHandler::ProcessReceipt: started processing receipt")
	// Decode the JSON request body
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		log.Printf("ReceiptHandler::ProcessReceipt: error decoding receipt: %v\n", err)
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}

	// Process the receipt using the service
	id, err := h.service.ProcessReceipt(receipt)
	if err != nil {
		log.Printf("ReceiptHandler::ProcessReceipt: error processing receipt: %v\n", err)
		http.Error(w, "Failed to process receipt", http.StatusInternalServerError)
		return
	}

	// Create Response
	response := models.ProcessReceiptResult{Id: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// GetPointsById retrieves the points for a receipt by its ID.
// @Summary Get points by receipt ID
// @Description Get the points associated with a specific receipt ID
// @Tags receipts
// @Param id path string true "Receipt ID"
// @Produce json
// @Success 200 {object} models.GetPointsResult
// @Failure 404 {string} string "Receipt not found"
// @Router /receipts/{id}/points [get]
func (h *ReceiptHandler) GetPointsById(w http.ResponseWriter, r *http.Request) {
	log.Println("ReceiptHandler::GetPointsById: started retrieving points")
	// Extract the receipt ID from the URL path
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the points using the service
	points, err := h.service.GetPointsById(id)
	if err != nil {
		log.Printf("ReceiptHandler::GetPointsById: error retrieving points for receipt ID %s: %v\n", id, err)
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Create the response
	response := models.GetPointsResult{Points: points}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("ReceiptHandler::GetPointsById: error encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	log.Printf("ReceiptHandler::GetPointsById: successfully retrieved points for receipt ID: %s\n", id)
}
