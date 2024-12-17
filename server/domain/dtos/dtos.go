package dtos

import "paper-server/server/domain"

type AddPaperInput struct {
	domain.Paper
}

type AddPaperOutput struct {
	success bool
	message string
}

type ListPaperInput struct{}

type ListPaperOuput struct {
	papers []domain.PaperData
}

type GetPaperDetailsInput struct {
	paperNumber int
}

type GetPaperDetailsOutput struct {
	domain.PaperData
}

type FetchPaperContentInput struct {
	paperNumber int
}

type FetchPaperContentOutput struct {
	domain.Paper
}
