package excelize_wrapper

import "github.com/xuri/excelize/v2"

type ExcelizeStyleOption func(style *excelize.Style)

func Alignment(a *excelize.Alignment) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Alignment = a
	}
}

func Borders(b []excelize.Border) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Border = b
	}
}

func Font(f *excelize.Font) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Font = f
	}
}

func Fill(f excelize.Fill) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Fill = f
	}
}

func NewStyle(options ...ExcelizeStyleOption) *excelize.Style {
	s := &excelize.Style{}
	for _, option := range options {
		option(s)
	}
	return s
}
