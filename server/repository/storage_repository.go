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

func (repository *StorageRepository) AddPaper(paper dtos.AddPaperInput) domain.IDomainError {

}

func (repository *StorageRepository) ListPapers() (dtos.ListPaperOuput, domain.IDomainError) {

}

func (repository *StorageRepository) GetPaperDetails(paperNumber int) (dtos.GetPaperDetailsOutput, domain.IDomainError) {
}

func (repository *StorageRepository) FetchPaperContent(paperNumber int) (dtos.FetchPaperContentOutput, domain.IDomainError) {
}
