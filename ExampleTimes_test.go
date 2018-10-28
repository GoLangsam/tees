// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees_test

import (
	"github.com/GoLangsam/tees"
)

func ExampleTimes() {
	// Create a new list
	var chess = tees.New("Schach")

	// Add some Beams
	var cols = chess.AddBeam("Spalten", "A", "B", "C", "D", "E", "F", "G", "H")
	var rows = chess.AddBeam("Zeilen", 1, 2, 3, 4, 5, 6, 7, 8)
	var game = chess.AddBeam("Partien", "Fischer / Kasparski", "John / Doe", "Foo / Bar")

	chess.Print("Starting")
	chess.PrintAtomValues("chess")
	chess.PrintAways("chess")

	// So far, we have 8 * 8 * 3 Atom's
	// All else is going to be built from these atoms
	rcg := tees.Times(rows, cols, game)
	rcg.Root().PrintAtomValues("rows.Times( cols, game )")
	chess.AddList(rcg)

	chess.AddList(tees.Times(game, tees.Times(cols, rows)))

	chess.Print("Growing")
	chess.PrintAtomValues("chess")
	chess.PrintAways("chess")

	// Output:
	// Starting: List=Schach | Total=3
	// chess: List=Schach	Schach|Spalten | Schach|Zeilen | Schach|Partien | Total=3
	// chess: List=Schach | Total=3
	// => : List=Spalten	A | B | C | D | E | F | G | H | Total=8
	// => : List=Zeilen	1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | Total=8
	// => : List=Partien	Fischer / Kasparski | John / Doe | Foo / Bar | Total=3
	// rows.Times( cols, game ): Element=Zeilen|Spalten|Partien.
	// Growing: List=Schach | Total=5
	// chess: List=Schach	Schach|Spalten | Schach|Zeilen | Schach|Partien | Schach|Zeilen|Spalten|Partien | Schach|Partien|Spalten|Zeilen | Total=5
	// chess: List=Schach | Total=5
	// => : List=Spalten	A | B | C | D | E | F | G | H | Total=8
	// => : List=Zeilen	1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | Total=8
	// => : List=Partien	Fischer / Kasparski | John / Doe | Foo / Bar | Total=3
	// => : List=Zeilen|Spalten|Partien	1|A|Fischer / Kasparski | 1|A|John / Doe | 1|A|Foo / Bar | 1|B|Fischer / Kasparski | 1|B|John / Doe | 1|B|Foo / Bar | 1|C|Fischer / Kasparski | 1|C|John / Doe | 1|C|Foo / Bar | 1|D|Fischer / Kasparski | 1|D|John / Doe | 1|D|Foo / Bar | 1|E|Fischer / Kasparski | 1|E|John / Doe | 1|E|Foo / Bar | 1|F|Fischer / Kasparski | 1|F|John / Doe | 1|F|Foo / Bar | 1|G|Fischer / Kasparski | 1|G|John / Doe | 1|G|Foo / Bar | 1|H|Fischer / Kasparski | 1|H|John / Doe | 1|H|Foo / Bar | 2|A|Fischer / Kasparski | 2|A|John / Doe | 2|A|Foo / Bar | 2|B|Fischer / Kasparski | 2|B|John / Doe | 2|B|Foo / Bar | 2|C|Fischer / Kasparski | 2|C|John / Doe | 2|C|Foo / Bar | 2|D|Fischer / Kasparski | 2|D|John / Doe | 2|D|Foo / Bar | 2|E|Fischer / Kasparski | 2|E|John / Doe | 2|E|Foo / Bar | 2|F|Fischer / Kasparski | 2|F|John / Doe | 2|F|Foo / Bar | 2|G|Fischer / Kasparski | 2|G|John / Doe | 2|G|Foo / Bar | 2|H|Fischer / Kasparski | 2|H|John / Doe | 2|H|Foo / Bar | 3|A|Fischer / Kasparski | 3|A|John / Doe | 3|A|Foo / Bar | 3|B|Fischer / Kasparski | 3|B|John / Doe | 3|B|Foo / Bar | 3|C|Fischer / Kasparski | 3|C|John / Doe | 3|C|Foo / Bar | 3|D|Fischer / Kasparski | 3|D|John / Doe | 3|D|Foo / Bar | 3|E|Fischer / Kasparski | 3|E|John / Doe | 3|E|Foo / Bar | 3|F|Fischer / Kasparski | 3|F|John / Doe | 3|F|Foo / Bar | 3|G|Fischer / Kasparski | 3|G|John / Doe | 3|G|Foo / Bar | 3|H|Fischer / Kasparski | 3|H|John / Doe | 3|H|Foo / Bar | 4|A|Fischer / Kasparski | 4|A|John / Doe | 4|A|Foo / Bar | 4|B|Fischer / Kasparski | 4|B|John / Doe | 4|B|Foo / Bar | 4|C|Fischer / Kasparski | 4|C|John / Doe | 4|C|Foo / Bar | 4|D|Fischer / Kasparski | 4|D|John / Doe | 4|D|Foo / Bar | 4|E|Fischer / Kasparski | 4|E|John / Doe | 4|E|Foo / Bar | 4|F|Fischer / Kasparski | 4|F|John / Doe | 4|F|Foo / Bar | 4|G|Fischer / Kasparski | 4|G|John / Doe | 4|G|Foo / Bar | 4|H|Fischer / Kasparski | 4|H|John / Doe | 4|H|Foo / Bar | 5|A|Fischer / Kasparski | 5|A|John / Doe | 5|A|Foo / Bar | 5|B|Fischer / Kasparski | 5|B|John / Doe | 5|B|Foo / Bar | 5|C|Fischer / Kasparski | 5|C|John / Doe | 5|C|Foo / Bar | 5|D|Fischer / Kasparski | 5|D|John / Doe | 5|D|Foo / Bar | 5|E|Fischer / Kasparski | 5|E|John / Doe | 5|E|Foo / Bar | 5|F|Fischer / Kasparski | 5|F|John / Doe | 5|F|Foo / Bar | 5|G|Fischer / Kasparski | 5|G|John / Doe | 5|G|Foo / Bar | 5|H|Fischer / Kasparski | 5|H|John / Doe | 5|H|Foo / Bar | 6|A|Fischer / Kasparski | 6|A|John / Doe | 6|A|Foo / Bar | 6|B|Fischer / Kasparski | 6|B|John / Doe | 6|B|Foo / Bar | 6|C|Fischer / Kasparski | 6|C|John / Doe | 6|C|Foo / Bar | 6|D|Fischer / Kasparski | 6|D|John / Doe | 6|D|Foo / Bar | 6|E|Fischer / Kasparski | 6|E|John / Doe | 6|E|Foo / Bar | 6|F|Fischer / Kasparski | 6|F|John / Doe | 6|F|Foo / Bar | 6|G|Fischer / Kasparski | 6|G|John / Doe | 6|G|Foo / Bar | 6|H|Fischer / Kasparski | 6|H|John / Doe | 6|H|Foo / Bar | 7|A|Fischer / Kasparski | 7|A|John / Doe | 7|A|Foo / Bar | 7|B|Fischer / Kasparski | 7|B|John / Doe | 7|B|Foo / Bar | 7|C|Fischer / Kasparski | 7|C|John / Doe | 7|C|Foo / Bar | 7|D|Fischer / Kasparski | 7|D|John / Doe | 7|D|Foo / Bar | 7|E|Fischer / Kasparski | 7|E|John / Doe | 7|E|Foo / Bar | 7|F|Fischer / Kasparski | 7|F|John / Doe | 7|F|Foo / Bar | 7|G|Fischer / Kasparski | 7|G|John / Doe | 7|G|Foo / Bar | 7|H|Fischer / Kasparski | 7|H|John / Doe | 7|H|Foo / Bar | 8|A|Fischer / Kasparski | 8|A|John / Doe | 8|A|Foo / Bar | 8|B|Fischer / Kasparski | 8|B|John / Doe | 8|B|Foo / Bar | 8|C|Fischer / Kasparski | 8|C|John / Doe | 8|C|Foo / Bar | 8|D|Fischer / Kasparski | 8|D|John / Doe | 8|D|Foo / Bar | 8|E|Fischer / Kasparski | 8|E|John / Doe | 8|E|Foo / Bar | 8|F|Fischer / Kasparski | 8|F|John / Doe | 8|F|Foo / Bar | 8|G|Fischer / Kasparski | 8|G|John / Doe | 8|G|Foo / Bar | 8|H|Fischer / Kasparski | 8|H|John / Doe | 8|H|Foo / Bar | Total=192
	// => : List=Partien|Spalten|Zeilen	Fischer / Kasparski|A|1 | Fischer / Kasparski|A|2 | Fischer / Kasparski|A|3 | Fischer / Kasparski|A|4 | Fischer / Kasparski|A|5 | Fischer / Kasparski|A|6 | Fischer / Kasparski|A|7 | Fischer / Kasparski|A|8 | Fischer / Kasparski|B|1 | Fischer / Kasparski|B|2 | Fischer / Kasparski|B|3 | Fischer / Kasparski|B|4 | Fischer / Kasparski|B|5 | Fischer / Kasparski|B|6 | Fischer / Kasparski|B|7 | Fischer / Kasparski|B|8 | Fischer / Kasparski|C|1 | Fischer / Kasparski|C|2 | Fischer / Kasparski|C|3 | Fischer / Kasparski|C|4 | Fischer / Kasparski|C|5 | Fischer / Kasparski|C|6 | Fischer / Kasparski|C|7 | Fischer / Kasparski|C|8 | Fischer / Kasparski|D|1 | Fischer / Kasparski|D|2 | Fischer / Kasparski|D|3 | Fischer / Kasparski|D|4 | Fischer / Kasparski|D|5 | Fischer / Kasparski|D|6 | Fischer / Kasparski|D|7 | Fischer / Kasparski|D|8 | Fischer / Kasparski|E|1 | Fischer / Kasparski|E|2 | Fischer / Kasparski|E|3 | Fischer / Kasparski|E|4 | Fischer / Kasparski|E|5 | Fischer / Kasparski|E|6 | Fischer / Kasparski|E|7 | Fischer / Kasparski|E|8 | Fischer / Kasparski|F|1 | Fischer / Kasparski|F|2 | Fischer / Kasparski|F|3 | Fischer / Kasparski|F|4 | Fischer / Kasparski|F|5 | Fischer / Kasparski|F|6 | Fischer / Kasparski|F|7 | Fischer / Kasparski|F|8 | Fischer / Kasparski|G|1 | Fischer / Kasparski|G|2 | Fischer / Kasparski|G|3 | Fischer / Kasparski|G|4 | Fischer / Kasparski|G|5 | Fischer / Kasparski|G|6 | Fischer / Kasparski|G|7 | Fischer / Kasparski|G|8 | Fischer / Kasparski|H|1 | Fischer / Kasparski|H|2 | Fischer / Kasparski|H|3 | Fischer / Kasparski|H|4 | Fischer / Kasparski|H|5 | Fischer / Kasparski|H|6 | Fischer / Kasparski|H|7 | Fischer / Kasparski|H|8 | John / Doe|A|1 | John / Doe|A|2 | John / Doe|A|3 | John / Doe|A|4 | John / Doe|A|5 | John / Doe|A|6 | John / Doe|A|7 | John / Doe|A|8 | John / Doe|B|1 | John / Doe|B|2 | John / Doe|B|3 | John / Doe|B|4 | John / Doe|B|5 | John / Doe|B|6 | John / Doe|B|7 | John / Doe|B|8 | John / Doe|C|1 | John / Doe|C|2 | John / Doe|C|3 | John / Doe|C|4 | John / Doe|C|5 | John / Doe|C|6 | John / Doe|C|7 | John / Doe|C|8 | John / Doe|D|1 | John / Doe|D|2 | John / Doe|D|3 | John / Doe|D|4 | John / Doe|D|5 | John / Doe|D|6 | John / Doe|D|7 | John / Doe|D|8 | John / Doe|E|1 | John / Doe|E|2 | John / Doe|E|3 | John / Doe|E|4 | John / Doe|E|5 | John / Doe|E|6 | John / Doe|E|7 | John / Doe|E|8 | John / Doe|F|1 | John / Doe|F|2 | John / Doe|F|3 | John / Doe|F|4 | John / Doe|F|5 | John / Doe|F|6 | John / Doe|F|7 | John / Doe|F|8 | John / Doe|G|1 | John / Doe|G|2 | John / Doe|G|3 | John / Doe|G|4 | John / Doe|G|5 | John / Doe|G|6 | John / Doe|G|7 | John / Doe|G|8 | John / Doe|H|1 | John / Doe|H|2 | John / Doe|H|3 | John / Doe|H|4 | John / Doe|H|5 | John / Doe|H|6 | John / Doe|H|7 | John / Doe|H|8 | Foo / Bar|A|1 | Foo / Bar|A|2 | Foo / Bar|A|3 | Foo / Bar|A|4 | Foo / Bar|A|5 | Foo / Bar|A|6 | Foo / Bar|A|7 | Foo / Bar|A|8 | Foo / Bar|B|1 | Foo / Bar|B|2 | Foo / Bar|B|3 | Foo / Bar|B|4 | Foo / Bar|B|5 | Foo / Bar|B|6 | Foo / Bar|B|7 | Foo / Bar|B|8 | Foo / Bar|C|1 | Foo / Bar|C|2 | Foo / Bar|C|3 | Foo / Bar|C|4 | Foo / Bar|C|5 | Foo / Bar|C|6 | Foo / Bar|C|7 | Foo / Bar|C|8 | Foo / Bar|D|1 | Foo / Bar|D|2 | Foo / Bar|D|3 | Foo / Bar|D|4 | Foo / Bar|D|5 | Foo / Bar|D|6 | Foo / Bar|D|7 | Foo / Bar|D|8 | Foo / Bar|E|1 | Foo / Bar|E|2 | Foo / Bar|E|3 | Foo / Bar|E|4 | Foo / Bar|E|5 | Foo / Bar|E|6 | Foo / Bar|E|7 | Foo / Bar|E|8 | Foo / Bar|F|1 | Foo / Bar|F|2 | Foo / Bar|F|3 | Foo / Bar|F|4 | Foo / Bar|F|5 | Foo / Bar|F|6 | Foo / Bar|F|7 | Foo / Bar|F|8 | Foo / Bar|G|1 | Foo / Bar|G|2 | Foo / Bar|G|3 | Foo / Bar|G|4 | Foo / Bar|G|5 | Foo / Bar|G|6 | Foo / Bar|G|7 | Foo / Bar|G|8 | Foo / Bar|H|1 | Foo / Bar|H|2 | Foo / Bar|H|3 | Foo / Bar|H|4 | Foo / Bar|H|5 | Foo / Bar|H|6 | Foo / Bar|H|7 | Foo / Bar|H|8 | Total=192

}
