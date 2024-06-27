package service

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/repository"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
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
	// Todo: Refactor CalculatePoints by abstracting rules.
	log.Println("ReceiptService::CalculatePoints: calculating points for receipt")
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("CalculatePoints: error parsing purchase date: %v\n", err)
	} else {
		if date.Day()%2 != 0 {
			points += 6
		}
	}

	//Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	time, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		log.Printf("CalculatePoints: error parsing purchase time: %v\n", err)
	} else {
		if time.Hour() >= 14 && time.Hour() < 16 {
			points += 10
		}
	}

	return points
}
