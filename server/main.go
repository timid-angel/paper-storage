package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"paper-server/config"
	"paper-server/server/controller"
	storage_repository "paper-server/server/repository"
	"paper-server/server/services/rabbitmq"
	paper_usecase "paper-server/server/usecase"
)

func main() {
	config.LoadEnvironmentVariables(".env")

	storageRepository := storage_repository.NewStorageRepository()
	notificationService := rabbitmq.NewRabbitMqService(os.Getenv("RABBIT_MQ_ADDRESS"), "add_notification")
	usecase := paper_usecase.NewPaperStorageUsecase(storageRepository, notificationService)
	controller := controller.NewPaperStorage(usecase)

	serverAddress := os.Getenv("HOST_ADDRESS")
	err := rpc.Register(controller)
	if err != nil {
		log.Fatalln("[ERROR] Failed to start server: " + err.Error())
	}

	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalln("[ERROR] Failed to start server: " + err.Error())
	}

	defer listener.Close()
	log.Println("Server listening on " + serverAddress)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ERROR] Connection error: " + err.Error())
			continue
		}

		go rpc.ServeConn(conn)
	}
}
