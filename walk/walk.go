// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

"GoTo" is an alias for a kind of function to reach another existing Here (or nil) from some Here e

"From" returns the Here (or nil) reached from e by applying some Kata's (=[]GoTo)

"Grab" is greedy and returns the slice of all Here reached from e

"Haul" is steady and returns the slice of Here reached from e keeping going until seen or nil

"Walker" returns an hauling iterator

TODO: "Turner" returns an hauling iterator roundrobbing it's akas
*/

package walk

// GoTo's represent basic functions to Step from some element to another existing element (or nil)
type GoTo uint8 // func( from Here ) (to Here)

// Distance represents the 'length' of some movement
// for a Kata it is the sum of it's GoTo distances
// for a Atas it is the 'radius', the max of it's Kata's distances
type Distance int

const (
	Next GoTo = iota
	Prev
	Away
	Root
	Front
	Back
	Home
)

// Walk represents an iterator. Usage pattern:
//
//  next := kata.Walker(e)
//  for e := next(), e != nil, e = next() { /* ... */ }
type Walk func() Here

// ========================================================
// from returns the Here (or nil) reached from e by applying GoTo
func (g GoTo) from(e Here) (Here, Distance) {

	var dist Distance = 0
	var next Here = e

	if e != nil {
		switch g {
		case Next:
			next = e.Next()
			if next != nil {
				dist = 1
			}
		case Prev:
			next = e.Prev()
			if next != nil {
				dist = 1
			}
		case Away:
			next = e.Away()

		case Root:
			next = e.Root()
			if next != nil {
				dist = Distance(e.List().Len())
			}

		case Front:
			next = e.Front()
			if next != nil {
				dist = Distance(e.List().Len())
			}
		case Back:
			next = e.Back()
			if next != nil {
				dist = Distance(e.List().Len())
			}

		case Home:
			next = e.Home()
			if next != nil {
				dist = Distance(e.List().Len())
			} // Homerun is long, but free :-)

		default:
			panic("Undefined value for GoTo")
		}
	}
	return next, dist
}
