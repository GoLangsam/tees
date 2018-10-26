// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// TrixSizes returns both: a map of the elements with their respective TrixSize and a map of the Sizes with their elements
func TrixSizes(l *list.List) (map[*list.Element]int, map[int][]*list.Element) {
	var elems = make(map[*list.Element]int, l.Len())
	var sizes = make(map[int][]*list.Element, l.Len())

	var size int
	for e := l.Front(); e != nil; e = e.Next() {
		size = TrixSize(e)
		elems[e] = size
		sizes[size] = append(sizes[size], e)
	}
	return elems, sizes
}

// TrixSize returns the sum of AwaySize of e's AwayList
func TrixSize(e *list.Element) int {
	size := 0
	for f := e.AwayList().Front(); f != nil; f = f.Next() {
		size += AwaySize(f.AwayList()) // same as f.AwayList().Size()
	}
	return size
}

// AwaySize returns the sum of Len() of the AwayLists along l - same as l.Size()
func AwaySize(l *list.List) int {
	size := 0
	for e := l.Front(); e != nil; e = e.Next() {
		size += e.AwayList().Len()
	}
	return size
}
