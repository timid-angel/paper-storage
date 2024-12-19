package paper_usecase

import (
	"fmt"
	"log"
	"paper-server/domain/dtos"
	"paper-server/domain/entities"
	"paper-server/server/domain"
)

type PaperStorageUsecase struct {
	storageRepository   domain.IStorageRepository
	notificationService domain.INotificationService
}

func NewPaperStorageUsecase(storageRepository domain.IStorageRepository, notificationService domain.INotificationService) *PaperStorageUsecase {
	return &PaperStorageUsecase{
		storageRepository:   storageRepository,
		notificationService: notificationService,
	}
}

func (usecase *PaperStorageUsecase) AddPaper(data *dtos.AddPaperInput) (int, domain.IDomainError) {
	data.Paper.PaperNumber = usecase.storageRepository.GetNewPaperNumber()
	paperNumber, err := data.Paper.PaperNumber, usecase.storageRepository.AddPaper(data.Paper)
	if err != nil {
		return paperNumber, err
	}

	err = usecase.notificationService.PublishNotification(fmt.Sprintf("New paper added: '%v' by %v with PNB = %v", data.Title, data.Author, paperNumber))
	if err != nil {
		log.Println("failed to send notification: " + err.Error())
	}
	return paperNumber, err
}

func (usecase *PaperStorageUsecase) ListPapers() (*[]entities.PaperData, domain.IDomainError) {
	return usecase.storageRepository.ListPapers()
}

func (usecase *PaperStorageUsecase) GetPaperDetails(paperNumber int) (*entities.PaperData, domain.IDomainError) {
	return usecase.storageRepository.GetPaperDetails(paperNumber)
}

func (usecase *PaperStorageUsecase) FetchPaperContent(paperNumber int) (*entities.Paper, domain.IDomainError) {
	return usecase.storageRepository.FetchPaperContent(paperNumber)
}
