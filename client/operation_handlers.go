package main

import (
	"fmt"
	"net/rpc"
	"os"
	"paper-server/domain/dtos"
	"paper-server/domain/entities"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func getAddParams(message string) []string {
	message += " "
	parts := []string{}
	cn := 0
	curr := ""
	for i := 0; i < len(message); i++ {
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
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "[ERROR] Unable to read file: " + err.Error()
	}

	paperData := entities.PaperData{
		Author: authorName,
		Title:  paperTitle,
	}
	fileExtension := strings.ToLower(strings.TrimPrefix(filepath.Ext(filePath), "."))
	args := dtos.AddPaperInput{
		Paper: &entities.Paper{
			PaperData: paperData,
			Content:   fileContent,
			Format:    fileExtension,
		},
	}
	reply := &dtos.AddPaperOutput{}
	err = client.Call("PaperStorage.AddPaper", args, reply)
	if err != nil {
		return "[ERROR] Unable to add paper: " + err.Error()
	}

	return fmt.Sprintf("[SUCCESS] Add paper '%v, %v' in %v format with PNB = %v", paperTitle, authorName, fileExtension, reply.PaperNumber)
}

func handleList(client *rpc.Client) string {
	args := dtos.ListPaperInput{}
	reply := &dtos.ListPaperOuput{}
	err := client.Call("PaperStorage.ListPapers", args, reply)
	if err != nil {
		return "[ERROR] Unable to list papers: " + err.Error()
	}

	if len(reply.Papers) == 0 {
		return "There are currently no papers in the server"
	}

	sort.Slice(reply.Papers, func(i, j int) bool {
		return reply.Papers[i].PaperNumber < reply.Papers[j].PaperNumber
	})
	response := "List of Papers: "
	for _, paper := range reply.Papers {
		response += fmt.Sprintf("\n\t\t \033[1;35mPNB = %v:\033[0m \033[1;34m'%v' by %v\033[0m", paper.PaperNumber, paper.Title, paper.Author)
	}

	return response
}

func handleDetail(client *rpc.Client, paperID string) string {
	paperNumber, err := strconv.ParseInt(paperID, 10, 32)
	if err != nil {
		return "[ERROR] Invalid paper number"
	}

	args := dtos.GetPaperDetailsInput{
		PaperNumber: int(paperNumber),
	}
	reply := &dtos.GetPaperDetailsOutput{}
	err = client.Call("PaperStorage.GetPaperDetails", args, reply)
	if err != nil {
		return "[ERROR] Unable to get the details of paper: " + err.Error()
	}

	response := "Paper details: \n"
	response += fmt.Sprintf("\t    %-10s \033[1;34m%v\033[0m\n", "ID:", reply.PaperData.PaperNumber)
	response += fmt.Sprintf("\t    %-10s \033[1;34m%v\033[0m\n", "Title:", reply.PaperData.Title)
	response += fmt.Sprintf("\t    %-10s \033[1;34m%v\033[0m\n", "Author:", reply.PaperData.Author)
	return response
}

func handleFetch(client *rpc.Client, paperID string) string {
	paperNumber, err := strconv.ParseInt(paperID, 10, 32)
	if err != nil {
		return "[ERROR] Invalid paper number"
	}

	args := dtos.FetchPaperContentInput{
		PaperNumber: int(paperNumber),
	}
	reply := &dtos.FetchPaperContentOutput{}
	err = client.Call("PaperStorage.FetchPaperContent", args, reply)
	if err != nil {
		return "[ERROR] Unable to fetch paper: " + err.Error()
	}

	response := "Paper details: \n"
	response += fmt.Sprintf("\t    %-10s \033[1;34m%v\033[0m\n", "ID:", reply.PaperData.PaperNumber)
	response += fmt.Sprintf("\t    %-10s \033[1;34m%v\033[0m\n", "Title:", reply.PaperData.Title)
	response += fmt.Sprintf("\t    %-10s \033[1;34m%v\033[0m\n", "Author:", reply.PaperData.Author)
	if reply.Format == "txt" {
		delimiter := fmt.Sprintf("\n%v\n", strings.Repeat("-", 100))
		response += fmt.Sprintf("\n\033[0;32mPaper content:\033[0m%v%v%v", delimiter, string(reply.Content), delimiter)
	} else {
		fileName := fmt.Sprintf("%s by %s.%s", reply.Title, reply.Author, reply.Format)
		err = os.WriteFile(fileName, reply.Content, 0644)
		if err != nil {
			return "[ERROR] Unable to write file: " + err.Error()
		}

		response += fmt.Sprintf("\n\t\033[0;32mPaper fetched and saved as '%v'\033[0m", fileName)
	}

	return response
}
