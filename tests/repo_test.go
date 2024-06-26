package tests

import (
	"cmd/main.go/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetById(t *testing.T) {
	// Initialize a new mock repository
	mockRepo := repository.NewRepository()
	generateId, _ := mockRepo.AddEntry(12)

	tests := []struct {
		name     string
		id       string
		expected int
	}{
		{
			name:     "RepoGetById Test 1",
			id:       generateId,
			expected: 12,
		},
		{
			name:     "RepoGetById Test 2",
			id:       "adfasdaf",
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			point, _ := mockRepo.GetById(test.id)
			assert.Equal(t, test.expected, point, "Failed tested")
		})
	}

}

func TestRepository_Create(t *testing.T) {
	mockRepo := repository.NewRepository()
	id1, err1 := mockRepo.AddEntry(10)
	id2, err2 := mockRepo.AddEntry(20)
	t.Run("RepoAddEntry Test", func(t *testing.T) {
		assert.NoError(t, err1, "Expected no error for valid input")
		assert.NotEmpty(t, id1, "Expected a non-empty ID")
		assert.NoError(t, err2, "Expected no error for valid input")
		assert.NotEmpty(t, id2, "Expected a non-empty ID")
	})

	t.Run("RepoAddEntry Unique Id test", func(t *testing.T) {
		assert.NotEqual(t, id1, id2, "Expected different IDs for different creations")
	})

}
