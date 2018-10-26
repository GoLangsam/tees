// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tees

import (
	"github.com/GoLangsam/tees/list"
)

// ===========================================================================

// CanIter allows to iterate forward by starting with Front() and, if non-nil, repeating Next() until Next() returns nil
type CanIter interface {
	Front() *list.Element
	Next() *list.Element
	ForEachNext(f func(*list.Element))
}

// CanReti allows to iterate backward from Front() and, if non-nil, repeating Next() until Next() returns nil
type CanReti interface {
	Back() *list.Element
	Prev() *list.Element
	ForEachPrev(f func(*list.Element))
}

// HasLen models any collection which can report it's Length.
type HasLen interface {
	Len() int
}

// HasValue models anything which has a Value of unspecific type.
type HasValue interface {
	Value() interface{}
}

// HasRoot models anything which has a root element.
type HasRoot interface {
	Root() *list.Element
}

// ==== composed interfaces

// canIterN
type canIterN interface {
	CanIter
	HasLen
}

// canRetiN
type canRetiN interface {
	CanReti
	HasLen
}

// canIterNV
type canIterNV interface {
	canIterN
	HasValue
}

// canRetiNV
type canRetiNV interface {
	canRetiN
	HasValue
}

// ===========================================================================
