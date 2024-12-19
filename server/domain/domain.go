package domain

import (
	"paper-server/domain/dtos"
	"paper-server/domain/entities"
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

func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}

type IPaperStorageUsecase interface {
	AddPaper(paper *dtos.AddPaperInput) (int, IDomainError)
	ListPapers() (*[]entities.PaperData, IDomainError)
	GetPaperDetails(paperNumber int) (*entities.PaperData, IDomainError)
	FetchPaperContent(paperNumber int) (*entities.Paper, IDomainError)
}

type IStorageRepository interface {
	AddPaper(data *entities.Paper) IDomainError
	ListPapers() (*[]entities.PaperData, IDomainError)
	GetPaperDetails(paperNumber int) (*entities.PaperData, IDomainError)
	FetchPaperContent(paperNumber int) (*entities.Paper, IDomainError)
	GetNewPaperNumber() int
}
