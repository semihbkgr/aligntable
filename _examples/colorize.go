package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/semihbkgr/aligntable"
	"gopkg.in/yaml.v3"
)

const colorizeTableData = `
- Resource: Frontend-Pod
  Kind: Pod
  Namespace: default
  CPU Requests: 0.5 core
  Memory Requests: 512 MiB
  CPU Limits: 1 core
  Memory Limits: 1 GiB

- Resource: Backend-Deploy
  Kind: Deployment
  Namespace: backend
  CPU Requests: 1 core
  Memory Requests: 1 GiB
  CPU Limits: 2 core
  Memory Limits: 2 GiB

- Resource: Database-Pod
  Kind: Pod
  Namespace: database
  CPU Requests: 0.25 core
  Memory Requests: 256 MiB
  CPU Limits: 0.5 core
  Memory Limits: 512 MiB

- Resource: Worker-Replica
  Kind: ReplicaSet
  Namespace: workers
  CPU Requests: 0.5 core
  Memory Requests: 512 MiB
  CPU Limits: 1 core
  Memory Limits: 1 GiB

- Resource: API-Service
  Kind: Service
  Namespace: default
  CPU Requests: "-"
  Memory Requests: "-"
  CPU Limits: "-"
  Memory Limits: "-"
`

func main() {
	tree()
}

func tree() {
	t := aligntable.New()
	t.Rows = []*aligntable.Row{
		{
			Cells: []*aligntable.Cell{
				{
					Text:      color.HiBlueString("RESOURCE"),
					Alignment: aligntable.AlignLeft,
				},
				{
					Text:      color.HiBlueString("KIND"),
					Alignment: aligntable.AlignLeft,
				},
				{
					Text:      color.HiBlueString("NAMESPACE"),
					Alignment: aligntable.AlignCenter,
				},
				{
					Text:      color.HiBlueString("CPU REQUESTS"),
					Alignment: aligntable.AlignCenter,
				},
				{
					Text:      color.HiBlueString("MEMORY REQUESTS"),
					Alignment: aligntable.AlignCenter,
				},
				{
					Text:      color.HiBlueString("CPU LIMITS"),
					Alignment: aligntable.AlignCenter,
				},
				{
					Text:      color.HiBlueString("MEMORY LIMITS"),
					Alignment: aligntable.AlignCenter,
				},
			},
		},
	}

	for _, entry := range unmarshalColorizeTableData() {
		t.Rows = append(t.Rows, &aligntable.Row{
			Cells: []*aligntable.Cell{
				{
					Text:      color.HiWhiteString(entry["Resource"].(string)),
					Alignment: aligntable.AlignLeft,
				},
				{
					Text:      color.HiCyanString(entry["Kind"].(string)),
					Alignment: aligntable.AlignLeft,
				},
				{
					Text:      color.HiBlackString(entry["Namespace"].(string)),
					Alignment: aligntable.AlignCenter,
				},
				{
					Text:      color.HiYellowString(entry["CPU Requests"].(string)),
					Alignment: aligntable.AlignRight,
				},
				{
					Text:      color.HiGreenString(entry["Memory Requests"].(string)),
					Alignment: aligntable.AlignRight,
				},
				{
					Text:      color.HiMagentaString(entry["CPU Limits"].(string)),
					Alignment: aligntable.AlignRight,
				},
				{
					Text:      color.HiRedString(entry["Memory Limits"].(string)),
					Alignment: aligntable.AlignRight,
				},
			},
		})
	}

	fmt.Println(t)
	/*
		RESOURCE         KIND         NAMESPACE   CPU REQUESTS   MEMORY REQUESTS   CPU LIMITS   MEMORY LIMITS
		Frontend-Pod     Pod           default        0.5 core           512 MiB       1 core           1 GiB
		Backend-Deploy   Deployment    backend          1 core             1 GiB       2 core           2 GiB
		Database-Pod     Pod          database       0.25 core           256 MiB     0.5 core         512 MiB
		Worker-Replica   ReplicaSet    workers        0.5 core           512 MiB       1 core           1 GiB
		API-Service      Service       default               -                 -            -               -
	*/
}

func unmarshalColorizeTableData() []map[string]any {
	var a []map[string]any
	err := yaml.Unmarshal([]byte(colorizeTableData), &a)
	if err != nil {
		panic(err)
	}
	return a
}
