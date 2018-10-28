// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"sort"
)

// Arrange sorts t in place according to weight.
// The sort is stable.
func Arrange(t Tees, weight func(This) int) Tees {
	if t.Len() < 1 {
		return t
	}

	these := t.Elements()
	less := func(i, j int) bool { return weight(these[i]) < weight(these[j]) }
	return Stable(these, less)
}

// Stable sorts these in place according to less.
// The sort is stable.
func Stable(these []This, less func(i, j int) bool) Tees {
	if len(these) < 1 {
		return nil
	}

	mark := these[0]
	sort.SliceStable(these, less)

	for _, e := range these {
		mark = e.MoveToPrevOf(mark)
	}
	return mark.List()
}

// Weight returns both:
//  a map of the elements with their respective weight and
//  a map of the weights with their elements
func Weight(t trailer, weight func(This) int) (map[This]int, map[int][]This) {
	var elems = make(map[This]int, t.Len())
	var sizes = make(map[int][]This, t.Len())

	var size int
	for e := t.Front(); e != nil; e = e.Next() {
		size = weight(e)
		elems[e] = size
		sizes[size] = append(sizes[size], e)
	}
	return elems, sizes
}
