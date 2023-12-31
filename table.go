package aligntable

import (
	"fmt"
	"regexp"
	"strings"
)

type Table struct {
	Rows      []*Row
	Separator string
}

type Row struct {
	Cells []*Cell
}

func (r *Row) rowLineLen() int {
	n := 0
	for _, c := range r.Cells {
		if l := len(c.Lines()); l > n {
			n = l
		}
	}
	return n
}

type Cell struct {
	Text      string
	Alignment Alignment
}

func (c *Cell) Len() int {
	n := 0
	for _, line := range c.Lines() {
		if l := len([]rune(line)); l > n {
			n = l
		}
	}
	return n
}

func (c *Cell) Lines() []string {
	return strings.Split(extractColor(c.Text), "\n")
}

func (c *Cell) AlignedStrings(w int) []string {
	switch c.Alignment {
	case AlignLeft:
		return eachLineRaw(c.Text, func(s string, r string) string {
			return fmt.Sprintf("%s%s", s, strings.Repeat(" ", w-len([]rune(r))))
		})
	case AlignCenter:
		return eachLineRaw(c.Text, func(s string, r string) string {
			sub := (w - len([]rune(r))) / 2
			rem := (w - len([]rune(r))) % 2
			return fmt.Sprintf("%s%s%s", strings.Repeat(" ", sub), s, strings.Repeat(" ", sub+rem))
		})
	case AlignRight:
		return eachLineRaw(c.Text, func(s string, r string) string {
			return fmt.Sprintf("%s%s", strings.Repeat(" ", w-len([]rune(r))), s)
		})
	default:
		return nil
	}
}

func eachLineRaw(s string, f func(line string, raw string) string) []string {
	lines := strings.Split(s, "\n")
	rawLines := strings.Split(extractColor(s), "\n")
	for i, line := range lines {
		lines[i] = f(line, rawLines[i])
	}
	return lines
}

func extractColor(input string) string {
	re := regexp.MustCompile("\x1b\\[(\\d+;)*\\d+m")
	cleaned := re.ReplaceAllString(input, "")
	return cleaned
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
			if l := cell.Len(); l > lens[i] {
				lens[i] = l
			}
		}
	}
	return lens
}

func (t *Table) String() string {
	b := strings.Builder{}
	widths := t.ColumnsWidth()
	for _, row := range t.Rows {
		var cellsLines [][]string
		lineLen := 0
		for i, cell := range row.Cells {
			lines := cell.AlignedStrings(widths[i])
			cellsLines = append(cellsLines, lines)
			if l := len(lines); l > lineLen {
				lineLen = l
			}
		}
		for i := 0; i < lineLen; i++ {
			for cellIndex, cellLines := range cellsLines {
				if len(cellLines) > i {
					b.WriteString(cellLines[i])
				} else {
					b.WriteString(strings.Repeat(" ", widths[cellIndex]))
				}
				if cellIndex != len(row.Cells)-1 {
					b.WriteString(t.Separator)
				}
			}
			b.WriteString("\n")
		}
	}
	return b.String()
}

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)
