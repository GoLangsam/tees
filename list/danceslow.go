// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
danceslow.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

	- l.DanceSlow( d *Dancing )
	- e.DanceSlow( d *Dancing )

*/
package list

// ========================================================

// DanceSlow l is where the dancing begins
func (l *List)     DanceSlow( d *Dancing ) {
	l.foldSlow(d)
	l.Root().DanceSlow( d )
	l.openSlow(d)
}

// DanceSlow e is where the dancing continues
func (e *Element)  DanceSlow( d *Dancing ) {
	e.ForEachNext( func(i *Element){
		d.OnGoal( i )		// Push
		i.Away().foldSlow(d)
		d.Dance()		// Dance d is where the dancing recurs to
		i.Away().openSlow(d)
		d.OnFail()		// Pop
	} )
}

// ========================================================
func (l *List)      foldSlow( d *Dancing ) { 					l.Root().Away().UnLeaf(d);	l.unListSlow(d)		}
func (l *List)      openSlow( d *Dancing ) { 					l.Root().Away().ReLeaf(d);	l.reListSlow(d)		}

func (e *Element)   foldSlow( d *Dancing ) { e.ForEachNext (func(i *Element) {	i.Away().List().foldSlow(d)	} ) }
func (e *Element)   openSlow( d *Dancing ) { e.ForEachPrev (func(i *Element) {	i.Away().List().openSlow(d)	} ) }

func (l *List)    unListSlow( d *Dancing ) { l.ForEachNext (func(i *Element) {	i.Away().unListSlow(d)		} ) }
func (l *List)    reListSlow( d *Dancing ) { l.ForEachPrev (func(i *Element) {	i.Away().reListSlow(d)		} ) }

func (e *Element) unListSlow( d *Dancing ) { e.ForEachNext (func(i *Element) {	i.Away().UnLeaf(d)		} ) }
func (e *Element) reListSlow( d *Dancing ) { e.ForEachPrev (func(i *Element) {	i.Away().ReLeaf(d)		} ) }
