// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"encoding/json"
	"io/ioutil"
	"github.com/favframework/db"
	"github.com/favframework/dump"
	"strings"
	"regexp"
	"log"
)

//mvc config struct
type conf struct {
	//0:development 1:testing 2:staging 3:production
	Environment int8

	//module list
	Modules map[string]Module

	//Misc
	FileServers map[string]FileServer

	//file servers
	//FileServers map[string]map[string]string

	//database connection string
	DB map[string]db.Conf

	//cache configs
	//Cache CacheConf
}

//init config
func (c *conf)ReadFromFile(p string) {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Println(err)
		return
	}

	txt := string(b)

	//log.Printf("%#v\n", txt)

	//replace path
	txt = regexp.MustCompile(`\{HomeDir\}`).ReplaceAllLiteralString(txt, HomeDir)

	//log.Printf("%#v\n", txt)

	//decode json
	dec := json.NewDecoder(strings.NewReader(txt))
	err = dec.Decode(c)
	if err != nil {
		log.Println(err)
		return
	}
}

//init config
func (c *conf)Init() {
	//default module
	if c.Modules == nil {
		c.Modules = make(map[string]Module)
		c.Modules["System"] = Module{Name: "System"}
	}

	//module name
	for k, m := range c.Modules {
		if m.Name == "" {
			m.Name = k //key as module name
			c.Modules[k] = m
		}
	}

	if c.FileServers == nil {
		c.FileServers = make(map[string]FileServer)
	}

	if c.DB == nil {
		c.DB = make(map[string]db.Conf)
	}

	//log.Printf("%#v\n", Conf)
	dump.Dump(Conf)
}
