// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
dance.go extends list.go with:

	- l.Dance( d *Dancing )
	- e.Dance( d *Dancing )
*/

package list

// ========================================================

// Dance l is where the dancing begins
func (l *List)     Dance( d *Dancing ) {
	l.fold(d)
	l.root.Dance( d )
	l.open(d)
}

// Dance e is where the dancing continues
func (e *Element)  Dance( d *Dancing ) {
	for i := e.next; i != e; i = i.next {
		d.OnGoal( i )  // Push
		i.away.fold(d) // fold
		d.Dance()      // Dance d is where the dancing recurs to
		i.away.open(d) // open
		d.OnFail()     // Pop
	}
}

// ========================================================
func (l *List)      fold( d *Dancing ) { 										l.root.away.UnLeaf(d);	l.unList(d)	}
func (l *List)      open( d *Dancing ) { 										l.root.away.ReLeaf(d);	l.reList(d)	}

func (e *Element)   fold( d *Dancing ) { for i := e.next; i != e; i = i.next { if i.list == nil || i != &e.list.root {	i.away.list.fold(d)	} } }
func (e *Element)   open( d *Dancing ) { for i := e.prev; i != e; i = i.prev { if i.list == nil || i != &e.list.root {	i.away.list.open(d)	} } }

func (l *List)    unList( d *Dancing ) { for i := l.root.next; i != &l.root; i = i.next { 				i.away.unList(d)	  } }
func (l *List)    reList( d *Dancing ) { for i := l.root.prev; i != &l.root; i = i.prev { 				i.away.reList(d)	  } }

func (e *Element) unList( d *Dancing ) { for i := e.next; i != e; i = i.next { if i.list == nil || i != &e.list.root {	i.away.UnLeaf(d)	} } }
func (e *Element) reList( d *Dancing ) { for i := e.prev; i != e; i = i.prev { if i.list == nil || i != &e.list.root {	i.away.ReLeaf(d)	} } }
