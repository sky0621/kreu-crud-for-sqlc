package excelize_wrapper

import "fmt"

func Cell(col string, row int) string {
	return fmt.Sprintf("%s%d", col, row)
}
