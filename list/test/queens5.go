// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/GoLangsam/tees/list"
)

// FiveQueens
func FiveQueens() *list.List {
	var cols = list.NewList("Board")

	// Ranks
	var R0 = cols.AddBeam("R0")
	var R1 = cols.AddBeam("R1")
	var R2 = cols.AddBeam("R2")
	var R3 = cols.AddBeam("R3")
	var R4 = cols.AddBeam("R4")
	// Files
	var F0 = cols.AddBeam("F0")
	var F1 = cols.AddBeam("F1")
	var F2 = cols.AddBeam("F2")
	var F3 = cols.AddBeam("F3")
	var F4 = cols.AddBeam("F4")

	var diag = list.NewList("Diagonals")
	// Diagonals
	var A0 = diag.AddBeam("A0")
	var A1 = diag.AddBeam("A1")
	var A2 = diag.AddBeam("A2")
	var A3 = diag.AddBeam("A3")
	var A4 = diag.AddBeam("A4")
	var A5 = diag.AddBeam("A5")
	var A6 = diag.AddBeam("A6")
	var A7 = diag.AddBeam("A7")
	var A8 = diag.AddBeam("A8")

	var scnd = list.NewList("secondary")
	var B0 = scnd.AddBeam("A0")
	var B1 = scnd.AddBeam("B1")
	var B2 = scnd.AddBeam("B2")
	var B3 = scnd.AddBeam("B3")
	var B4 = scnd.AddBeam("B4")
	var B5 = scnd.AddBeam("B5")
	var B6 = scnd.AddBeam("B6")
	var B7 = scnd.AddBeam("B7")
	var B8 = scnd.AddBeam("B8")
	diag.Junk(scnd)

	var rows = list.NewList("Rows")

	// TODO: AddOnes vs AddJunk
	rows.AddJunk("R-00", R0, F0, A0, B3)
	rows.AddJunk("R-01", R0, F1, A1, B4)
	rows.AddJunk("R-02", R0, F2, A2, B5)
	rows.AddJunk("R-03", R0, F3, A3, B6)
	rows.AddJunk("R-04", R0, F4, A4, B7)

	rows.AddJunk("R-10", R1, F0, A1, B2)
	rows.AddJunk("R-11", R1, F1, A2, B3)
	rows.AddJunk("R-12", R1, F2, A3, B4)
	rows.AddJunk("R-13", R1, F3, A4, B5)
	rows.AddJunk("R-14", R1, F4, A5, B6)

	rows.AddJunk("R-20", R2, F0, A2, B1)
	rows.AddJunk("R-21", R2, F1, A3, B2)
	rows.AddJunk("R-22", R2, F2, A4, B3)
	rows.AddJunk("R-23", R2, F3, A5, B4)
	rows.AddJunk("R-24", R2, F4, A6, B5)

	rows.AddJunk("R-30", R3, F0, A3, B0)
	rows.AddJunk("R-31", R3, F1, A4, B1)
	rows.AddJunk("R-32", R3, F2, A5, B2)
	rows.AddJunk("R-33", R3, F3, A6, B3)
	rows.AddJunk("R-34", R3, F4, A7, B4)

	rows.AddJunk("R-40", R4, F0, A4, B8)
	rows.AddJunk("R-41", R4, F1, A5, B0)
	rows.AddJunk("R-42", R4, F2, A6, B1)
	rows.AddJunk("R-43", R4, F3, A7, B2)
	rows.AddJunk("R-44", R4, F4, A8, B3)
	cols.Junk(rows)
	return cols
}
