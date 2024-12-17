package controller

import (
	"paper-server/domain"
	"paper-server/domain/dtos"
)

type PaperStorage struct {
	usecase domain.IPaperStorageUsecase
}

func (controller *PaperStorage) AddPaper(args dtos.AddPaperInput, reply *dtos.AddPaperOutput) {

}

func (controller *PaperStorage) ListPapers(args dtos.ListPaperInput, reply *dtos.ListPaperOuput) {

}

func (controller *PaperStorage) GetPaperDetails(args dtos.GetPaperDetailsInput, reply *dtos.GetPaperDetailsOutput) {

}

func (controller *PaperStorage) FetchPaperContent(args dtos.FetchPaperContentInput, reply *dtos.FetchPaperContentOutput) {

}
