package controller

import (
	"fmt"
	"paper-server/domain/dtos"
	"paper-server/server/domain"
)

type PaperStorage struct {
	usecase domain.IPaperStorageUsecase
}

func NewPaperStorage(usecase domain.IPaperStorageUsecase) *PaperStorage {
	return &PaperStorage{
		usecase: usecase,
	}
}

func (controller *PaperStorage) AddPaper(args dtos.AddPaperInput, reply *dtos.AddPaperOutput) error {
	err := controller.usecase.AddPaper(&args)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

func (controller *PaperStorage) ListPapers(args dtos.ListPaperInput, reply *dtos.ListPaperOuput) error {
	paperList, err := controller.usecase.ListPapers()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	reply.Papers = *paperList
	return nil
}

func (controller *PaperStorage) GetPaperDetails(args dtos.GetPaperDetailsInput, reply *dtos.GetPaperDetailsOutput) error {
	paperDetails, err := controller.usecase.GetPaperDetails(args.PaperNumber)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	*reply = dtos.GetPaperDetailsOutput{
		PaperData: paperDetails,
	}

	return nil
}

func (controller *PaperStorage) FetchPaperContent(args dtos.FetchPaperContentInput, reply *dtos.FetchPaperContentOutput) error {
	paper, err := controller.usecase.FetchPaperContent(args.PaperNumber)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	*reply = dtos.FetchPaperContentOutput{
		Paper: paper,
	}

	return nil
}
