// +build appengine

package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// c := appengine.NewContext(r)

	// t := strconv.Itoa(c)
	// fmt.Println(t)
	// fmt.Println("type:", reflect.TypeOf(c))
	fmt.Fprint(w, "Hello!")
}
