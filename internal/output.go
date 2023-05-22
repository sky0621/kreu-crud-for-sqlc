package internal

import (
	"fmt"

	ex "github.com/sky0621/excelize-wrapper"
)

const sheetName = "CRUD"
const fileName = "CRUD.xlsx"

var tableColSet = []string{"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U"}

func Output(sqlParseResults []*SQLParseResult) error {
	tableNames := CollectTableNames(sqlParseResults)

	ew, closeFunc := ex.NewExcelizeWrapper(ex.SheetName(sheetName))
	defer func() {
		if err := closeFunc(); err != nil {
			panic(err)
		}
	}()

	ew.Set("A1", "No")
	ew.Set("B1", "SQL関数名")
	ew.Set("C1", "SQLファイル名")
	for i, tableName := range tableNames {
		col := tableColSet[i]
		ew.Set(ex.Cell(col, 1), tableName)
		ew.Width(col, calcWidthByLength(len(tableName), 1.0))
	}

	var noMaxWordCount = 1
	var sqlNameMaxWordCount = 1
	var sqlFileNameMaxWordCount = 1

	for i, x := range sqlParseResults {
		no := i + 1
		rowCnt := i + 2

		ew.Set(ex.Cell("A", rowCnt), no)
		cnt := lenInt(no)
		if cnt > noMaxWordCount {
			ew.Width("A", calcWidthByLength(cnt, 2.0))
			noMaxWordCount = cnt
		}

		ew.Set(ex.Cell("B", rowCnt), x.SQLName.ToString())
		cnt2 := len(x.SQLName.ToString())
		if cnt2 > sqlNameMaxWordCount {
			ew.Width("B", calcWidthByLength(cnt2, 3.0))
			sqlNameMaxWordCount = cnt2
		}

		ew.Set(ex.Cell("C", rowCnt), x.SQLFileName.ToString())
		cnt3 := len(x.SQLFileName.ToString())
		if cnt3 > sqlFileNameMaxWordCount {
			ew.Width("C", calcWidthByLength(cnt3, 3.0))
			sqlFileNameMaxWordCount = cnt3
		}

		for _, y := range x.TableNameWithCRUDSlice {
			for i3, tableName := range tableNames {
				if tableName == y.TableName.ToString() {
					targetCell := ex.Cell(tableColSet[i3], rowCnt)
					already, err := ew.Get(targetCell)
					if err != nil {
						return err
					}
					if already == "" {
						ew.Set(targetCell, y.CRUD.ToShortName())
					} else {
						ew.Set(targetCell, already+", "+y.CRUD.ToShortName())
					}
				}
			}
		}
	}

	ew.SaveAs(fileName)

	return nil
}

func lenInt(val int) int {
	return len(fmt.Sprintf("%d", val))
}

func calcWidthByLength(l int, coefficient float64) float64 {
	return float64(l) + coefficient
}
