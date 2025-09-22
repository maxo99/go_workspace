package main

// https://gowebexamples.com/http-server/


import (
	"net/http"
)



func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "Welcome to my website!")
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))

	})
	http.ListenAndServe(":8080", nil)

	// curl localhost:8080/static/image.png
}