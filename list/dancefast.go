// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
dancefast.go extends list.go with:

	- l.DanceFast( d *Dancing )
	- e.DanceFast( d *Dancing )
*/

package list

// ========================================================

// DanceFast l is where the dancing begins
func (l *List) DanceFast(d *Dancing) {
	var a, p, n *Element

	{
		{
			{
				a = l.root.away
				p, n = a.prev, a.next
				n.prev, p.next = p, n
				a.list.len--
				for y := l.root.next; y != &l.root; y = y.next {
					for z := y.away.next; z != y.away; z = z.next {
						if z.list == nil || z != &l.root.list.root {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = p, n
							a.list.len--
						}
					}
				}
			}
		}
	}
	for x := l.root.next; x != &l.root; x = x.next {
		d.OnGoal(x) // Push
		for i := x.away.next; i != x.away; i = i.next {
			if i.list == nil || i != &x.away.list.root {
				a = i.away.list.root.away
				p, n = a.prev, a.next
				n.prev, p.next = p, n
				a.list.len--
				for y := i.away.list.root.next; y != &i.away.list.root; y = y.next {
					for z := y.away.next; z != y.away; z = z.next {
						if z.list == nil || z != &y.away.list.root {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = p, n
							a.list.len--
						}
					}
				}
			}
		}
		d.Dance() // Dance d is where the dancing recurs to
		for i := x.away.prev; i != x.away; i = i.prev {
			if i.list == nil || i != &x.away.list.root {
				for y := i.away.list.root.prev; y != &i.away.list.root; y = y.prev {
					for z := y.away.prev; z != y.away; z = z.prev {
						if z.list == nil || z != &y.away.list.root {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = a, a
							a.list.len++
						}
					}
				}
				a = i.away.list.root.away
				p, n = a.prev, a.next
				n.prev, p.next = a, a
				a.list.len++
			}
		}
		d.OnFail() // Pop
	}
	{
		{
			{
				for y := l.root.prev; y != &l.root; y = y.prev {
					for z := y.away.prev; z != y.away; z = z.prev {
						if z.list == nil || z != &y.away.list.root {
							a = z.away
							p, n = a.prev, a.next
							n.prev, p.next = a, a
							a.list.len++
						}
					}
				}
				a = l.root.away
				p, n = a.prev, a.next
				n.prev, p.next = a, a
				a.list.len++
			}
		}
	}
}

// DanceFast e is where the dancing continues.
//  Note: You must not call this directly - implemented only for symmetry, see interface Spot.
func (e *Element) DanceFast(d *Dancing) {
	panic("You cannot call this directly - implemented only for symmetry, see interface Spot")
}
