package excelize_wrapper

import "github.com/xuri/excelize/v2"

type Border struct {
	Color string
	Style int
}

func (b *Border) setBorder(s string) excelize.Border {
	return excelize.Border{Type: s, Color: b.Color, Style: b.Style}
}

func (b *Border) LeftBorder() excelize.Border {
	return b.setBorder("left")
}

func (b *Border) TopBorder() excelize.Border {
	return b.setBorder("top")
}

func (b *Border) RightBorder() excelize.Border {
	return b.setBorder("right")
}

func (b *Border) BottomBorder() excelize.Border {
	return b.setBorder("bottom")
}

func (b *Border) FullBorder() []excelize.Border {
	return []excelize.Border{
		b.LeftBorder(),
		b.TopBorder(),
		b.RightBorder(),
		b.BottomBorder(),
	}
}

func (b *Border) SideBorder() []excelize.Border {
	return []excelize.Border{
		b.LeftBorder(),
		b.RightBorder(),
	}
}

func (b *Border) TopBottomBorder() []excelize.Border {
	return []excelize.Border{
		b.TopBorder(),
		b.BottomBorder(),
	}
}

func (b *Border) LeftTopBottomBorder() []excelize.Border {
	return []excelize.Border{
		b.LeftBorder(),
		b.TopBorder(),
		b.BottomBorder(),
	}
}

func (b *Border) RightTopBottomBorder() []excelize.Border {
	return []excelize.Border{
		b.RightBorder(),
		b.TopBorder(),
		b.BottomBorder(),
	}
}

func (b *Border) SideBottomBorder() []excelize.Border {
	return []excelize.Border{
		b.LeftBorder(),
		b.RightBorder(),
		b.BottomBorder(),
	}
}
