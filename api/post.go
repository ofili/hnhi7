package api


func Post(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("network")
	phoneNumber := r.FormValue("number")
	amount := r.FormValue("amount")
	//fmt.Println(code, phoneNumber, amount)
	
	api.Purchase(w, r, code, phoneNumber, amount)
}