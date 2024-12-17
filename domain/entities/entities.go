package entities

type PaperData struct {
	paperNumber int
	author      string
	title       string
}

type Paper struct {
	PaperData
	format  string
	content []byte
}
