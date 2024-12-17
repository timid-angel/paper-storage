package domain

import (
	"paper-server/domain/dtos"
)

type IDomainError interface {
	error
}

type DomainError struct {
	message string
}

func (err *DomainError) Error() string {
	return err.message
}

type IPaperStorageUsecase interface {
	AddPaper(paper dtos.AddPaperInput) IDomainError
	ListPapers() (dtos.ListPaperOuput, IDomainError)
	GetPaperDetails(paperNumber int) (dtos.GetPaperDetailsOutput, IDomainError)
	FetchPaperContent(paperNumber int) (dtos.FetchPaperContentOutput, IDomainError)
}

type IStorageRepository interface {
	AddPaper(paper dtos.AddPaperInput) IDomainError
	ListPapers() (dtos.ListPaperOuput, IDomainError)
	GetPaperDetails(paperNumber int) (dtos.GetPaperDetailsOutput, IDomainError)
	FetchPaperContent(paperNumber int) (dtos.FetchPaperContentOutput, IDomainError)
}
