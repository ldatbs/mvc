// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"log"
	"fmt"
)

func Url(m string, c string, a string, args ...interface{}) string {
	module, s := Conf.Modules[m]
	if s == false {
		log.Printf("module not exists: %s\n", m)
		return ""
	}

	return fmt.Sprintf("http://%s/%s/%s", module.Listen, c, a)
}

func Misc(p string) string {
	f, s := Conf.FileServers["Misc"]
	if s == false {
		log.Printf("misc file server config not exists\n")
		return ""
	}

	return fmt.Sprintf("http://%s%s", f.Listen, p)
}

func FavUI(p string) string {
	f, s := Conf.FileServers["FavUI"]
	if s == false {
		log.Printf("FavUI config not exists\n")
		return ""
	}

	return fmt.Sprintf("http://%s%s", f.Listen, p)
}

func Theme(p string) string {
	//theme
	t := "default"

	//full path
	fp := fmt.Sprintf("/themes/%s%s", t, p)
	return Misc(fp)
}
