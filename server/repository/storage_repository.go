package storage_repository

import (
	"paper-server/domain/entities"
	"paper-server/server/domain"
	"sync"
)

type StorageRepository struct {
	mu              sync.Mutex
	paper           map[int]entities.Paper
	lastPaperNumber int
}

func NewStorageRepository() *StorageRepository {
	return &StorageRepository{
		paper:           make(map[int]entities.Paper),
		lastPaperNumber: 0,
	}
}

func (r *StorageRepository) AddPaper(data *entities.Paper) domain.IDomainError {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.paper[data.PaperNumber] = *data
	return nil
}

func (r *StorageRepository) ListPapers() (*[]entities.PaperData, domain.IDomainError) {
	papers := []entities.PaperData{}
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, paper := range r.paper {
		papers = append(papers, paper.PaperData)
	}

	return &papers, nil
}

func (r *StorageRepository) GetPaperDetails(paperNumber int) (*entities.PaperData, domain.IDomainError) {
	r.mu.Lock()
	defer r.mu.Unlock()
	paper, exists := r.paper[paperNumber]
	if !exists {
		return nil, domain.NewDomainError("Paper with given number does not exist")
	}

	return &paper.PaperData, nil
}

func (r *StorageRepository) FetchPaperContent(paperNumber int) (*entities.Paper, domain.IDomainError) {
	r.mu.Lock()
	defer r.mu.Unlock()
	paper, exists := r.paper[paperNumber]
	if !exists {
		return nil, domain.NewDomainError("Paper with given number does not exist")
	}

	return &paper, nil
}

func (r *StorageRepository) GetNewPaperNumber() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lastPaperNumber++
	return r.lastPaperNumber
}
