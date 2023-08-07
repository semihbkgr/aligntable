package aligntable

import "fmt"

type Tree struct {
	Nodes []*Node
}

type Node struct {
	Text     string
	SubNodes []*Node
}

func (n *Node) Rows(skip int, column int, last bool, parentLast bool) []*Row {
	rows := make([]*Row, 1)
	cells := make([]Cell, column)
	if skip != 0 {

		for i := 0; i < skip; i++ {
			if i == skip-2 && parentLast {
				break
			}
			cells[i].Text = "│"
		}

		if last {
			cells[skip-1].Text = "└───────────────────"
		} else {
			cells[skip-1].Text = "├───────────────────"
		}

	}
	cells[skip] = Cell{Text: n.Text}
	rows[0] = &Row{Cells: cells}
	for i, node := range n.SubNodes {
		rows = append(rows, node.Rows(skip+1, column, i == len(n.SubNodes)-1, last)...)
	}
	return rows
}

func (n *Node) Count(i int) int {
	i++
	c := i
	for _, node := range n.SubNodes {
		t := node.Count(i)
		if t > c {
			c = t
		}
	}
	return c
}

func (t *Tree) Table() *Table {
	table := New()
	c := (&Node{SubNodes: t.Nodes}).Count(-1)
	fmt.Println(c)
	for i, node := range t.Nodes {
		table.Rows = append(table.Rows, node.Rows(0, c, i == len(t.Nodes)-1, false)...)
	}
	table.Separator = " "
	return table
}
