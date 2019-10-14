package main

import (
	"encoding/csv"
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	if len(os.Args) > 1 {
		namefile := os.Args[1]

		var newFile *xlsx.File
		var newSheet *xlsx.Sheet
		var newRow *xlsx.Row
		var newCell *xlsx.Cell

		year, month, day := time.Now().Date()
		output := fmt.Sprintf("%v%v%v_", year, int(month), day)

		fixUtf := func(r rune) rune {
			if r == utf8.RuneError {
				return -1
			}
			return r
		}

		newFile = xlsx.NewFile()

		excelFileName := namefile
		xlFile, err := xlsx.OpenFile(excelFileName)
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		for _, sheet := range xlFile.Sheets {

			file, err := os.Create(output + sheet.Name + ".csv")
			if err != nil {
				log.Fatal("cannot create file", err)
			}
			fmt.Printf("generate ===> %s\n", output+sheet.Name+".csv")
			defer file.Close()

			writer := csv.NewWriter(file)
			writer.Comma = ';'
			defer writer.Flush()

			newSheet, err = newFile.AddSheet(sheet.Name)
			if err != nil {
				fmt.Printf(err.Error())
			}

			j := 0
			for _, row := range sheet.Rows {

				var vals []string
				i := 0
				newRow = newSheet.AddRow()

				for _, cell := range row.Cells {

					text := cell.String()
					if len(text) > 0 {
						if j == 0 {
							newCell = newRow.AddCell()
							text2 := strings.ReplaceAll(text, "\t", " ")
							text3 := strings.ReplaceAll(text2, "\\'", "")
							text4 := strings.ReplaceAll(text3, ",", "")
							text5 := strings.ReplaceAll(text4, "'", "")
							text6 := strings.ReplaceAll(text5, ";", "")
							text7 := strings.ReplaceAll(text6, "\"", "")
							text8 := strings.ReplaceAll(text7, ".", "")
							text9 := strings.ToLower(text8)
							text10 := strings.Map(fixUtf, text9)
							newCell.Value = text10
							vals = append(vals, text10)
						} else {
							newCell = newRow.AddCell()
							text2 := strings.ReplaceAll(text, "\t", "")
							text3 := strings.ReplaceAll(text2, "\\'", "")
							text4 := strings.ReplaceAll(text3, ",", "")
							text5 := strings.ReplaceAll(text4, "'", "")
							text6 := strings.ReplaceAll(text5, ";", "")
							text7 := strings.ReplaceAll(text6, "\"", "")
							text8 := strings.ReplaceAll(text7, "\n", " ")
							text9 := strings.ReplaceAll(text8, "＋", "+")
							text10 := strings.Map(fixUtf, text9)
							newCell.Value = text10
							vals = append(vals, text10)
						}
						i++
					}
				}
				err := writer.Write(vals)
				if err != nil {
					log.Fatal("Cannot write to file", err)
				}
				j++
			}

			err = newFile.Save("newFile.xlsx")
		}

	} else {
		fmt.Printf("\nUsage : \n")
		fmt.Printf("        go run myroutine.go my.xlsx \n")
		fmt.Printf("\n")
	}
}
