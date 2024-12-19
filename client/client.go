package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"slices"
	"strings"
)

var PORT = 8080

func getClient(address string) (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err.Error())
	}

	return client, nil
}

func handleOperation(message string) string {
	if message == "" {
		return "[ERROR] Empty operation"
	}

	divParts := strings.Split(message, " ")
	parts := []string{}
	for _, v := range divParts {
		v = strings.TrimSpace(v)
		if v != "" && v != " " {
			parts = append(parts, v)
		}
	}

	if len(parts) <= 2 {
		return "[ERROR] Invalid command"
	}

	if parts[0] != "paperclient" {
		return "[ERROR] Only paper client is allowed"
	}

	operations := []string{"add", "list", "detail", "fetch"}
	if !slices.Contains(operations, parts[1]) {
		return "[ERROR] Invalid operation"
	}

	client, err := getClient(parts[2])
	if err != nil {
		return "[ERROR] Unable to connect to rpc server: " + err.Error()
	}

	switch parts[1] {
	case "add":
		addParams := getAddParams(message)
		if len(addParams) != 6 {
			return "[ERROR] Invalid operation syntax"
		}

		return handleAdd(client, addParams[3], addParams[4], addParams[5])
	case "list":
		if len(parts) != 3 {
			return "[ERROR] Invalid operation syntax"
		}

		return handleList(client)

	case "detail":
		if len(parts) != 4 {
			return "[ERROR] Invalid operation syntax"
		}

		return handleDetail(client, parts[3])

	case "fetch":
		if len(parts) != 4 {
			return "[ERROR] Invalid operation syntax"
		}

		return handleFetch(client, parts[3])
	}

	return "[ERROR] Invalid operation"
}

func main() {
	terminalReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\033[3;36m>\033[0m \033[1;3;35m")
		message, _ := terminalReader.ReadString('\n')
		fmt.Print("\033[0m")
		response := handleOperation(message)

		if strings.HasPrefix(response, "[ERROR]") {
			response = fmt.Sprintf("\033[0;31m%v\033[0m", response)
		}

		if strings.HasPrefix(response, "[SUCCESS]") {
			response = fmt.Sprintf("\033[0;32m%v\033[0m", response)
		}

		fmt.Printf("\n\t> %v\n\n", response)
	}
}
