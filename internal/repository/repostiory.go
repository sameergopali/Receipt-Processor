package repository

import (
	"errors"

	"github.com/google/uuid"
)

type Repository struct {
	receipts map[string]int
}

func NewRepository() *Repository {
	return &Repository{
		receipts: make(map[string]int),
	}
}

func (repo *Repository) GetById(id string) (int, error) {
	points, exists := repo.receipts[id]
	if !exists {
		return -1, errors.New("receipt Id not found")
	}
	return points, nil
}

func (repo *Repository) AddEntry(points int) (string, error) {
	id := uuid.New().String()
	repo.receipts[id] = points
	return id, nil
}
