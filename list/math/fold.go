// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// foldr _ z [] = z
// foldr f z (x:xs) = f x (foldr f z xs)

// FoldAny folds a list (from front to next) using f and initial
func FoldAny(
	f func(*list.Element, interface{}) interface{},
	l *list.List,
	initial interface{}) interface{} {
	if l == nil {
		return initial
	}

	var result = initial
	for e := l.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}
	return result
}

// FoldInt folds a list (from front to next) using f(*Element, int) and initial
func FoldInt(
	f func(*list.Element, int) int,
	l *list.List,
	initial int) int {
	if l == nil {
		return initial
	}

	var result = initial
	for e := l.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}
	return result
}

// FoldString folds a list (from front to next) using f(*Element, string) and initial
func FoldString(
	f func(*list.Element, string) string,
	l *list.List,
	initial string) string {
	if l == nil {
		return initial
	}

	var result = initial
	for e := l.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}
	return result
}
