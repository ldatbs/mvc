// Copyright 2014 The fav Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mvc

import (
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Router struct {
}

//Registered routers
var routers map[string]reflect.Value = make(map[string]reflect.Value)

//new router register
func NewRouter(c interface{}) {
	value := reflect.ValueOf(c)
	routers[value.Elem().Type().Name()] = value
}

//Every request run this function
func newHandler(fn actionLoadFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := newRequest(r)

		log.Printf("\n\n") //空行分隔
		log.Printf("%#v\n", req)

		controller, s := routers[req.Module+req.Controller]
		if s == false {
			log.Printf("controller not found: %s%s\n", req.Module, req.Controller)
			return
		}

		//Invoke Load() function
		if fn != nil {
			req.inited = fn(req)
			if req.inited.Num < 0 {
				req.Controller = "Index"
				req.Action = "Err"

				log.Printf("Load falure: %s\n", req.inited.Msg)
			}
		}

		rq := controller.Elem().FieldByName("Request")
		rq.Set(reflect.ValueOf(req))

		writer := View{Request: req, Data: make(map[string]interface{}, 0), ResponseWriter: w}

		rw := controller.Elem().FieldByName("View")
		rw.Set(reflect.ValueOf(&writer))

		method := req.Action

		//action
		action := controller.MethodByName(method)
		if action.IsValid() {
			log.Printf("method [%s] found\n", method)

			typ := action.Type()
			numIn := typ.NumIn()

			if len(req.args) >= numIn {
				pass := true
				in := make([]reflect.Value, numIn)

				for i := 0; i < numIn; i++ {
					actionIn := typ.In(i)
					kind := actionIn.Kind()
					v, err := paramConversion(kind, req.args[i])
					if err != nil {
						pass = false
						log.Printf("string convert to %s failure: %s\n", kind, req.args[i])
					} else {
						in[i] = v
						req.Args[actionIn.Name()] = v
					}
				}

				if pass == true {
					resultSli := action.Call(in)
					result := resultSli[0].Interface().(Result)
					//log.Printf("%#v\n", result)
					writer.realRender(result)
				} else {
					log.Printf("%s's paramters failure\n", method)
				}
			} else {
				log.Printf("method [%s]'s in arguments wrong\n", method)
			}

		} else {
			log.Printf("method [%s] not found\n", method)
		}
	}
}

func paramConversion(kind reflect.Kind, arg string) (reflect.Value, error) {
	var v reflect.Value
	var err error

	switch kind {
	case reflect.String:
		v = reflect.ValueOf(arg)
	case reflect.Int64:
		var i64 int64
		i64, err = strconv.ParseInt(arg, 10, 64)
		v = reflect.ValueOf(i64)
	default:
		log.Printf("string convert to int failure: %s\n", arg)
	}

	return v, err
}
