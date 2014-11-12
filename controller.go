// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"log"
)

//controller struct
type Controller struct {
	//Request
	Request *Request

	//Responsewriter
	View *View
}

//General error page
func (c *Controller) IndexErr(status int16, msg string) Result {
	log.Printf("SystemIndexErr\n")

	return c.View.Render()
}

