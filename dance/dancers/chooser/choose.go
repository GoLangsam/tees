// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


// Package chooser defines some Choose-functions to choose sth for a dancer to dance on
package chooser

// ========================================================

// ChooseNil returns nil & false - useful for abort
func ChooseNil(among *Tees) (*Tees, bool) {
	return nil, false
}

// ========================================================

// ChooseFront simply returns the first column
func ChooseFront(among *Tees) (*Tees, bool) {

	if among == nil { panic( "Cannot choose from is nil!") }
	if Verbose {	among.PrintAways("Choosing from: ")	}

	e := among.Front()
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
//	- it returns false on zero-length as there is no point in trying such
func ChooseShort(among *Tees) (*Tees, bool) {
	var found bool
	var c *Tees
	s := 999999999
	for j := among.Front(); j != nil; j = j.Next() {
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
func ChooseUpto0(among *Tees) (*Tees, bool) {
	return ChooseUpto(among, 0)
}

// ChooseUpto1 returns the first column not longer than 1
func ChooseUpto1(among *Tees) (*Tees, bool) {
	return ChooseUpto(among, 1)
}

// ChooseUpto2 returns the first column not longer than 2
func ChooseUpto2(among *Tees) (*Tees, bool) {
	return ChooseUpto(among, 2)
}

// ChooseUpto3 returns the first column not longer than 3
func ChooseUpto3(among *Tees) (*Tees, bool) {
	return ChooseUpto(among, 3)
}

// ========================================================

// ChooseUpto returns the first column not longer than min
func ChooseUpto(among *Tees, min int) (*Tees, bool) {
	var found bool
	var c *Tees
	s := 999999999
	for j := among.Front(); j != nil; j = j.Next() {
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
