// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/GoLangsam/tees/list"
)

// SmallMatrix - Donald E. Knuth's sample
func SmallMatrix() *list.List {
	// Create a new list and put some beams in it.
	var cols = list.NewList("Cols")
	var C1 = cols.AddBeam("A")
	var C2 = cols.AddBeam("B")
	var C3 = cols.AddBeam("C")
	var C4 = cols.AddBeam("D")
	var C5 = cols.AddBeam("E")
	var C6 = cols.AddBeam("F")
	var C7 = cols.AddBeam("G")

	/*
		A B C D E F G

		0 0 1 0 1 1 0
		1 0 0 1 0 0 1
		0 1 1 0 0 1 0
		1 0 0 1 0 0 0
		0 1 0 0 0 0 1
		0 0 0 1 1 0 1

	*/

	var rows = list.NewList("Rows")
	rows.Root().Junk(cols.Root())

	rows = list.NewList("Dis-Disorderd Rows")

	// TODO: AddOnes vs AddJunk
	rows.AddJunk("R-4", C1, C4)
	rows.AddJunk("R-5", C7, C2)
	rows.AddJunk("R-1", C6, C3, C5)
	rows.AddJunk("R-2", C7, C4, C1)
	rows.AddJunk("R-3", C6, C2, C3)
	rows.AddJunk("R-6", C7, C5, C4)

	return cols
}
