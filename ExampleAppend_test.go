// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees_test

import (
	"github.com/GoLangsam/tees"
)

func ExampleAppend() {
	// Create a new list
	var atom = tees.New("Atom")

	// Add some Beams
	var ds = atom.AddBeam("Zahlen", 1, 2, 3)       // int
	var as = atom.AddBeam("Staben", "A", "B", "C") // char
	var ws = atom.AddBeam("Wörter", "foo", "bar")  // string
	var bs = atom.AddBeam("Boolen", true, false)   // bool

	atom.Print("Starting")
	atom.PrintAways("Atom")

	// So far, we have 3 * 3 * 2 * 2 Atom's
	// All else is going to be built from these atoms
	var dark = tees.New("Dark")

	var plus = dark.AddBeam("Append")
	plus.AddList(tees.Append(nil))
	plus.AddList(tees.Append(nil, ds))
	plus.AddList(tees.Append(ds))
	plus.AddList(tees.Append(ds, as))
	plus.AddList(tees.Append(ds, as, nil))
	plus.AddList(tees.Append(ds, as, bs))

	var mult = dark.AddBeam("Times")
	mult.AddList(tees.Times(nil))
	mult.AddList(tees.Times(nil, ds))
	mult.AddList(tees.Times(ds, nil))
	mult.AddList(tees.Times(ds))
	mult.AddList(tees.Times(ds, ws))
	mult.AddList(tees.Times(ds, ws, bs))
	mult.AddList(tees.Times(bs, ws, ds))

	dark.PrintAways("Dark")
	plus.PrintAways("Plus")
	mult.PrintAways("Mult")

	// Output:
	// Starting: List=Atom | Total=4
	// Atom: List=Atom | Total=4
	// => : List=Zahlen	1 | 2 | 3 | Total=3
	// => : List=Staben	A | B | C | Total=3
	// => : List=Wörter	foo | bar | Total=2
	// => : List=Boolen	true | false | Total=2
	// Dark: List=Dark | Total=2
	// => : List=Append	Append|<nil> | Append|<nil> | Append|Zahlen | Append|Zahlen|Staben | Append|Zahlen|Staben | Append|Zahlen|Staben|Boolen | Total=6
	// => : List=Times	Times|<nil> | Times|<nil> | Times|Zahlen | Times|Zahlen | Times|Zahlen|Wörter | Times|Zahlen|Wörter|Boolen | Times|Boolen|Wörter|Zahlen | Total=7
	// Plus: List=Append | Total=6
	// => : List=<nil>	Total=0
	// => : List=<nil>	Total=0
	// => : List=Zahlen	1 | 2 | 3 | Total=3
	// => : List=Zahlen|Staben	1 | 2 | 3 | A | B | C | Total=6
	// => : List=Zahlen|Staben	1 | 2 | 3 | A | B | C | Total=6
	// => : List=Zahlen|Staben|Boolen	1 | 2 | 3 | A | B | C | true | false | Total=8
	// Mult: List=Times | Total=7
	// => : List=<nil>	Total=0
	// => : List=<nil>	Total=0
	// => : List=Zahlen	Total=0
	// => : List=Zahlen	Total=0
	// => : List=Zahlen|Wörter	1|foo | 1|bar | 2|foo | 2|bar | 3|foo | 3|bar | Total=6
	// => : List=Zahlen|Wörter|Boolen	1|foo|true | 1|foo|false | 1|bar|true | 1|bar|false | 2|foo|true | 2|foo|false | 2|bar|true | 2|bar|false | 3|foo|true | 3|foo|false | 3|bar|true | 3|bar|false | Total=12
	// => : List=Boolen|Wörter|Zahlen	true|foo|1 | true|foo|2 | true|foo|3 | true|bar|1 | true|bar|2 | true|bar|3 | false|foo|1 | false|foo|2 | false|foo|3 | false|bar|1 | false|bar|2 | false|bar|3 | Total=12
}
