package entities

type PaperData struct {
	PaperNumber int
	Author      string
	Title       string
}

type Paper struct {
	PaperData
	Format  string
	Content []byte
}
