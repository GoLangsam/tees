// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinner

// ========================================================

// GetTurners returns all turners defined in this package as a slice,
// which can be useful in benchmarking comparisons
func GetTurners() []Turner {
	var c = make([]Turner, 0, 5)
	c = append(c, TurnFast)
	c = append(c, TurnEver)
	return c
}

// TurnerName returns the name of the turner at index i of the result of GetTurners
func TurnerName(i int) string {
	switch {
		case i == 0:	return "TurnFast"
		case i == 1:	return "TurnEver"
		default:	return "<unknown>"
	}
}
