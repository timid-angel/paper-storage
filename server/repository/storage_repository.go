package storage_repository

import (
	"paper-server/domain"
	"paper-server/domain/dtos"
	"paper-server/domain/entities"
	"sync"
)

type StorageRepository struct {
	mu    sync.Mutex
	paper map[int]entities.Paper
}

func (r *StorageRepository) AddPaper(data *dtos.AddPaperInput) domain.IDomainError {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.paper[data.PaperNumber] = data.Paper
	return nil
}

func (r *StorageRepository) ListPapers() (*dtos.ListPaperOuput, domain.IDomainError) {
	papers := []entities.PaperData{}
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, paper := range r.paper {
		papers = append(papers, paper.PaperData)
	}

	return &dtos.ListPaperOuput{
		Papers: papers,
	}, nil
}

func (r *StorageRepository) GetPaperDetails(paperNumber int) (*dtos.GetPaperDetailsOutput, domain.IDomainError) {
	r.mu.Lock()
	defer r.mu.Unlock()
	paper, exists := r.paper[paperNumber]
	if !exists {
		return nil, domain.NewDomainError("Paper with given number does not exist")
	}

	return &dtos.GetPaperDetailsOutput{
		PaperData: paper.PaperData,
	}, nil
}

func (r *StorageRepository) FetchPaperContent(paperNumber int) (*dtos.FetchPaperContentOutput, domain.IDomainError) {
	r.mu.Lock()
	defer r.mu.Unlock()
	paper, exists := r.paper[paperNumber]
	if !exists {
		return nil, domain.NewDomainError("Paper with given number does not exist")
	}

	return &dtos.FetchPaperContentOutput{
		Paper: paper,
	}, nil
}
