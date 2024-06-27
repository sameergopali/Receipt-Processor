package repository

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

// Repository Interface
type Repository interface {
	GetById(id string) (int, error)
	AddEntry(points int) (string, error)
}

// InMemory Repo is the data repository for storing receipt points.

type MemRepository struct {
	receipts map[string]int
}

// NewRepository creates a new instance of Repository.
func NewMemRepository() MemRepository {
	return MemRepository{
		receipts: make(map[string]int),
	}
}

// GetById retrieves points associated with a receipt ID.
func (repo MemRepository) GetById(id string) (int, error) {
	points, exists := repo.receipts[id]
	if !exists {
		log.Printf("Repository::GetById: receipt ID %s not found\n", id)
		return -1, errors.New("receipt Id not found")
	}
	log.Printf("Repository::GetById: successfully retrieved points for receipt ID %s\n", id)
	return points, nil
}

// AddEntry adds a new receipt entry with points and returns the receipt ID.
func (repo MemRepository) AddEntry(points int) (string, error) {
	log.Println("Repository::AddEntry: adding a new receipt entry")
	id := uuid.New().String()
	repo.receipts[id] = points
	log.Printf("Repository::AddEntry: successfully added receipt entry with ID %s\n", id)
	return id, nil
}
