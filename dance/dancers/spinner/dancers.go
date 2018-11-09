// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spinner

// ========================================================

// GetDancers returns all dancers defined in this package as a slice,
// which can be useful in benchmarking comparisons
func GetDancers() []Spinner {
	var c = make([]Spinner, 0, 5)
	c = append(c, Dance)
	return c
}

// DancerName returns the name of the dancer at index i of the result of GetDancers
func DancerName(i int) string {
	switch {
		case i == 0:	return "Dance"
		default:	return "<unknown>"
	}
}
