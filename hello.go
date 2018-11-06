package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Google App Engine for Go world.")
}
