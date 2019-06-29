// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"sort"
)

// Arrange sorts t in place according to weight.
// The sort is stable.
func Arrange(l slicer, weight func(*aTee) int) *Tees {
	if l.Len() < 1 {
		return l.(*Tees)
	}

	these := l.Elements()
	less := func(i, j int) bool { return weight(these[i]) < weight(these[j]) }
	return Stable(these, less)
}

// Stable sorts these in place according to less.
// The sort is stable.
func Stable(these []*aTee, less func(i, j int) bool) *Tees {
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
func Weight(l trailer, weight func(*aTee) int) (map[*aTee]int, map[int][]*aTee) {
	var elems = make(map[*aTee]int, l.Len())
	var sizes = make(map[int][]*aTee, l.Len())

	var size int
	for e := l.Front(); e != nil; e = e.Next() {
		size = weight(e)
		elems[e] = size
		sizes[size] = append(sizes[size], e)
	}
	return elems, sizes
}
