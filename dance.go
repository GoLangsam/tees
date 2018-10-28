// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// ========================================================
type Dancer interface {
	Away() *list.Element
	AwayList() *list.List
	ForEachNext(f func(*list.Element))
	ForEachPrev(f func(*list.Element))
	Root() *list.Element
}

// ========================================================

// Dance l is where the dancing begins
func Dance(l Dancer, d *list.Dancing) {
	unC(l)
	dance(l.Root(), d)
	reC(l)
}

// dance e is where the dancing continues
func dance(e Dancer, d *list.Dancing) {
	ForEachNext(e, func(i *list.Element) {
		d.OnGoal(i) // Publish candidate
		unR(i.Away())
		d.Dance() // Dance d is where the dancing recurs to
		reR(i.Away())
		d.OnFail() // Pop
	})
}

// ========================================================
func unC(c Dancer) { c.Away().UnLink(); unL(c) }
func reC(c Dancer) { c.Away().ReLink(); reL(c) }

func unR(r Dancer) { ForEachNext(r, unAf) }
func reR(r Dancer) { ForEachPrev(r, reAf) }

func unL(l Dancer) { ForEachNext(l, unEf) }
func reL(l Dancer) { ForEachPrev(l, reEf) }

func unE(e Dancer) { ForEachNext(e, unKf) }
func reE(e Dancer) { ForEachPrev(e, reKf) }

// ========================================================

var unAf = func(i *list.Element) { unC(i.AwayList()) }
var reAf = func(i *list.Element) { reC(i.AwayList()) }

var unEf = func(i *list.Element) { unE(i.Away()) }
var reEf = func(i *list.Element) { reE(i.Away()) }

var unKf = func(i *list.Element) { i.Away().UnLink() }
var reKf = func(i *list.Element) { i.Away().ReLink() }

// ========================================================

func ForEachNext(s Dancer, f func(*list.Element)) { s.ForEachNext(f) }
func ForEachPrev(s Dancer, f func(*list.Element)) { s.ForEachPrev(f) }
