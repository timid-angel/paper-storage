package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

var PORT = 8080

func getClient(address string) (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: " + err.Error())
	}

	return client, nil
}

func handleOperation(message string) string {
	return ""
}

func main() {
	terminalReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\033[3;36m>\033[0m \033[1;3;35m")
		message, _ := terminalReader.ReadString('\n')
		fmt.Print("\033[0m")
		response := handleOperation(message)
		fmt.Println("\033[92m\n\tMessage from server: \033[0m", response)
	}
}
