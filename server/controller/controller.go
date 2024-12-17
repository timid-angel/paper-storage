package controller

import (
	"paper-server/domain"
	"paper-server/domain/dtos"
)

type PaperStorage struct {
	usecase domain.IPaperStorageUsecase
}

func NewPaperStorage(usecase domain.IPaperStorageUsecase) *PaperStorage {
	return &PaperStorage{
		usecase: usecase,
	}
}

func (controller *PaperStorage) AddPaper(args dtos.AddPaperInput, reply *dtos.AddPaperOutput) {
	err := controller.usecase.AddPaper(&args)
	if err != nil {
		*reply = dtos.AddPaperOutput{
			Response: dtos.Response{
				Success: false,
				Message: err.Error(),
			},
		}

		return
	}

	*reply = dtos.AddPaperOutput{
		Response: dtos.Response{
			Success: true,
			Message: "Paper added successfully",
		},
	}
}

func (controller *PaperStorage) ListPapers(args dtos.ListPaperInput, reply *dtos.ListPaperOuput) {
	paperList, err := controller.usecase.ListPapers()
	if err != nil {
		*reply = dtos.ListPaperOuput{
			Response: dtos.Response{
				Success: false,
				Message: err.Error(),
			},
		}

		return
	}

	*reply = dtos.ListPaperOuput{
		Response: dtos.Response{
			Success: true,
			Message: "Paper list fetched successfully",
		},
		Papers: paperList,
	}
}

func (controller *PaperStorage) GetPaperDetails(args dtos.GetPaperDetailsInput, reply *dtos.GetPaperDetailsOutput) {
	paperDetails, err := controller.usecase.GetPaperDetails(args.PaperNumber)
	if err != nil {
		*reply = dtos.GetPaperDetailsOutput{
			Response: dtos.Response{
				Success: false,
				Message: err.Error(),
			},
		}

		return
	}

	*reply = dtos.GetPaperDetailsOutput{
		Response: dtos.Response{
			Success: true,
			Message: "Paper details fetched successfully",
		},
		PaperData: paperDetails,
	}
}

func (controller *PaperStorage) FetchPaperContent(args dtos.FetchPaperContentInput, reply *dtos.FetchPaperContentOutput) {
	paper, err := controller.usecase.FetchPaperContent(args.PaperNumber)
	if err != nil {
		*reply = dtos.FetchPaperContentOutput{
			Response: dtos.Response{
				Success: false,
				Message: err.Error(),
			},
		}

		return
	}

	*reply = dtos.FetchPaperContentOutput{
		Response: dtos.Response{
			Success: true,
			Message: "Paper data fetched successfully",
		},
		Paper: paper,
	}
}
