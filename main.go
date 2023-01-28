package main

import (
	"fmt"
	"net/http"
)

/**
* After all said and done more is actual said than done
 */

func main() {
	fmt.Println("Go has started from here now")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", nil)
}
