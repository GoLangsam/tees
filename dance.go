// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ========================================================

// Dance e is where the dancing begins.
func Dance(e Dancer, d Dancing) {
	unC(e) //   fold Col
	forEachNext(e, func(i This) {
		d.OnGoal(i)   // Publish candidate
		unR(i.Away()) //   fold Row
		d.Dance()     // Dance d is where the dancing recurs to
		reR(i.Away()) //   open Row
		d.OnFail()    // Plopp :-(
	})
	reC(e) //   open Col
}

// ========================================================
func unC(c Dancer) { c.Away().UnLink(); unL(c) }
func reC(c Dancer) { c.Away().ReLink(); reL(c) }

func unR(r Dancer) { forEachNext(r, func(i This) { unC(i.Away().List()) }) }
func reR(r Dancer) { forEachPrev(r, func(i This) { reC(i.Away().List()) }) }

func unL(l Dancer) { forEachNext(l, func(i This) { unE(i.Away()) }) }
func reL(l Dancer) { forEachPrev(l, func(i This) { reE(i.Away()) }) }

func unE(e Dancer) { forEachNext(e, func(i This) { i.Away().UnLink() }) }
func reE(e Dancer) { forEachPrev(e, func(i This) { i.Away().ReLink() }) }

// ========================================================

func forEachNext(s Dancer, f func(This)) { s.ForEachNext(f) }
func forEachPrev(s Dancer, f func(This)) { s.ForEachPrev(f) }
