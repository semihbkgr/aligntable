package main

import (
	"fmt"
	"gopkg.in/yaml.v3"

	"github.com/semihbkgr/aligntable"
)

const multilineTableData = `
- ID: C101
  City: Skylight City
  Country: United States
  Population: 1,200,000
  Description: |-
    A bustling metropolis known for its futuristic
    architecture and vibrant cultural scene.

- ID: C102
  City: Serenity Springs
  Country: New Zealand
  Population: 350,000
  Description: |-
    A charming town renowned for its geothermal springs
    and serene natural landscapes.

- ID: C103
  City: Solaris City
  Country: Australia
  Population: 1,100,000
  Description: |-
    An eco-friendly city powered by renewable energy
    sources and known for its sustainable infrastructure.
`

func main() {
	t := aligntable.New()
	t.Rows = []*aligntable.Row{
		{
			Cells: []*aligntable.Cell{
				{Text: "ID"},
				{Text: "City"},
				{Text: "Country"},
				{Text: "Population"},
				{Text: "Description"},
			},
		},
	}

	for _, entry := range unmarshalMultilineTableData() {
		t.Rows = append(t.Rows, &aligntable.Row{
			Cells: []*aligntable.Cell{
				{Text: entry["ID"].(string)},
				{Text: entry["City"].(string)},
				{Text: entry["Country"].(string)},
				{Text: entry["Population"].(string)},
				{Text: entry["Description"].(string)},
			},
		})
	}

	fmt.Println(t)
	/*
		ID     City               Country         Population   Description
		C101   Skylight City      United States   1,200,000    A bustling metropolis known for its futuristic
		                                                       architecture and vibrant cultural scene.
		C102   Serenity Springs   New Zealand     350,000      A charming town renowned for its geothermal springs
		                                                       and serene natural landscapes.
		C103   Solaris City       Australia       1,100,000    An eco-friendly city powered by renewable energy
		                                                       sources and known for its sustainable infrastructure.
	*/
}

func unmarshalMultilineTableData() []map[string]any {
	var a []map[string]any
	err := yaml.Unmarshal([]byte(multilineTableData), &a)
	if err != nil {
		panic(err)
	}
	return a
}
