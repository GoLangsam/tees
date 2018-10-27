// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"github.com/GoLangsam/tees/list"
)

// Kata (Japanese) is a form of movement (in martial arts)
type Kata []GoTo // Japanese

// ========================================================

// From returns the Element (or nil) reached from e by steps
func (steps Kata) From(e *list.Element) (*list.Element, Distance) {
	var dist, dnow Distance
	var goal = e

	for _, step := range steps {
		if goal == nil {
			return nil, dist
		}
		goal, dnow = step.from(goal)
		dist += dnow
	}
	return goal, dist
}

// ========================================================

// Grab returns all Elements reached from e by steps
// until nil or same is found
// Note: Grab may be useful in debugging, as it returns a full trace
// To Grab is not intended for regular use - Don't be greedy :-)
func (steps Kata) Grab(e *list.Element) ([]*list.Element, Distance) {
	var dist, dnow Distance
	var goal = e
	last := goal
	goals := make([]*list.Element, 0, len(steps))
	for _, step := range steps {
		last = goal
		goal, dnow = step.from(goal)
		if goal == nil || goal == last {
			continue
		}
		goals = append(goals, goal)
		dist += dnow
	}
	return goals, dist
}

// ========================================================

// Haul returns the Elements (or nil) From e by repeating steps
// until nil or seen is found
func (steps Kata) Haul(e *list.Element) ([]*list.Element, Distance) {
	var seen = make(map[*list.Element]bool)
	var dist, dnow Distance
	var goal = e
	goals := make([]*list.Element, 0, len(steps)*8)
	for {
		goal, dnow = steps.From(goal)
		if goal == nil || seen[goal] {
			return goals, dist
		}
		seen[goal] = true
		goals = append(goals, goal)
		dist += dnow
	}
}

// ========================================================
// Iterator

// Walker returns an iterator repeating Kata.From(e) ...
func (steps Kata) Walker(e *list.Element) Walk {
	var seen = make(map[*list.Element]bool)
	var curr = e
	var kata = steps

	var move Walk = func() *list.Element {
		seen[curr] = true
		if curr == nil {
			return nil
		}
		curr, _ = kata.From(curr)
		if seen[curr] {
			return nil
		} // already seen
		return curr
	}
	return move
}
