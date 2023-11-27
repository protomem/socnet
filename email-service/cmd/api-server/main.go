package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8086", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "Email Service v0.1.0")
	}))
}
