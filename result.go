// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"fmt"
)

//action result
type Result struct {
	//status
	Num int64

	//Message
	Msg string
}

//implement error interface
func (r *Result) Error() string {
	return fmt.Sprintf("%d: %s", r.Num, r.Msg)
}
