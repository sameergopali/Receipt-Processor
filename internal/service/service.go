package service

import (
	"cmd/main.go/internal/models"
	"cmd/main.go/internal/repository"
	"math"
	"strconv"
	"strings"
	"time"
)

type ReceiptService struct {
	repo *repository.Repository
}

func NewReceiptService(repo *repository.Repository) *ReceiptService {
	return &ReceiptService{repo: repo}
}

func (s *ReceiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	points := s.CalcuatePoints(receipt)
	id, err := s.repo.AddEntry(points)
	return id, err

}

func (s *ReceiptService) GetPointsById(id string) (int, error) {
	return s.repo.GetById(id)
}

func (s *ReceiptService) CalcuatePoints(receipt models.Receipt) int {
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
	// 6 points if the day in the purchase date is odd.
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	time, _ := time.Parse("15:04", receipt.PurchaseTime)
	if time.Hour() == 14 {
		points += 10
	}

	return points
}
