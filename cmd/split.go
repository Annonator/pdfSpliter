package cmd

import (
	"fmt"
	"github.com/ledongthuc/pdf"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	rootCommand.AddCommand(splitPdfByPageCommand)
}

var splitPdfByPageCommand = &cobra.Command{
	Use:   "split [inputFile] [outputDirectory]",
	Short: "split a PDF file into multiple files",
	Long:  "split a PDF file into multiple files. It takes a PDF file and a list of page numbers and creates a new PDF file for each page.",
	Run: func(cmd *cobra.Command, args []string) {
		if !isInputFileValid(args[0]) {
			panic("Input file not found or not a PDF file")
		}

		if !isOutputDirectoryValid(args[1]) {
			panic("Output directory not found or not empty")
		}

		err := os.MkdirAll(args[1], 0755)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic("Can't open input file")
		}
		importer := gofpdi.NewImporter()
		//Yes, we do need to use another PDF package to get the number of pages because gofpdf or gofpdi are not compatible with each other, its super broken
		_, reader, err := pdf.Open(args[0])
		if err != nil {
			panic("Can't open input file to read page count")
		}
		pageCount := reader.NumPage()

		fmt.Println("Start creating files")

		for i := 1; i <= pageCount; i++ {
			newPdf := gofpdf.New("P", "mm", "A4", "")
			template := importer.ImportPage(newPdf, args[0], i, "/MediaBox")

			newPdf.AddPage()

			importer.UseImportedTemplate(newPdf, template, 0, 0, 210, 297)

			err := newPdf.OutputFileAndClose(args[1] + "/page_" + strconv.Itoa(i) + ".pdf")
			if err != nil {
				panic("Can't write to output file")
			}
		}
		fmt.Println("Done")
	},
}

func isInputFileValid(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	if filepath.Ext(file) != ".pdf" {
		return false
	}

	return true
}

// we require the output directory to be empty / not existing
func isOutputDirectoryValid(dir string) bool {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return false
	}

	return true
}
