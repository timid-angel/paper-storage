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

func (usecase *PaperStorageUsecase) AddPaper(data *dtos.AddPaperInput) (int, domain.IDomainError) {
	data.Paper.PaperNumber = usecase.storageRepository.GetNewPaperNumber()
	return data.Paper.PaperNumber, usecase.storageRepository.AddPaper(data.Paper)
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
