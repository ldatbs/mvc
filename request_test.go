// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"testing"
)

func TestPathToMethod(t *testing.T) {
	path := "browse_by_set"
	method := pathToMethod(path)
	if method != "BrowseBySet" {
		t.Error("pathToMethod failure")
	}
}
