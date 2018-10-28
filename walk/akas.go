// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"github.com/GoLangsam/tees/list"
)

// Akas (Burmese) is a slice of forms of movement (in martial arts, and dances :-) )
type Akas []Kata // Burmese

// ========================================================

// From returns the non-nil Elements reached from e by jumps
func (jumps Akas) From(e *list.Element) ([]*list.Element, Distance) {
	var dist, dnow Distance
	var goal = e
	goals := make([]*list.Element, 0, len(jumps))
	for _, steps := range jumps {
		goal, dnow = steps.From(e)
		if goal == nil {
			continue
		}
		goals = append(goals, goal)
		dist += dnow
	}
	return goals, dist
}

// ========================================================

// Grab returns all Elements reached from e by jumps
// Note: Grab may be useful in debugging, as it returns a full trace
// To Grab is not intended for regular use - Don't be greedy :-)
func (jumps Akas) Grab(e *list.Element) ([]*list.Element, Distance) {
	var dist Distance
	goals := make([]*list.Element, 0, len(jumps)*len(jumps))
	for _, steps := range jumps {
		goal, dnow := steps.Grab(e)
		if goal == nil || len(goal) == 0 {
			continue
		}
		goals = append(goals, goal...)
		dist += dnow
	}
	return goals, dist
}

// ========================================================

// Haul returns the Elements (or nil) From e by hauling jumps
// Note: From any new goal, just the current Kata is repeated!
// Not all jumps are done again - this would imply loops.
func (jumps Akas) Haul(e *list.Element) ([]*list.Element, Distance) {
	var dist Distance
	goals := make([]*list.Element, 0, len(jumps)*len(jumps)*8)
	for _, steps := range jumps {
		goal, dnow := steps.Haul(e)
		if goal == nil || len(goal) == 0 {
			continue
		}
		goals = append(goals, goal...)
		dist += dnow
	}
	return goals, dist
}

// ========================================================
// Iterator

// Walker returns an iterator walking all Kata.From(e) ...
func (jumps Akas) Walker(e *list.Element) Walk {

	var curr = e
	var akas = jumps
	var maxi = len(akas) - 1
	var aidx int // index of akas
	var next = akas[aidx].Walker(curr)

	var move Walk = func() *list.Element {
	next:
		goal := next()
		if goal == nil && aidx < maxi {
			aidx++
			next = akas[aidx].Walker(curr)
			goto next
		}
		return goal
	}
	return move
}
