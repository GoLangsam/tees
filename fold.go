// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

// ===========================================================================

// foldr _ z [] = z
// foldr f z (x:xs) = f x (foldr f z xs)

// FoldAny folds Tees using f and initial i.
func FoldAny(
	f func(*aTee, interface{}) interface{},
	l iterator,
	i interface{},
) (result interface{}) {

	result = i

	if l != nil {
		for e := l.Front(); e != nil; e = e.Next() {
			result = f(e, result)
		}
	}

	return
}

// FoldInt folds Tees using f(*aTee, int) and initial i.
func FoldInt(
	f func(*aTee, int) int,
	l iterator,
	i int,
) (result int) {

	result = i

	if l != nil {
		for e := l.Front(); e != nil; e = e.Next() {
			result = f(e, result)
		}
	}

	return
}

// FoldString folds Tees using f(*aTee, string) and initial
func FoldString(
	f func(*aTee, string) string,
	l iterator,
	i string,
) (result string) {

	result = i

	if l != nil {
		for e := l.Front(); e != nil; e = e.Next() {
			result = f(e, result)
		}
	}

	return
}
