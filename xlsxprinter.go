package textslicer

import (
	"container/list"
	"fmt"

	"github.com/tealeg/xlsx"
)

type XlsxPrinter struct {
	file       *xlsx.File
	sheetIndex int
}

func NewXlsxPrinter(f *xlsx.File) *XlsxPrinter {
	return &XlsxPrinter{
		file: f,
	}

}

func (xp *XlsxPrinter) Proc(chunk *list.List) {
	xp.sheetIndex++
	sheet, err := xp.file.AddSheet(fmt.Sprintf("sheet%d", xp.sheetIndex))
	if err != nil {
		panic(err)
	}
	rowIndex := 0
	for e := chunk.Front(); e != nil; e = e.Next() {
		row := sheet.AddRow()
		row.WriteSlice(e.Value, rowIndex)
		rowIndex++
	}
}
