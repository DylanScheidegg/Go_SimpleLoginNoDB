package main

import (
	"net/http"

	"./controllers/accountcontroller"
)

//http://localhost:10000/account

func main() {
	http.HandleFunc("/account", accountcontroller.Index)
	http.HandleFunc("/account/index", accountcontroller.Index)
	http.HandleFunc("/account/login", accountcontroller.Login)
	http.HandleFunc("/account/welcome", accountcontroller.Welcome)
	http.HandleFunc("/account/logout", accountcontroller.Logout)
	http.ListenAndServe(":10000", nil)
}
