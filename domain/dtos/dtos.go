package dtos

import "paper-server/domain/entities"

type Response struct {
	Success bool
	Message string
}

type AddPaperInput struct {
	*entities.Paper
}

type AddPaperOutput struct {
	Response
}

type ListPaperInput struct{}

type ListPaperOuput struct {
	Papers *[]entities.PaperData
	Response
}

type GetPaperDetailsInput struct {
	PaperNumber int
}

type GetPaperDetailsOutput struct {
	*entities.PaperData
	Response
}

type FetchPaperContentInput struct {
	PaperNumber int
}

type FetchPaperContentOutput struct {
	*entities.Paper
	Response
}
