package handler

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ReceiptHandler struct {
	service *service.ReceiptService
}

func NewReceiptHandler(service *service.ReceiptService) *ReceiptHandler {
	return &ReceiptHandler{service: service}
}

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
	// Decode the JSON request body
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid", http.StatusBadRequest)
		return
	}
	id, _ := h.service.ProcessReceipt(receipt)
	response := models.ProcessReceiptResult{Id: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// @Summary Get points by receipt ID
// @Description Get the points associated with a specific receipt ID
// @Tags receipts
// @Param id path string true "Receipt ID"
// @Produce json
// @Success 200 {object} models.GetPointsResult
// @Failure 404 {string} string "Receipt not found"
// @Router /receipts/{id}/points [get]
func (h *ReceiptHandler) GetPointsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, err := h.service.GetPointsById(id)
	if err != nil {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := models.GetPointsResult{Points: points}
	json.NewEncoder(w).Encode(response)
}
