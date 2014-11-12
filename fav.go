// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"log"
	"net/http"
	"os"
	"strings"
)

//application root path
var HomeDir string

//mvc config
var Conf conf

//action load function
type actionLoadFunc func(r *Request) Result

//file server
type FileServer struct {
	//Listen
	Listen string

	//Path
	Path string

	//Root
	Root string
}

//Init fav
func init() {
	//init Home
	var err error
	HomeDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	HomeDir = strings.Replace(HomeDir, "\\", "/", -1)

	log.Printf("%#v\n", HomeDir)

	//init configs
	Conf.ReadFromFile(HomeDir + "/fav.json")
	Conf.Init()

	//init log
	initLog()
}

func initLog() {
	//var Log *log.Logger
	//Log = log.New(os.Stderr, "", flag)

	switch Conf.Environment {
	case 0:
		log.SetOutput(os.Stderr)
		log.SetPrefix("")
		log.SetFlags(log.Lshortfile)
	case 1:
		log.SetOutput(os.Stderr)
		log.SetPrefix("")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	case 2, 3:
		log.SetOutput(os.Stderr)
		log.SetPrefix("")
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}
}

//Start new HTPP server
func NewServer(fn actionLoadFunc) {
	muxList := make(map[string]*http.ServeMux)

	//misc
	for _, v := range Conf.FileServers {
		_, s := muxList[v.Listen]
		if s == false {
			muxList[v.Listen] = http.NewServeMux()
		}

		if v.Path == "/" {
			muxList[v.Listen].Handle(v.Path, http.FileServer(http.Dir(v.Root)))
		} else {
			// To serve a directory on disk (/tmp) under an alternate URL
			// path (/tmpfiles/), use StripPrefix to modify the request
			// URL's path before the FileServer sees it:
			muxList[v.Listen].Handle(v.Path, http.StripPrefix(v.Path, http.FileServer(http.Dir(v.Root))))
		}
	}

	//file server
	/*for listen, v := range Conf.FileServers {
		_, s := muxList[listen]
		if s == false {
			muxList[listen] = http.NewServeMux()
		}

		for path, root := range v {
			if path == "/" {
				muxList[listen].Handle(path, http.FileServer(http.Dir(root)))
			} else {
				muxList[listen].Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(root))))
			}
		}
	}*/

	//modules
	for _, m := range Conf.Modules {
		_, s := muxList[m.Listen]
		if s == false {
			muxList[m.Listen] = http.NewServeMux()
		}

		muxList[m.Listen].HandleFunc("/" + m.Name + "/", newHandler(fn))
	}

	//log.Printf("%#v\n", muxList)

	l := len(muxList)
	i := 0
	for listen, mux := range muxList {
		i++

		//log.Printf("%#v\n", mux)

		if i ==l {
			//why the last listenning no go?
			NewHost(listen, mux)
		} else {
			go NewHost(listen, mux)
		}
	}
}

//new host
func NewHost(listen string, mux *http.ServeMux){
	err := http.ListenAndServe(listen, mux)
	if err != nil {
		log.Fatal(err)
	}
}
