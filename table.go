package aligntable

import (
	"fmt"
	"strings"
)

type Table struct {
	Rows      []*Row
	Separator string
}

type Row struct {
	Cells []Cell
}

type Cell struct {
	Text string
}

func New() *Table {
	return &Table{
		Separator: "   ",
	}
}

func (t *Table) ColumnsWidth() []int {
	if len(t.Rows) == 0 {
		return nil
	}
	lens := make([]int, len(t.Rows[0].Cells))
	for _, row := range t.Rows {
		for i, cell := range row.Cells {
			if l := len([]rune(cell.Text)); l > lens[i] {
				lens[i] = l
			}
		}
	}
	return lens
}

func (t *Table) String() string {
	b := strings.Builder{}
	widths := t.ColumnsWidth()
	fmt.Println(widths)
	for _, row := range t.Rows {
		for i, cell := range row.Cells {
			b.WriteString(cell.Text)
			if s := widths[i] - len([]rune(cell.Text)); s > 0 {
				b.WriteString(strings.Repeat(" ", s))
			}
			if i != len(row.Cells)-1 {
				b.WriteString(t.Separator)
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
