package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	_, _ = fmt.Fprintf(w, "Hello, %s", name)
}

func GreeterHandler( w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}
func main()  {
	_ = http.ListenAndServe(":3002", http.HandlerFunc(GreeterHandler))
}
