package main

import (
	"fmt"

	"github.com/semihbkgr/aligntable"
)

func main() {
	t := aligntable.Tree{}
	t.Nodes = []*aligntable.Node{
		{
			Text: "Fruits",
			SubNodes: []*aligntable.Node{
				{
					Text: "Citrus",
					SubNodes: []*aligntable.Node{
						{Text: "Orange"},
						{Text: "Lemon"},
						{Text: "Lime"},
					},
				},
				{
					Text: "Berries",
					SubNodes: []*aligntable.Node{
						{Text: "Strawberry"},
						{Text: "Lemon"},
						{Text: "Blueberry"},
					},
				},
				{
					Text: "Tropical",
					SubNodes: []*aligntable.Node{
						{Text: "Mango"},
						{Text: "Pineapple"},
						{Text: "Blueberry"},
					},
				},
			},
		}, {
			Text: "Vegetables",
			SubNodes: []*aligntable.Node{
				{
					Text: "Leafy Greens",
					SubNodes: []*aligntable.Node{
						{Text: "Lettuce"},
						{Text: "Spinach"},
					},
				},
				{
					Text: "Root Vegetables",
					SubNodes: []*aligntable.Node{
						{Text: "Carrot"},
						{Text: "Potato"},
						{Text: "Beetroot"},
					},
				},
			},
		},
	}

	fmt.Println(t.Table())

	/*
		Fruits
		├───────── Citrus
		│          ├────────────── Orange
		│          ├────────────── Lemon
		│          └────────────── Lime
		├───────── Berries
		│          ├────────────── Strawberry
		│          ├────────────── Lemon
		│          └────────────── Blueberry
		└───────── Tropical
		           ├────────────── Mango
		           ├────────────── Pineapple
		           └────────────── Blueberry
		Vegetables
		├───────── Leafy Greens
		│          ├────────────── Lettuce
		│          └────────────── Spinach
		└───────── Root Vegetables
		           ├────────────── Carrot
		           ├────────────── Potato
		           └────────────── Beetroot
	*/
}
