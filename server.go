package main

import (
	"fmt"
	"net/http"
)

type MyServer struct{}

func (m MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my workspace!")
}
func main() {
	var m MyServer
	http.ListenAndServe("localhost:4000", m)
}
