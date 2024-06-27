package service

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/repository"
	"log"
)

// ReceiptService provides services related to processing receipts.
type ReceiptService struct {
	repo repository.Repository
}

// NewReceiptService creates a new instance of ReceiptService.
func NewReceiptService(repo repository.Repository) *ReceiptService {
	return &ReceiptService{repo: repo}
}

// ProcessReceipt processes a receipt and returns the receipt ID.
func (s *ReceiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	log.Println("ReceiptService::ProcessReceipt: started processing receipt")
	points := s.CalculatePoints(receipt)
	id, err := s.repo.AddEntry(points)
	if err != nil {
		log.Printf("ReceiptService::ProcessReceipt: error adding receipt entry: %v\n", err)
		return "", err
	}
	log.Printf("ReceiptService::ProcessReceipt: successfully processed receipt with ID: %s\n", id)
	return id, err

}

// GetPointsById retrieves the points for a receipt by its ID.
func (s *ReceiptService) GetPointsById(id string) (int, error) {
	log.Printf("ReceiptService::GetPointsById: retrieving points for receipt ID %s\n", id)
	points, err := s.repo.GetById(id)
	if err != nil {
		log.Printf("ReceiptService::GetPointsById: error retrieving points for receipt ID %s: %v\n", id, err)
		return -1, err
	}

	log.Printf("ReceiptService::GetPointsById: successfully retrieved points for receipt ID %s\n", id)
	return points, nil
}

// CalculatePoints calculates the points for a given receipt.
func (s *ReceiptService) CalculatePoints(receipt models.Receipt) int {
	log.Println("ReceiptService::CalculatePoints: calculating points for receipt")
	points := 0
	rules := []Rule{
		Rule1{},
		Rule2{},
		Rule3{},
		Rule4{},
		Rule5{},
		Rule6{},
		Rule7{},
	}

	for _, rule := range rules {
		points += rule.Calculate(receipt)
	}
	return points
}
