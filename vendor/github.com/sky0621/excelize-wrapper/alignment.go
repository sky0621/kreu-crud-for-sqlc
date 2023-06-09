package excelize_wrapper

import "github.com/xuri/excelize/v2"

type AlignmentOption func(a *excelize.Alignment)

const (
	HLeft   = "left"
	HCenter = "center"
	HRight  = "right"

	VCenter = "center"
)

var VhCenterAlignment = NewAlignment(Horizontal(HCenter), Vertical(VCenter), WrapText(true), Indent(1))
var HRightAlignment = NewAlignment(Horizontal(HRight), Vertical(VCenter), WrapText(true), Indent(1))
var HLeftAlignment = NewAlignment(Horizontal(HLeft), Vertical(VCenter), WrapText(true), Indent(1))
var HLeftAlignmentIndent2 = NewAlignment(Horizontal(HLeft), Vertical(VCenter), WrapText(true), Indent(2))
var HLeftAlignmentIndent3 = NewAlignment(Horizontal(HLeft), Vertical(VCenter), WrapText(true), Indent(3))
var HLeftAlignmentIndent4 = NewAlignment(Horizontal(HLeft), Vertical(VCenter), WrapText(true), Indent(4))

func Horizontal(s string) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.Horizontal = s
	}
}

func Indent(i int) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.Indent = i
	}
}

func JustifyLastLine(b bool) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.JustifyLastLine = b
	}
}

func ReadingOrder(r uint64) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.ReadingOrder = r
	}
}

func RelativeIndent(r int) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.RelativeIndent = r
	}
}

func ShrinkToFit(b bool) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.ShrinkToFit = b
	}
}

func TextRotation(t int) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.TextRotation = t
	}
}

func Vertical(v string) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.Vertical = v
	}
}

func WrapText(w bool) AlignmentOption {
	return func(a *excelize.Alignment) {
		a.WrapText = w
	}
}

func NewAlignment(options ...AlignmentOption) *excelize.Alignment {
	s := &excelize.Alignment{}
	for _, option := range options {
		option(s)
	}
	return s
}
