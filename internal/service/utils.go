package service

import (
	"cmd/main.go/internal/models"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

// Rule defines the interface for a points calculation rule
type Rule interface {
	Calculate(receipt models.Receipt) int
}

// Rule1 implements Rule for alphanumeric characters in the retailer name
type Rule1 struct{}

func (r Rule1) Calculate(receipt models.Receipt) int {
	points := 0
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}
	return points
}

// Rule2 implements Rule for round dollar amount total
type Rule2 struct{}

func (r Rule2) Calculate(receipt models.Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("rule2: error parsing total: %v\n", err)
		return 0
	}
	if total == float64(int(total)) {
		return 50
	}
	return 0
}

// Rule3 implements Rule for total being a multiple of 0.25
type Rule3 struct{}

func (r Rule3) Calculate(receipt models.Receipt) int {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("rule3: error parsing total: %v\n", err)
		return 0
	}
	if math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

// Rule4 implements Rule for every two items on the receipt
type Rule4 struct{}

func (r Rule4) Calculate(receipt models.Receipt) int {
	return (len(receipt.Items) / 2) * 5
}

// Rule5 implements Rule for item descriptions
type Rule5 struct{}

func (r Rule5) Calculate(receipt models.Receipt) int {
	points := 0
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Printf("rule5: error parsing item price: %v\n", err)
				continue
			}
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

// Rule6 implements Rule for odd day in the purchase date
type Rule6 struct{}

func (r Rule6) Calculate(receipt models.Receipt) int {
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		log.Printf("rule6: error parsing purchase date: %v\n", err)
		return 0
	}
	if date.Day()%2 != 0 {
		return 6
	}
	return 0
}

// Rule7 implements Rule for purchase time between 2:00pm and 4:00pm
type Rule7 struct{}

func (r Rule7) Calculate(receipt models.Receipt) int {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		log.Printf("rule7: error parsing purchase time: %v\n", err)
		return 0
	}
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		return 10
	}
	return 0
}
