package api


func Purchase(w http.ResponseWriter, r *http.Request,code, number, amt string) {
	//fmt.Println("From form",amt)
	parseamt,_ := strconv.ParseInt(amt, 10, 64)
	//fmt.Println(parseamt)
	NewClient := &ClientRequest {
		Code: code,
		Amount: parseamt,
		Phonenumber: number,
		SecretKey: secretKey,
	}
	const publicToken = "uvjqzm5xl6bw"  // Todo: to change
 	var bearer = "Bearer " + publicToken
	
	requestBody, err := json.Marshal(NewClient); if err!= nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

  client := &http.Client {
  }
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

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

  responsebody := response{}
  json.NewDecoder(res.Body).Decode(&responsebody)

  if res.Status == "200 OK" {
	http.Redirect(w, r, "/success", http.StatusFound)  // popup message showing success
  }else {
	  w.WriteHeader(http.StatusBadRequest)
	  fmt.Fprintln(w, responsebody.Message)	// popup message showing error message
  }

  
  //fmt.Println(responsebody.Message)

  //body, err := ioutil.ReadAll(res.Body)

  //fmt.Println("Response body:",string(body))
  //fmt.Println(string(body))
  fmt.Println("response Status:", res.Status)
}