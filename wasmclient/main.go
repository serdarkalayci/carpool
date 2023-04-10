package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall/js"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
)

func main() {
	js.Global().Set("login", js.FuncOf(login))
	c := make(chan int)
	<-c
}

func login(this js.Value, args []js.Value) interface{} {
	email := args[0].String()
	password := args[1].String()
	executeLogin(email, password)
	return nil
}

func executeLogin(email string, password string) {
	document := js.Global().Get("document")
	navdiv := document.Call("getElementById", "nav")
	h2 := document.Call("createElement", "h2")
	go func() {
		loginDTO := dto.LoginRequest{Email: email, Password: password}
		fmt.Printf("loginDTO: %v", loginDTO)
		payload, err := json.Marshal(loginDTO)
		if err != nil {
			log.Fatal(err)
		}
		// prefreq, err := http.NewRequest("OPTIONS", "http://localhost:5500/login", bytes.NewBuffer(payload))
		// prefreq.Header.Set("Content-Type", "application/json")
		// prefreq.Header.Set("Access-Control-Request-Method", "POST")
		// prefreq.Header.Set("Origin", "http://localhost:5800")
		// // prefreq.Header.Set("Access-Control-Request-Headers", "Content-Type, lang")
		// prefresp, err := http.DefaultClient.Do(prefreq)
		// if err == nil && prefresp.StatusCode == 200 {
		// fmt.Printf("prefresp: %v", prefresp)
		request, err := http.NewRequest("PUT", "http://localhost:5550/login", bytes.NewBuffer(payload))
		// request.Header.Set("Content-Type", "application/json")
		// request.Header.Set("Origin", "http://localhost:5800")
		// request.Header.Set("Access-Control-Request-Method", "PUT")
		// request.Header.Set("Accept", "*/*")
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			h2.Set("innerHTML", "error")
		} else {
			for _, cookie := range response.Cookies() {
				fmt.Println(cookie.Name)
			}
			if response.StatusCode == 200 {
				h2.Set("innerHTML", "login success")
				wrapperdiv := document.Call("getElementById", "wrapper")
				logindiv := document.Call("getElementById", "login")
				wrapperdiv.Call("removeChild", logindiv)
			}
			_, err = ioutil.ReadAll(response.Body)
			if err != nil {
				h2.Set("innerHTML", "error")
			} else {

				h2.Set("innerHTML", "post"+response.Status)
			}
			navdiv.Call("appendChild", h2)
		}
		// } else {
		// 	if err != nil {
		// 		h2.Set("innerHTML", err.Error())
		// 	} else {
		// 		h2.Set("innerHTML", "options"+prefresp.Status)
		// 	}
		// 	navdiv.Call("appendChild", h2)
		// }
	}()
}
