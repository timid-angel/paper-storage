package paper_usecase

import (
	"paper-server/domain/dtos"
	"paper-server/domain/entities"
	"paper-server/server/domain"
)

type PaperStorageUsecase struct {
	storageRepository domain.IStorageRepository
}

func NewPaperStorageUsecase(storageRepository domain.IStorageRepository) *PaperStorageUsecase {
	return &PaperStorageUsecase{
		storageRepository: storageRepository,
	}
}

func (usecase *PaperStorageUsecase) AddPaper(paper *dtos.AddPaperInput) domain.IDomainError {
	return usecase.storageRepository.AddPaper(paper.Paper)
}

func (usecase *PaperStorageUsecase) ListPapers() (*[]entities.PaperData, domain.IDomainError) {
	return usecase.storageRepository.ListPapers()
}

func (usecase *PaperStorageUsecase) GetPaperDetails(paperNumber int) (*entities.PaperData, domain.IDomainError) {
	return usecase.storageRepository.GetPaperDetails(paperNumber)
}

func (usecase *PaperStorageUsecase) FetchPaperContent(paperNumber int) (*entities.Paper, domain.IDomainError) {
	return usecase.storageRepository.FetchPaperContent(paperNumber)
}
