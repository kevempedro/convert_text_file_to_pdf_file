package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	fmt.Println("---------- CONVERTING FILE... ----------")

	fileName := "file_text.txt"

	fileReaded, err := readTextFile(fileName)

	if err != nil {
		log.Fatalf("'%s' text file doesn't exist!", fileName)
	}

	generatePdfFile(fileReaded)

	fmt.Println("---------- CONVERTION FINISHED! ----------")
}

func readTextFile(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func generatePdfFile(content []byte) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "", 16)
	pdf.MultiCell(200, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("file_pdf.pdf")
}
