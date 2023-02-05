package main

import (
	"fmt"
	"net/http"
)

/**
* After all said and done more is actual said than done
 */

//  type balance struct {
// 	Name string
// 	amount float32
//  }

func main() {
	fmt.Println("Go has started from here now")

	// Balance := balance {
	// 	Name: ,
	// }

	var BALANCE = map[string]interface{}{
		"marcel": 100000,
		"chidi":  4000,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/balance", func(w http.ResponseWriter, r *http.Request) {

		// I am going to be printing the outputs of the users balance in the parameters here.
		user := r.URL.Query().Get("user")
		balance := BALANCE[user]
		fmt.Println("The balance is: \t", balance)
		fmt.Fprintf(w, "We are trying to get the balance")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", nil)
}
