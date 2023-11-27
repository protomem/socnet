package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8085", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintln(w, "Recommended Service v0.1.0")
	}))
}
