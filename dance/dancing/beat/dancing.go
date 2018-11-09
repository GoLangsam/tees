// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package beat

import (
	"github.com/GoLangsam/tees"
)

// Dancing is what dancer needs to dance.
type Dancing struct {
	Dance func(*tees.Tees)

	OnGoal func()
	OnFail func()
	OnLeaf func()
}
