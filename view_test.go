// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"testing"
)

func TestMethodToPath(t *testing.T) {
	method := "BrowseBySet"
	path := methodToPath(method)
	if path != "browse_by_set" {
		t.Errorf("methodToPath failure: %#v", path)
	}
}
