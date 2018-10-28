// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ========================================================

// Dance e is where the dancing begins.
func Dance(e dancer, d Dancing) {
	unC(e) //   fold Col
	forEachNext(e, func(i *aTee) {
		d.OnGoal(i)   // Publish candidate
		unR(i.Away()) //   fold Row
		d.Dance()     // Dance d is where the dancing recurs to
		reR(i.Away()) //   open Row
		d.OnFail()    // Plopp :-(
	})
	reC(e) //   open Col
}

// ========================================================
func unC(c dancer) { c.Away().UnLink(); unL(c) }
func reC(c dancer) { c.Away().ReLink(); reL(c) }

func unR(r dancer) { forEachNext(r, func(i *aTee) { unC(i.Away().List()) }) }
func reR(r dancer) { forEachPrev(r, func(i *aTee) { reC(i.Away().List()) }) }

func unL(l dancer) { forEachNext(l, func(i *aTee) { unE(i.Away()) }) }
func reL(l dancer) { forEachPrev(l, func(i *aTee) { reE(i.Away()) }) }

func unE(e dancer) { forEachNext(e, func(i *aTee) { i.Away().UnLink() }) }
func reE(e dancer) { forEachPrev(e, func(i *aTee) { i.Away().ReLink() }) }

// ========================================================

func forEachNext(s dancer, f func(*aTee)) { s.ForEachNext(f) }
func forEachPrev(s dancer, f func(*aTee)) { s.ForEachPrev(f) }
