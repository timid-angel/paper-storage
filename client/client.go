package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"slices"
	"strconv"
	"strings"
)

var PORT = 8080

func getClient(address string) (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: " + err.Error())
	}

	return client, nil
}

func getAddParams(message string) []string {
	message += " "
	parts := []string{}
	cn := 0
	curr := ""
	for i, _ := range message {
		c := string(message[i])
		if c == "'" {
			cn = (cn + 1) % 2
			continue
		}

		if c == " " && cn == 0 {
			if len(curr) > 0 {
				parts = append(parts, strings.TrimSpace(curr))
			}

			curr = ""
		}

		curr += c
	}

	return parts
}

func handleAdd(client *rpc.Client, authorName string, paperTitle string, filePath string) string {
	return "ADD called;" + authorName + ";" + paperTitle + ";" + filePath
}

func handleList(client *rpc.Client) string {
	return "LIST called"
}

func handleDetail(client *rpc.Client, paperID string) string {
	paperNumber, err := strconv.ParseInt(paperID, 10, 32)
	if err != nil {
		return "[ERROR] Invalid paper number"
	}

	return "DETAIL called" + string(paperNumber)
}

func handleFetch(client *rpc.Client, paperID string) string {
	paperNumber, err := strconv.ParseInt(paperID, 10, 32)
	if err != nil {
		return "[ERROR] Invalid paper number"
	}

	return "FETCH called" + string(paperNumber)
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
		fmt.Printf("\033[92m\n\tMessage from server:\033[0m %v\n\n", response)
	}
}
