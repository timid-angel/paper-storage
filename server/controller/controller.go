package controller

import (
	"fmt"
	"log"
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
	log.Printf("Called AddPaper with args: %+v", args)
	paperNumber, err := controller.usecase.AddPaper(&args)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	reply.PaperNumber = paperNumber
	return nil
}

func (controller *PaperStorage) ListPapers(args dtos.ListPaperInput, reply *dtos.ListPaperOuput) error {
	log.Printf("Called ListPapers with args: %+v", args)
	paperList, err := controller.usecase.ListPapers()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	reply.Papers = *paperList
	return nil
}

func (controller *PaperStorage) GetPaperDetails(args dtos.GetPaperDetailsInput, reply *dtos.GetPaperDetailsOutput) error {
	log.Printf("Called GetPaperDetails with args: %+v", args)
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
	log.Printf("Called FetchPaperContent with args: %+v", args)
	paper, err := controller.usecase.FetchPaperContent(args.PaperNumber)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	*reply = dtos.FetchPaperContentOutput{
		Paper: paper,
	}

	return nil
}
