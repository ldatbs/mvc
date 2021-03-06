// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"net/http"
	"strings"
)

//Request struct, containe *http.Request, it's about to more variable.
type Request struct {
	//Module name
	Module string

	//Lowercase
	//ModuleLow string

	//Controller name
	Controller string

	//Lowercase
	//ControllerLow string

	//Action name
	Action string

	//Lowercase
	//ActionLow string

	//Method arguments sli.
	args []string

	//Method arguments map
	Args map[string]interface{}

	//Inited status
	inited Result

	//*http.Request
	HTTPRequest *http.Request
}


//New request
func newRequest(r *http.Request) *Request {
	nodes := strings.Split(r.URL.Path, "/")
	l := len(nodes)

	req := Request{Module: "System", Controller: "Index", Action: "Index", args: make([]string, 0), Args: make(map[string]interface{}, 0), HTTPRequest: r}

	if l > 1 && nodes[1] != "" {
		req.Module = nodes[1] //strings.Title(nodes[1])
	}
	if l > 2 && nodes[2] != "" {
		req.Controller = nodes[2] //pathToMethod(nodes[2])
	}
	if l > 3 && nodes[3] != "" {
		req.Action = nodes[3] //pathToMethod(nodes[3])
	}
	if l > 4 {
		req.args = nodes[4:]
	}

	return &req
}

// for example: transfer browse_by_set to BrowseBySet
func pathToMethod(path string) string {
	var method string
	sli := strings.Split(path, "_")
	for _, v := range sli {
		method += strings.Title(v)
	}
	return method
}
