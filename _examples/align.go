package main

import (
	"fmt"
	"strconv"

	"github.com/semihbkgr/aligntable"
	"gopkg.in/yaml.v3"
)

const alignTableData = `
- Image ID: '987654'
  Repository: web-app
  Tag: latest
  Size (MB): 350
  Created Date: 2023-07-15 08:30

- Image ID: '654321'
  Repository: database
  Tag: v2
  Size (MB): 480
  Created Date: 2023-06-28 14:15

- Image ID: '789012'
  Repository: nginx
  Tag: stable
  Size (MB): 180
  Created Date: 2023-07-02 11:45

- Image ID: '123789'
  Repository: api-service
  Tag: prod
  Size (MB): 220
  Created Date: 2023-06-15 18:20
`

func main() {
	t := aligntable.New()
	t.Rows = []*aligntable.Row{
		{
			Cells: []*aligntable.Cell{
				{Text: "Image ID"},
				{Text: "Repository"},
				{
					Text:      "Tag",
					Alignment: aligntable.AlignCenter,
				},
				{Text: "Size (MB)"},
				{
					Text:      "Created Date",
					Alignment: aligntable.AlignRight,
				},
			},
		},
	}

	for _, entry := range unmarshalAlignTableData() {
		t.Rows = append(t.Rows, &aligntable.Row{
			Cells: []*aligntable.Cell{
				{
					Text:      entry["Image ID"].(string),
					Alignment: aligntable.AlignRight,
				},
				{Text: entry["Repository"].(string)},
				{
					Text:      entry["Tag"].(string),
					Alignment: aligntable.AlignCenter,
				},
				{
					Text:      strconv.Itoa(entry["Size (MB)"].(int)),
					Alignment: aligntable.AlignCenter,
				},
				{Text: entry["Created Date"].(string)},
			},
		})
	}

	fmt.Println(t)
	/*
		Image ID   Repository     Tag     Size (MB)       Created Date
		  987654   web-app       latest      350      2023-07-15 08:30
		  654321   database        v2        480      2023-06-28 14:15
		  789012   nginx         stable      180      2023-07-02 11:45
		  123789   api-service    prod       220      2023-06-15 18:20
	*/
}

func unmarshalAlignTableData() []map[string]any {
	var a []map[string]any
	err := yaml.Unmarshal([]byte(alignTableData), &a)
	if err != nil {
		panic(err)
	}
	return a
}
