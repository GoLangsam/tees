// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
package deeh provides a normal (=non-concurrency-safe) stack
for *Element

*/
package deeh

import (
	"fmt"
)

// Stack implements a normal (=non-concurrency-safe) stack
// for *Element

type Stack struct {
	stack []*Element
}

// New returns a new empty stack with given initial capacity
func New(cap int) *Stack {
	return new(Stack).Init(cap)
}

// Init returns an empty stack with given initial capacity
func (s *Stack) Init(cap int) *Stack {
	s.stack = make([]*Element, 0, cap)
	return s
}

// Push sth onto the current stack
func (s *Stack) Push(l *Element) {
	//	s.Lock()
	//	defer s.Unlock()

	s.stack = append(s.stack, l)
}

// Pop sth off the current stack
func (s *Stack) Pop() *Element {
	//	s.Lock()
	//	defer s.Unlock()

	var p = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return p
}

// Drop pops sth silently off the current stack
func (s *Stack) Drop() {
	//	s.Lock()
	//	defer s.Unlock()

	//	var p = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return // p
}

// Top returns the top of the current stack
func (s *Stack) Top() *Element {
	//	s.Lock()
	//	defer s.Unlock()

	return s.stack[len(s.stack)-1]
}

// Get returns a copy of the current stack
func (s *Stack) Get() []*Element {
	//	s.Lock()
	//	defer s.Unlock()

	var stack = make([]*Element, len(s.stack))
	copy(stack, s.stack)
	return stack
}

// Len returns the length of the current stack
func (s *Stack) Len() int {
	//	s.Lock()
	//	defer s.Unlock()

	return len(s.stack)
}

// ========================================================

// Print prints the current stack (= the solution)
func (s *Stack) Print() {
	//	s.Lock()
	//	defer s.Unlock()

	fmt.Print("Solution: ")
	fmt.Println(s.Len())
	for _, c := range s.stack {
		c.PrintValue()
		fmt.Print(": ")
		for e := c.Front(); e != nil; e = e.Next() {
			e.Away().List().Root().PrintValue()
			fmt.Print(" ")
		}
		fmt.Println(".")
	}
}
