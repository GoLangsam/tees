// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================

// foldr _ z [] = z
// foldr f z (x:xs) = f x (foldr f z xs)

// FoldAny folds Tees using f and initial
func FoldAny(
	f func(*aTee, interface{}) interface{},
	i iterator,
	initial interface{}) interface{} {
	if i == nil {
		return initial
	}

	var result = initial
	for e := i.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}
	return result
}

// FoldInt folds Tees using f(*aTee, int) and initial
func FoldInt(
	f func(*aTee, int) int,
	i iterator,
	initial int) int {
	if i == nil {
		return initial
	}

	var result = initial
	for e := i.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}
	return result
}

// FoldString folds Tees using f(*aTee, string) and initial
func FoldString(
	f func(*aTee, string) string,
	i iterator,
	initial string) string {
	if i == nil {
		return initial
	}

	var result = initial
	for e := i.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}
	return result
}
