// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/tees/list"

	"fmt"
)

func ExampleList_AddBeam() {
	fmt.Println("Starting")

	// Create a new list and put some beams in it.
	var root = list.NewList("Board")
	root.AddBeam("Spalten", "A", "B", "C", "D", "E", "F", "G", "H")
	root.AddBeam("Zeilen", 1, 2, 3, 4, 5, 6, 7, 8)

	root.Print("root")
	root.PrintAtomValues("root")
	root.PrintAways()

	// Output:
	// Starting
	// root: List=Board | Total=2
	// root: List=Board	Board|Spalten | Board|Zeilen | Total=2
	// List=Board | Total=2
	// => : List=Spalten	A | B | C | D | E | F | G | H | Total=8
	// => : List=Zeilen	1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | Total=8
}
