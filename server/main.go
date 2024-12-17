package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"paper-server/server/controller"
	storage_repository "paper-server/server/repository"
	paper_usecase "paper-server/server/usecase"
)

func main() {
	repository := storage_repository.NewStorageRepository()
	usecase := paper_usecase.NewPaperStorageUsecase(repository)
	controller := controller.NewPaperStorage(usecase)
	if len(os.Args) != 2 {
		log.Fatalln("[ERROR] The port must be passed as an argument")
	}

	serverAddress := os.Args[2]
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
