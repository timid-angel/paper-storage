package usecase

import (
	"paper-server/domain"
	"paper-server/domain/dtos"
)

type PaperStorageUsecase struct{}

func (usecase *PaperStorageUsecase) AddPaper(paper dtos.AddPaperInput) domain.IDomainError {
}

func (usecase *PaperStorageUsecase) ListPapers() (dtos.ListPaperOuput, domain.IDomainError) {
}

func (usecase *PaperStorageUsecase) GetPaperDetails(paperNumber int) (dtos.GetPaperDetailsOutput, domain.IDomainError) {
}

func (usecase *PaperStorageUsecase) FetchPaperContent(paperNumber int) (dtos.FetchPaperContentOutput, domain.IDomainError) {
}
