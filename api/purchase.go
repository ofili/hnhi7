package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Client ...
type Client struct {
	Code string
	Amount int64
	Phonenumber string
	SecretKey string
}

const secretKey = "hfucj5jatq8h"
var url = "https://sandbox.wallets.africa/bills/airtime/purchase"

// Purchase ...
func Purchase(w http.ResponseWriter, r *http.Request,code, number, amt string) {
	
	cost,_ := strconv.ParseInt(amt, 10, 64)
	
	NewClient := &Client {
		Code: code,
		Amount: cost,
		Phonenumber: number,
		SecretKey: secretKey,
	}

	// Token
	const publicToken = "uvjqzm5xl6bw"
 	var bearer = "Bearer " + publicToken
	
	 // Body of response
	m, err := json.Marshal(NewClient); if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

  client := &http.Client {
  }
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(m))
  if err != nil {
	  w.WriteHeader(http.StatusInternalServerError)
    fmt.Println(err)
  }

  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", bearer)

  res, err := client.Do(req); if err!= nil{
	w.WriteHeader(http.StatusBadRequest)
    log.Println(err)
  }

  defer res.Body.Close()

  rBody := Responses{}
  json.NewDecoder(res.Body).Decode(&rBody)

  if res.Status == "200 OK" {
	  // Show popup message (successful)
	http.Redirect(w, r, "/success", http.StatusFound)
  }else {
	  w.WriteHeader(http.StatusBadRequest)
	  // Show popup message (fail)
	  fmt.Fprintln(w, rBody.Message)
  }

// Print status
  fmt.Println("response Status:", res.Status)
}