// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package turn

import (
	"github.com/GoLangsam/tees"
)

type Dancing struct {
	Dance   	func(*tees.Tees)

	OnGoal		func(*tees.Tees) 		bool
	OnFail	   	func(*tees.Tees) (*tees.Tees,	bool)
	OnLeaf		func(*tees.Tees) 		bool
}
