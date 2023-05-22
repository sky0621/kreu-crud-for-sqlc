package excelize_wrapper

import "github.com/xuri/excelize/v2"

// TODO: １シートのみ使うことを前提としている。複数シート使う場合は再検討。

type ExcelizeWrapperOption func(file *excelize.File)

func SheetName(name string) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.SetSheetName("Sheet1", name); err != nil {
			panic(err)
		}
	}
}

func PaperSize(size int) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.SetPageLayout(getDefaultSheetName(f), &excelize.PageLayoutOptions{
			Size: &size,
		}); err != nil {
			panic(err)
		}
	}
}

func SheetPassword(pass string) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.ProtectSheet(getDefaultSheetName(f), &excelize.SheetProtectionOptions{
			AlgorithmName: "SHA-512",
			Password:      pass,
			EditScenarios: false,
		}); err != nil {
			panic(err)
		}
	}
}

func SheetView() ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.SetSheetView(getDefaultSheetName(f), 0, &excelize.ViewOptions{
			ShowGridLines: func() *bool {
				b := false
				return &b
			}(),
		}); err != nil {
			panic(err)
		}
	}
}

type CloseFunc = func() error

func NewExcelizeWrapper(options ...ExcelizeWrapperOption) (Wrapper, CloseFunc) {
	f := excelize.NewFile()
	for _, option := range options {
		option(f)
	}
	return &wrapper{f: f}, func() error {
		if f != nil {
			if err := f.Close(); err != nil {
				return err
			}
		}
		return nil
	}
}

type Wrapper interface {
	Set(position string, val any)
	Get(position string) (string, error)
	Merge(from, to string)
	Height(row int, h float64)
	Width(cell string, wd float64)
	RangeWidth(start, end string, wd float64)
	Text(cell string, settings []excelize.RichTextRun)
	CellStyle(cell string, style *excelize.Style)
	CellRangeStyle(start, end string, style *excelize.Style)
	CellExternalHyperLink(cell, url string)
	AddPicture(cell, path string)
	InsertPageBreak(cell string)

	SaveAs(name string)
}

type wrapper struct {
	f *excelize.File
}

func getDefaultSheetName(f *excelize.File) string {
	return f.GetSheetName(0)
}

func (w *wrapper) Set(position string, val interface{}) {
	if err := w.f.SetCellValue(getDefaultSheetName(w.f), position, val); err != nil {
		panic(err)
	}
}

func (w *wrapper) Get(targetCell string) (string, error) {
	already, err := w.f.GetCellValue(getDefaultSheetName(w.f), targetCell)
	if err != nil {
		return "", err
	}
	return already, nil
}

func (w *wrapper) Merge(from, to string) {
	if err := w.f.MergeCell(getDefaultSheetName(w.f), from, to); err != nil {
		panic(err)
	}
}

func (w *wrapper) Height(row int, h float64) {
	if err := w.f.SetRowHeight(getDefaultSheetName(w.f), row, h); err != nil {
		panic(err)
	}
}

func (w *wrapper) Width(cell string, wd float64) {
	if err := w.f.SetColWidth(getDefaultSheetName(w.f), cell, cell, wd); err != nil {
		panic(err)
	}
}

func (w *wrapper) RangeWidth(start, end string, wd float64) {
	if err := w.f.SetColWidth(getDefaultSheetName(w.f), start, end, wd); err != nil {
		panic(err)
	}
}

func (w *wrapper) Text(cell string, settings []excelize.RichTextRun) {
	if err := w.f.SetCellRichText(getDefaultSheetName(w.f), cell, settings); err != nil {
		panic(err)
	}
}

func (w *wrapper) style(s *excelize.Style) int {
	styleID, err := w.f.NewStyle(s)
	if err != nil {
		panic(err)
	}
	return styleID
}

func (w *wrapper) CellStyle(cell string, style *excelize.Style) {
	if err := w.f.SetCellStyle(getDefaultSheetName(w.f), cell, cell, w.style(style)); err != nil {
		panic(err)
	}
}
func (w *wrapper) CellRangeStyle(start, end string, style *excelize.Style) {
	if err := w.f.SetCellStyle(getDefaultSheetName(w.f), start, end, w.style(style)); err != nil {
		panic(err)
	}
}

func (w *wrapper) CellExternalHyperLink(cell, url string) {
	if err := w.f.SetCellHyperLink(getDefaultSheetName(w.f), cell, url, "External", excelize.HyperlinkOpts{
		Display: &url,
	}); err != nil {
		panic(err)
	}
}

func (w *wrapper) AddPicture(cell, path string) {
	if err := w.f.AddPicture(getDefaultSheetName(w.f), cell, path, nil); err != nil {
		panic(err)
	}
}

func (w *wrapper) InsertPageBreak(cell string) {
	if err := w.f.InsertPageBreak(getDefaultSheetName(w.f), cell); err != nil {
		panic(err)
	}
}

func (w *wrapper) SaveAs(name string) {
	if err := w.f.SaveAs(name); err != nil {
		panic(err)
	}
}
