// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
away.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

	- l.AwayLists()		[]*List
	- e.AwayLists()		[]*List


*/
package list

// ===========================================================================

func (l *List) AwayLists() []*List {
	var data = make([]*List, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e.AwayList())
	}
	return data
}

func (e *Element) AwayLists() []*List {
	return e.List().AwayLists()
}
