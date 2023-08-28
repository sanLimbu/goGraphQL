package gopher

import (
	"errors"
	"sync"
)

// InMemoryRepository is a storage for gophers that uses a map to store them

type InMemoryRepository struct {
	//gophers is our super storage for gophers
	gophers []Gopher
	sync.Mutex
}

// NewMemoryRepository initializes a memory with mock data
func NewMemoryRepository() *InMemoryRepository {
	gophers := []Gopher{
		{
			ID:         "1",
			Name:       "Santosh",
			Hired:      true,
			Profession: "Software professional",
		},
		{
			ID:         "2",
			Name:       "Marvin",
			Hired:      true,
			Profession: "Son",
		},
	}
	return &InMemoryRepository{
		gophers: gophers,
	}
}

// GetGophers returns all gophers
func (im *InMemoryRepository) GetGophers() ([]Gopher, error) {
	return im.gophers, nil
}

// GetGopher will return a gopher by its ID
func (im *InMemoryRepository) GetGopher(id string) (Gopher, error) {

	for _, gopher := range im.gophers {
		if gopher.ID == id {
			return gopher, nil
		}
	}
	return Gopher{}, errors.New("no such gopher exists")
}
