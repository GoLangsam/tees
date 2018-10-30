// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
chooser.go defines some Choose-functions to choose sth for a dancer to dance on

*/
package chooser

import (
	"github.com/GoLangsam/tees"
)

// ========================================================

// ChooseNil returns nil & false - useful for abort
func ChooseNil(l *tees.Tees) (*tees.Tees, bool) {
	return nil, false
}

// ========================================================

// ChooseFront simply returns the first column
func ChooseFront(l *tees.Tees) (*tees.Tees, bool) {

	if l == nil { panic( "List to choose from is nil!") }
	if Verbose {	l.PrintAways("Choosing from: ")	}

	e := l.Front()
	if e == nil { return nil, false }
	e = e.Away()
	if e == nil { return nil, false }

	c := e.List()
	if c == nil { return nil, false }
	if Verbose { c.PrintAways("Chosen: ") }
	return c, c.Len() > 0
}

// ========================================================

// ChooseShort returns the shortest column
//	- it's the preferred method of choice
//	- it returns false on zero-lenght as there is no point in trying such
func ChooseShort(l *tees.Tees) (*tees.Tees, bool) {
	var found bool = false
	var c *tees.Tees
	s := 999999999
	for j := l.Front(); j != nil; j = j.Next() {
		list := j.AwayList()
		if list != nil {
			leng := list.Len()
			if leng < s {
				c = list
				s = leng
				found = true
				if s == 0 {
					break
				}
			}
		}
	}
	if Verbose { c.PrintAways("Chosen: ") }
	return c, found && c.Len() > 0
}

// Note: return ChooseUpto(list, 0) is equivalent, but ChooseUpto0 would cost another function call

// ========================================================

// ChooseUpto0 returns the first column not longer than 0
func ChooseUpto0(list *tees.Tees) (*tees.Tees, bool) {
	return ChooseUpto(list, 0)
}

// ChooseUpto1 returns the first column not longer than 1
func ChooseUpto1(list *tees.Tees) (*tees.Tees, bool) {
	return ChooseUpto(list, 1)
}

// ChooseUpto1 returns the first column not longer than 2
func ChooseUpto2(list *tees.Tees) (*tees.Tees, bool) {
	return ChooseUpto(list, 2)
}

// ChooseUpto1 returns the first column not longer than 3
func ChooseUpto3(list *tees.Tees) (*tees.Tees, bool) {
	return ChooseUpto(list, 3)
}

// ========================================================

// ChooseBelow returns the first column not longer than min
func ChooseUpto(l *tees.Tees, min int) (*tees.Tees, bool) {
	var found bool = false
	var c *tees.Tees
	s := 999999999
	for j := l.Front(); j != nil; j = j.Next() {
		list := j.AwayList()
		if list != nil {
			leng := list.Len()
			if leng < s {
				c = list
				s = leng
				found = true
				if s <= min {
					break
				}

			}
		}
	}
	if Verbose { c.PrintAways("Chosen: ") }
	return c, found && c.Len() > 0
}

// ========================================================

// TODO: ChooseOrgan also considers the total length of the aways (=rows) - the higher the better

/*
	currlist curraway
	looklist lookaway

	d := looklist - currlist
	a := lookaway - curraway

	if a > d * weight { curr = look }

	switch d {
		d < 0: if a < 0 {if a < d	{ curr = look } }	// look is shorter; but should not have much less aways (not a>=0 or
		d = 0: if a > 0			{ curr = look }		// look has same len; look has more aways
		d > 0: if a > d * weight	{ curr = look }		// look is longer; look has weight more aways

	}


	if looklist <= currlist

	> currlist + curraway

*/
