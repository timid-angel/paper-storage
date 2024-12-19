package dtos

import "paper-server/domain/entities"

type AddPaperInput struct {
	*entities.Paper
}

type AddPaperOutput struct {
}

type ListPaperInput struct{}

type ListPaperOuput struct {
	Papers []entities.PaperData
}

type GetPaperDetailsInput struct {
	PaperNumber int
}

type GetPaperDetailsOutput struct {
	*entities.PaperData
}

type FetchPaperContentInput struct {
	PaperNumber int
}

type FetchPaperContentOutput struct {
	*entities.Paper
}
