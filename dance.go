// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// ========================================================
// A Dancer loves dancing his Dance.
type Dancer interface {
	Away() *list.Element
	ForEachNext(f func(*list.Element))
	ForEachPrev(f func(*list.Element))
}

// ========================================================

// Dance e is where the dancing begins.
func Dance(e Dancer, d *list.Dancing) {
	unC(e)		//   fold Col
	dance(e, d)	// dance dance dance
	reC(e)		//   open Col
}

// dance e is where the dancing continues
func dance(e Dancer, d *list.Dancing) {
	ForEachNext(e, func(i *list.Element) {
		d.OnGoal(i)	// Publish candidate
		unR(i.Away())	//   fold Row
		d.Dance()	// Dance d is where the dancing recurs to
		reR(i.Away())	//   open Roe
		d.OnFail()	// Plopp :-(
	})
}

// ========================================================
func unC(c Dancer) { c.Away().UnLink(); unL(c) }
func reC(c Dancer) { c.Away().ReLink(); reL(c) }

func unR(r Dancer) { ForEachNext(r, func(i *list.Element) { unC(i.Away().List()) }) }
func reR(r Dancer) { ForEachPrev(r, func(i *list.Element) { reC(i.Away().List()) }) }

func unL(l Dancer) { ForEachNext(l, func(i *list.Element) { unE(i.Away()) }) }
func reL(l Dancer) { ForEachPrev(l, func(i *list.Element) { reE(i.Away()) }) }

func unE(e Dancer) { ForEachNext(e, func(i *list.Element) { i.Away().UnLink() }) }
func reE(e Dancer) { ForEachPrev(e, func(i *list.Element) { i.Away().ReLink() }) }

// ========================================================

func ForEachNext(s Dancer, f func(*list.Element)) { s.ForEachNext(f) }
func ForEachPrev(s Dancer, f func(*list.Element)) { s.ForEachPrev(f) }
