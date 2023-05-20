package internal

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const sheetName = "CRUD"

var tableColSet = []string{"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U"}

func Output(sqlParseResults []*SQLParseResult) error {
	tableNames := CollectTableNames(sqlParseResults)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	if err := f.SetSheetName("Sheet1", sheetName); err != nil {
		return err
	}

	if err := f.SetCellStr(sheetName, "A1", "No"); err != nil {
		return err
	}
	if err := f.SetCellStr(sheetName, "B1", "SQL関数名"); err != nil {
		return err
	}
	if err := f.SetCellStr(sheetName, "C1", "SQLファイル名"); err != nil {
		return err
	}
	for i, tableName := range tableNames {
		col := tableColSet[i]
		cnt, err := SetCellStr(f, col, 1, tableName)
		if err != nil {
			return err
		}
		w := calcWidthByLength(cnt, 1.0)
		if err := f.SetColWidth(sheetName, col, col, w); err != nil {
			return err
		}
	}

	var noMaxWordCount = 1
	var sqlNameMaxWordCount = 1
	var sqlFileNameMaxWordCount = 1

	for i, x := range sqlParseResults {
		no := i + 1
		rowCnt := i + 2

		cnt, err := SetCellInt(f, "A", rowCnt, no)
		if err != nil {
			return err
		}
		if cnt > noMaxWordCount {
			if err := f.SetColWidth(sheetName, "A", "A", calcWidthByLength(cnt, 2.0)); err != nil {
				return err
			}
			noMaxWordCount = cnt
		}

		cnt2, err := SetCellStr(f, "B", rowCnt, x.SQLName.ToString())
		if err != nil {
			return err
		}
		if cnt2 > sqlNameMaxWordCount {
			if err := f.SetColWidth(sheetName, "B", "B", calcWidthByLength(cnt2, 3.0)); err != nil {
				return err
			}
			sqlNameMaxWordCount = cnt2
		}

		cnt3, err := SetCellStr(f, "C", rowCnt, x.SQLFileName.ToString())
		if err != nil {
			return err
		}
		if cnt3 > sqlFileNameMaxWordCount {
			if err := f.SetColWidth(sheetName, "C", "C", calcWidthByLength(cnt3, 3.0)); err != nil {
				return err
			}
			sqlFileNameMaxWordCount = cnt3
		}

		for _, y := range x.TableNameWithCRUDSlice {
			for i3, tableName := range tableNames {
				if tableName == y.TableName.ToString() {
					targetCell := fmt.Sprintf("%s%d", tableColSet[i3], rowCnt)
					already, err := f.GetCellValue(sheetName, targetCell)
					if err != nil {
						return err
					}
					if already == "" {
						if err := f.SetCellStr(sheetName, targetCell, y.CRUD.ToShortName()); err != nil {
							return err
						}
					} else {
						if err := f.SetCellStr(sheetName, targetCell, already+", "+y.CRUD.ToShortName()); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	if err := f.SaveAs("CRUD.xlsx"); err != nil {
		return err
	}

	return nil
}

func SetCellInt(f *excelize.File, col string, row int, val int) (int, error) {
	if err := f.SetCellInt(sheetName, fmt.Sprintf("%s%d", col, row), val); err != nil {
		return 0, err
	}
	return len(fmt.Sprintf("%d", val)), nil
}

func SetCellStr(f *excelize.File, col string, row int, val string) (int, error) {
	if err := f.SetCellStr(sheetName, fmt.Sprintf("%s%d", col, row), val); err != nil {
		return 0, err
	}
	return len(val), nil
}

func calcWidthByLength(l int, coefficient float64) float64 {
	return float64(l) + coefficient
}
