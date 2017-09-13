package main

import (
	"fmt"
	"Go-freexl/freexl"
)

func main() {
	version := freexl.FreeXLVerison()
	fmt.Println(version)
	handle, err := freexl.FreeXLOpen("t.xls")
	if err != nil {
		fmt.Println("FreeXLOpen error:", err)
	}
	
		var index uint16
		index, err = freexl.FreeXLGetActiveWorksheet(handle)
		fmt.Println("FreeXLGetActiveWorksheet:", index)
		err = freexl.FreeXLSelectActiveWorksheet(handle, 0)
		if err != nil {
			fmt.Println("FreeXLSelectActiveWorksheetr:", err)
		}
	
	rows, columns, err := freexl.FreeXLWorksheetDimensions(handle)
	fmt.Println("FreeXLWorksheetDimensions:", rows, columns, err)

	for rownum := uint(0); rownum < rows; rownum++ {
		fmt.Printf("[%d]\t", rownum)
		for column := uint16(0); column < columns; column++ {
			col := ""
			col, err = freexl.FreeXLGetCellValue(handle, rownum, column)
			if err != nil {
				fmt.Println("FreeXLGetCellValue:", err)
			}
			fmt.Printf(" [%d] %s \t", column, col)
		}
		fmt.Printf("\n")
	}

	err = freexl.FreeXLClose(handle)
	if err != nil {
		fmt.Println("FreeXLClose:", err)
	}
}
