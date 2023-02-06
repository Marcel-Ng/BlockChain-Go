package main

import (
	"fmt"
	"net/http"
	"strconv"
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

	// this is the function that would be making a post request to this application
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// This is just a test code although we are going to rewrite this to make use of a post request
		new_user := r.URL.Query().Get("user")
		BALANCE[new_user] = 0
		fmt.Println(BALANCE)
	})

	http.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
		// This is to make transfers inbetween users and the should still be in
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
		var user_balance int
		user_balance = BALANCE[from].(int)
		if err != nil {
			// handle the error that is going to happen here incase we can't convert to interger
			fmt.Println(err)
			// should be able to end the function here.
		}

		if amount > user_balance {
			fmt.Println("User do not have a sufficient balance to make this trade")
			// should be also be able to end the function here
		}

		BALANCE[to] = amount + user_balance

		fmt.Print("User balance for ", to, "is: \t", BALANCE[to])
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3000", nil)
}
