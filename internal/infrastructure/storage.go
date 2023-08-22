package infrastructure

import (
	"awesomeProject1/internal/domain"
	"context"
	"sync"
)

type InMemoryRepository struct {
	responses map[string]domain.ProxyResponse
	mu        sync.Mutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		responses: make(map[string]domain.ProxyResponse),
	}
}

func (r *InMemoryRepository) SaveResponse(ctx context.Context, response domain.ProxyResponse) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.responses[response.ID] = response
	return nil
}

func (r *InMemoryRepository) GetResponse(ctx context.Context, id string) (domain.ProxyResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	response, exists := r.responses[id]
	if !exists {
		return domain.ProxyResponse{}, domain.ErrResponseNotFound
	}
	return response, nil
}
