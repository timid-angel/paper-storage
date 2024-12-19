package main

import (
	"fmt"
	"net/rpc"
	"os"
	"paper-server/domain/dtos"
	"paper-server/domain/entities"
	"path/filepath"
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

	return "ADD called;" + authorName + ";" + paperTitle + ";" + filePath
}

func handleList(client *rpc.Client) string {
	args := dtos.ListPaperInput{}
	reply := &dtos.ListPaperOuput{}
	err := client.Call("PaperStorage.ListPapers", args, reply)
	if err != nil {
		return "[ERROR] Unable to list papers: " + err.Error()
	}

	return fmt.Sprintf("%v", reply.Papers)
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
	return fmt.Sprintf("%v", reply)
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

	fileName := fmt.Sprintf("%s - %s.%s", reply.Title, reply.Author, reply.Format)
	err = os.WriteFile(fileName, reply.Content, 0644)
	if err != nil {
		return "[ERROR] Unable to write file: " + err.Error()
	}

	return fmt.Sprintf("Paper fetched and saved as %s", fileName)
}
