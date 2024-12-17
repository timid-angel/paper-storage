package dtos

import "paper-server/domain/entities"

type AddPaperInput struct {
	entities.Paper
}

type AddPaperOutput struct {
	success bool
	message string
}

type ListPaperInput struct{}

type ListPaperOuput struct {
	papers []entities.PaperData
}

type GetPaperDetailsInput struct {
	paperNumber int
}

type GetPaperDetailsOutput struct {
	entities.PaperData
}

type FetchPaperContentInput struct {
	paperNumber int
}

type FetchPaperContentOutput struct {
	entities.Paper
}
