package memory

import (
	"sync"

	"ml-feature-store-api/models"
)

type InMemoryStore struct {
	Data map[string]models.Feature
	mu   sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		Data: make(map[string]models.Feature),
	}
}

func (s *InMemoryStore) Set(key string, feature models.Feature) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Data[key] = feature
}

func (s *InMemoryStore) Get(key string) (models.Feature, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.Data[key]
	return val, ok
}

func (s *InMemoryStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.Data, key)
}
