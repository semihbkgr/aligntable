package main

import (
	"fmt"
	"strconv"

	"github.com/semihbkgr/aligntable"
	"gopkg.in/yaml.v3"
)

const tableData = `
- Name: Alice
  Age: 28
  City: New York
- Name: Bob
  Age: 35
  City: Los Angeles
- Name: Connor
  Age: 22
  City: Chicago
- Name: Daniel
  Age: 30
  City: Houston
- Name: Emily
  Age: 29
  City: San Francisco
`

func main() {
	t := aligntable.New()
	t.Rows = []*aligntable.Row{
		{
			Cells: []*aligntable.Cell{
				{Text: "NAME"},
				{Text: "AGE"},
				{Text: "CITY"},
			},
		},
	}

	for _, entry := range unmarshalTableData() {
		t.Rows = append(t.Rows, &aligntable.Row{
			Cells: []*aligntable.Cell{
				{Text: entry["Name"].(string)},
				{Text: strconv.Itoa(entry["Age"].(int))},
				{Text: entry["City"].(string)},
			},
		})
	}

	fmt.Println(t)
	/*
		NAME     AGE   CITY
		Alice    28    New York
		Bob      35    Los Angeles
		Connor   22    Chicago
		Daniel   30    Houston
		Emily    29    San Francisco
	*/
}

func unmarshalTableData() []map[string]any {
	var a []map[string]any
	err := yaml.Unmarshal([]byte(tableData), &a)
	if err != nil {
		panic(err)
	}
	return a
}
