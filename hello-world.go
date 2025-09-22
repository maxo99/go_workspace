package main

// https://gowebexamples.com/hello-world/

import (
	"fmt"
	"net/http"
)

func hello_world() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)

	// curl localhost:8080
}
