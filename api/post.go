package api

import (
	"net/http"
)

// Post ...
func Post(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("network")
	phoneNumber := r.FormValue("number")
	amount := r.FormValue("amount")
	//fmt.Println(code, phoneNumber, amount)
	
	Purchase(w, r, code, phoneNumber, amount)
}