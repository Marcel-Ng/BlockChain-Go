package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// this is for
)

const PORT = ":3000"

/**
* After all said and done more is actual said than done
 */

var BALANCE = map[string]interface{}{
	"marcel": 100000,
	"chidi":  4000,
}

type TransferPayload struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func transfer(w http.ResponseWriter, r *http.Request) {
	// of course you should handle the exception if the request method is not the POSt method
	decoder := json.NewDecoder(r.Body)
	var transferPayLoad TransferPayload
	error := decoder.Decode(&transferPayLoad)

	if error != nil {
		fmt.Println("Error decoding the data in the Json payload: \t", error)
	}
	fmt.Println(transferPayLoad)
	from := transferPayLoad.From
	to := transferPayLoad.To

	user_balance := BALANCE[from].(int)
	current_reciever_balance := BALANCE[to].(int)
	amount, err := strconv.Atoi(transferPayLoad.Amount)
	if err != nil {
		// handle the error that is going to happen here incase we can't convert to interger
		fmt.Println(err)
		return
	}
	if amount > user_balance {
		fmt.Println("User do not have a sufficient balance to make this trade")
		// should be also be able to end the function here
	}
	fmt.Println(BALANCE[to])
	fmt.Println(user_balance)
	fmt.Println(amount)
	BALANCE[from] = user_balance - amount
	BALANCE[to] = amount + current_reciever_balance

	fmt.Print("User balance for ", to, " is: \t", BALANCE[to])
	// this should be able to return a JSON
}

// this is the function that would be making a post request to this application
func createUser(w http.ResponseWriter, r *http.Request) {
	// This is just a test code although we are going to rewrite this to make use of a post request
	new_user := r.URL.Query().Get("user")
	BALANCE[new_user] = 0
	fmt.Println(BALANCE)
}

func getUserBalance(w http.ResponseWriter, r *http.Request) {
	// I am going to be printing the outputs of the users balance in the parameters here.
	user := r.URL.Query().Get("user")
	balance := BALANCE[user]
	fmt.Println("The balance is: \t", balance)
	fmt.Fprintf(w, "We are trying to get the balance")
}

func main() {
	fmt.Println("Go has started from here now")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/balance", getUserBalance)
	http.HandleFunc("/user", createUser)
	http.HandleFunc("/transfer", transfer)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(PORT, nil)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// This is a test function that is used in making sure that we now how to get a post request
// 	// after we might also use the mux thing here if we want to

// 	// Supposed to handle the case of wrong url
// 	fmt.Println("This tt route has actually started")

// 	switch r.Method {
// 	case "GET":
// 		http.ServeFile(w, r, "form.html")
// 	case "POST":

// 		reqDump, err := httputil.DumpRequest(r, true)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Printf("REQUEST:\n%s", string(reqDump))
// 		// w.Write([]byte("Hello World"))
// 		r.ParseForm()
// 		name := r.Form.Get("name")
// 		age := r.Form.Get("age")
// 		// decoder := json.NewDecoder(r.Body)
// 		// var data json_data
// 		// error := decoder.Decode(&data)
// 		// if error != nil {
// 		// 	// handle the error in case the data is not decoded properly
// 		// 	fmt.Println("Error occured while decoding the data: ", error)
// 		// 	return
// 		// }
// 		// fmt.Println(data.name)
// 		// fmt.Printf("REQUEST:\n%s", data)
// 		fmt.Println("name and age: ", name, " ", age)
// 		fmt.Println("Form Request: \n%v", json.Marshal(r.Form))

// 		// err := r.ParseForm()

// 		// if err != nil {
// 		// 	fmt.Fprintf(w, "there was an error trying to get the body of this request %v", err)
// 		// 	return
// 		// }

// 		// name := r.FormValue("name")
// 		// age := r.FormValue("age")
// 		// fmt.Println("The values in the form is: ", name, " and ", age)

// 	}
// }
