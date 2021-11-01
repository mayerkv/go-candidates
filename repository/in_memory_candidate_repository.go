package repository

import (
	"github.com/mayerkv/go-candidates/domain"
	"sync"
)

type InMemoryCandidateRepository struct {
	sync.Mutex
	items map[string]domain.Candidate
}

func NewInMemoryCandidateRepository() domain.CandidateRepository {
	return &InMemoryCandidateRepository{
		items: map[string]domain.Candidate{},
	}
}

func (r *InMemoryCandidateRepository) Save(candidate *domain.Candidate) error {
	r.Lock()
	defer r.Unlock()

	r.items[candidate.Id] = *candidate

	return nil
}

func (r *InMemoryCandidateRepository) FindById(id string) (*domain.Candidate, error) {
	r.Lock()
	defer r.Unlock()

	if candidate, ok := r.items[id]; ok {
		return &candidate, nil
	}

	return nil, nil
}

func (r *InMemoryCandidateRepository) FindAll(pageable domain.Pageable) (domain.CandidatePage, error) {
	r.Lock()
	defer r.Unlock()

	var candidates []domain.Candidate
	for _, c := range r.items {
		candidates = append(candidates, c)
	}

	return domain.CandidatePage{
		Items:      r.skip(candidates, pageable),
		TotalCount: len(candidates),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}, nil
}

func (r *InMemoryCandidateRepository) skip(candidates []domain.Candidate, pageable domain.Pageable) []domain.Candidate {
	start := (pageable.Page - 1) * pageable.Size
	if start > len(candidates)-1 {
		return []domain.Candidate{}
	}

	end := start + pageable.Size
	if end > len(candidates)-1 {
		return candidates[start:]
	}

	return candidates[start:end]
}
