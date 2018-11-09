// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
choosers.go provides the Choosers, and a poor-mans lookup of their names.

TODO(AP): Redo this stupid implemntation

*/
package chooser

// ========================================================

// GetChoosers returns all choosers defined in this package as a slice,
// which can be useful in benchmarking comparisons
func GetChoosers() []Chooser {
	var c = make([]Chooser, 0, 5)
	c = append(c, ChooseUpto3)
	c = append(c, ChooseUpto2)
	c = append(c, ChooseUpto1)
	c = append(c, ChooseUpto0)
	c = append(c, ChooseShort)
	c = append(c, ChooseFront)
	return c
}

// ChooserName returns the name of the chooser at index i of the result of GetChoosers
func ChooserName(i int) string {
	switch {
		case i == 0:	return "ChooseUpto3"
		case i == 1:	return "ChooseUpto2"
		case i == 2:	return "ChooseUpto1"
		case i == 3:	return "ChooseUpto0"
		case i == 4:	return "ChooseShort"
		case i == 5:	return "ChooseFront"
		default:	return "<unknown>"
	}
}
