package internal

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

var tableColSet = []string{"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U"}

func Output(sqlParseResults []*SQLParseResult) error {
	tableNames := CollectTableNames(sqlParseResults)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	sheetName := "CRUD"
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
		if err := f.SetCellStr(sheetName, fmt.Sprintf("%s1", tableColSet[i]), tableName); err != nil {
			return err
		}
	}

	for i, x := range sqlParseResults {
		if err := f.SetCellInt(sheetName, fmt.Sprintf("A%d", i+2), i+1); err != nil {
			return err
		}
		if err := f.SetCellStr(sheetName, fmt.Sprintf("B%d", i+2), x.SQLName.ToString()); err != nil {
			return err
		}
		if err := f.SetCellStr(sheetName, fmt.Sprintf("C%d", i+2), x.SQLFileName.ToString()); err != nil {
			return err
		}
		for _, y := range x.TableNameWithCRUDSlice {
			for i3, tableName := range tableNames {
				if tableName == y.TableName.ToString() {
					targetCell := fmt.Sprintf("%s%d", tableColSet[i3], i+2)
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
