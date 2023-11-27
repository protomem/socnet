package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8083", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "Auth Service v0.1.0")
	}))
}
