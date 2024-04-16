package main

import (
	"fmt"
	"log"

	"github.com/go-gota/gota/dataframe"

	"github.com/xuri/excelize/v2"
)

func loadExcel(sheetPath string, sheetName string) [][]string {
	f, err := excelize.OpenFile(sheetPath)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func main() {
	sheetPath := "example.xlsx"
	sheetName := "Sheet1"
	rows := loadExcel(sheetPath, sheetName)

	df := dataframe.LoadRecords(rows, dataframe.HasHeader(true))

	fmt.Println(df)
}
