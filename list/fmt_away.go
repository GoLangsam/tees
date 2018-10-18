// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
fmt_away.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpful, such as:

	- l.PrintAways()
	- e.PrintAways()

*/

package list

import (
	"fmt"
)

// ===========================================================================

// PrintAways prints all away lists.
func (l *List) PrintAways(args ...interface{}) {
	if l.print(args...) {
		//		fmt.Print( "List=" )
		l.Print()
		//		fmt.Print( "\t" )
		// Iterate through list and print its contents.
		for e := l.Front(); e != nil; e = e.Next() {
			e.PrintAways()
			//			fmt.Print( " | ")
		}
		//		fmt.Print( "Total=" )
		//		fmt.Println( l.len )
	}
}

// PrintAways prints its away list.
func (e *Element) PrintAways(args ...interface{}) {
	if e.print(args...) {
		if e.away == nil {
			fmt.Println("Element's Away is nil!")
			return
		}
		e.away.list.PrintAtomValues("=> ")
	}
}
