package aligntable

import (
	"fmt"
	"strings"
)

type Tree struct {
	Nodes []*Node
}

type Node struct {
	Text     string
	SubNodes []*Node
}

func (n *Node) Len() int {
	lines := strings.Split(n.Text, "\n")
	i := 0
	for _, line := range lines {
		if l := len([]rune(line)); l > i {
			i = l
		}
	}
	return i
}

func (n *Node) NodeRows(skip int, column int) []*Row {
	rows := make([]*Row, 1, 1+len(n.SubNodes))
	cells := make([]*Cell, column)
	for i := 0; i < skip; i++ {
		cells[i] = &Cell{}
	}
	cells[skip] = &Cell{Text: n.Text, Alignment: AlignLeft}
	for i := skip + 1; i < column; i++ {
		cells[i] = &Cell{}
	}
	rows[0] = &Row{Cells: cells}
	for _, node := range n.SubNodes {
		rows = append(rows, node.NodeRows(skip+1, column)...)
	}
	return rows
}

func (n *Node) Width(i int) int {
	i++
	c := i
	for _, node := range n.SubNodes {
		t := node.Width(i)
		if t > c {
			c = t
		}
	}
	return c
}

func (t *Tree) Table() *Table {
	table := New()
	r := &Node{SubNodes: t.Nodes}
	w := r.Width(-1)
	for _, node := range t.Nodes {
		table.Rows = append(table.Rows, node.NodeRows(0, w)...)
	}
	setArrows(table)
	table.Separator = " "
	return table
}

func setArrows(t *Table) {
	colsW := t.ColumnsWidth()
	for i := 0; i < len(colsW); i++ {
		var parentIndex int
		var childIndexes []int
		for rowIndex, row := range t.Rows {
			if row.Cells[i].Text != "" {
				cellArrows(t, parentIndex, childIndexes, i, colsW[i])
				parentIndex = rowIndex
				childIndexes = []int{}
			} else if i+1 < len(row.Cells) && row.Cells[i+1].Text != "" {
				childIndexes = append(childIndexes, rowIndex)
			}
		}
		cellArrows(t, parentIndex, childIndexes, i, colsW[i])
	}
}

func cellArrows(t *Table, parentIndex int, childIndexes []int, column int, columnWidth int) {
	if len(childIndexes) == 0 {
		return
	}
	for arrowIndex := parentIndex + 1; arrowIndex <= childIndexes[len(childIndexes)-1]; arrowIndex++ {
		cell := t.Rows[arrowIndex].Cells[column]
		if containsInt(childIndexes, arrowIndex) {
			if arrowIndex == childIndexes[len(childIndexes)-1] {
				cell.Text = fmt.Sprintf("└%s", strings.Repeat("─", columnWidth-1))
			} else {
				cell.Text = fmt.Sprintf("├%s", strings.Repeat("─", columnWidth-1))
				for i := 1; i < t.Rows[arrowIndex].rowLineLen(); i++ {
					if i == 0 {
						cell.Text += "│"
					} else {
						cell.Text += "\n│"
					}
				}
			}
		} else {
			for i := 0; i < t.Rows[arrowIndex].rowLineLen(); i++ {
				if i == 0 {
					cell.Text += "│"
				} else {
					cell.Text += "\n│"
				}
			}
		}
		cell.Alignment = AlignLeft
	}
}

func containsInt(s []int, n int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}
