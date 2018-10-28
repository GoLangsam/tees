// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================

//  let rec sigma f = function
//     | [] -> 0
//     | x :: l -> f x + sigma f l;;
//
// val sigma : ('a -> int) -> 'a list -> int = <fun>

// Sigma returns the sum of the results of applying a given function f to each tee
func Sigma(
	f func(*aTee) int,
	i iterator,
) int {
	if i == nil {
		return 0
	}

	var result = 0
	for e := i.Front(); e != nil; e = e.Next() {
		result += f(e)
	}
	return result
}

// Note: Sigma may also be expressed using
//	FoldInt(func(e *aTee, r int) int {return r + f(e)}, i, 0)

// SigmaInt returns the sum of the results of applying a given function f to each tee using
//	FoldInt(func(e *aTee, r int) int {return r + f(e)}, i, 0)
func SigmaInt(
	f func(*aTee) int,
	i iterator,
) int {
	return FoldInt(func(e *aTee, r int) int { return r + f(e) }, i, 0)
}
