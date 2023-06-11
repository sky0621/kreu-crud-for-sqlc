package internal

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/sky0621/kreu-crud-for-sqlc/internal/parser"

	ex "github.com/sky0621/excelize-wrapper"
)

const sheetName = "CRUD"
const fileName = "CRUD.xlsx"

var tableColSet = []string{
	"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ",
	"BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ",
	"CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL", "CM", "CN", "CO", "CP", "CQ", "CR", "CS", "CT", "CU", "CV", "CW", "CX", "CY", "CZ",
	"DA", "DB", "DC", "DD", "DE", "DF", "DG", "DH", "DI", "DJ", "DK", "DL", "DM", "DN", "DO", "DP", "DQ", "DR", "DS", "DT", "DU", "DV", "DW", "DX", "DY", "DZ",
	"EA", "EB", "EC", "ED", "EE", "EF", "EG", "EH", "EI", "EJ", "EK", "EL", "EM", "EN", "EO", "EP", "EQ", "ER", "ES", "ET", "EU", "EV", "EW", "EX", "EY", "EZ",
	"FA", "FB", "FC", "FD", "FE", "FF", "FG", "FH", "FI", "FJ", "FK", "FL", "FM", "FN", "FO", "FP", "FQ", "FR", "FS", "FT", "FU", "FV", "FW", "FX", "FY", "FZ",
	"GA", "GB", "GC", "GD", "GE", "GF", "GG", "GH", "GI", "GJ", "GK", "GL", "GM", "GN", "GO", "GP", "GQ", "GR", "GS", "GT", "GU", "GV", "GW", "GX", "GY", "GZ",
	"HA", "HB", "HC", "HD", "HE", "HF", "HG", "HH", "HI", "HJ", "HK", "HL", "HM", "HN", "HO", "HP", "HQ", "HR", "HS", "HT", "HU", "HV", "HW", "HX", "HY", "HZ",
	"IA", "IB", "IC", "ID", "IE", "IF", "IG", "IH", "II", "IJ", "IK", "IL", "IM", "IN", "IO", "IP", "IQ", "IR", "IS", "IT", "IU", "IV", "IW", "IX", "IY", "IZ",
	"JA", "JB", "JC", "JD", "JE", "JF", "JG", "JH", "JI", "JJ", "JK", "JL", "JM", "JN", "JO", "JP", "JQ", "JR", "JS", "JT", "JU", "JV", "JW", "JX", "JY", "JZ",
	"KA", "KB", "KC", "KD", "KE", "KF", "KG", "KH", "KI", "KJ", "KK", "KL", "KM", "KN", "KO", "KP", "KQ", "KR", "KS", "KT", "KU", "KV", "KW", "KX", "KY", "KZ",
	"LA", "LB", "LC", "LD", "LE", "LF", "LG", "LH", "LI", "LJ", "LK", "LL", "LM", "LN", "LO", "LP", "LQ", "LR", "LS", "LT", "LU", "LV", "LW", "LX", "LY", "LZ",
	"MA", "MB", "MC", "MD", "ME", "MF", "MG", "MH", "MI", "MJ", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ",
}

func Output(rootPath string, sqlParseResults []*parser.SQLParseResult) error {
	if sqlParseResults == nil {
		return nil
	}
	for _, res := range sqlParseResults {
		log.Printf("[SQLFileName:%s][SQLName:%s]\n", res.SQLFileName, res.SQLName)
	}
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

	ew.SaveAs(filepath.Join(rootPath, fileName))

	return nil
}

func lenInt(val int) int {
	return len(fmt.Sprintf("%d", val))
}

func calcWidthByLength(l int, coefficient float64) float64 {
	return float64(l) + coefficient
}
