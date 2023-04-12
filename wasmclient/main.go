// go:build (js && wasm)

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"syscall/js"

	"github.com/serdarkalayci/carpool/api/adapters/comm/rest/dto"
)

const baseurl = "http://localhost:5550"

type errorMessage struct {
	Error string `json:"error"`
}

func main() {
	js.Global().Set("login", js.FuncOf(login))
	setUpUI()
	c := make(chan int)
	<-c
}

func setUpUI() {
	document := js.Global().Get("document")
	wrapperdiv := document.Call("getElementById", "wrapper")
	logindiv := document.Call("getElementById", "login")
	maindiv := document.Call("getElementById", "main")
	logincookie := document.Get("carpooltoken").String()
	fmt.Println("cookie:" + logincookie)
	loggedin := refreshtoken(logincookie)
	if loggedin {
		wrapperdiv.Call("removeChild", logindiv)
	} else {
		wrapperdiv.Call("removeChild", maindiv)
	}
}

func login(this js.Value, args []js.Value) interface{} {
	email := args[0].String()
	password := args[1].String()
	executeLogin(email, password)
	return nil
}

func refreshtoken(coookie string) bool {
	errorp := js.Global().Get("document").Call("getElementById", "error")
	result := make(chan bool)
	go func(c chan bool) {
		request, err := http.NewRequest("PUT", baseurl+"/login/refresh", nil)
		if err != nil {
			errorp.Set("innerHTML", err.Error())
		}
		request.Header.Set("Origin", "http://localhost:5800")
		request.AddCookie(&http.Cookie{Name: "carpooltoken", Value: coookie})
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			errorp.Set("innerHTML", err.Error())
		}
		if response.StatusCode == 200 {
			c <- true
		}
		c <- false
	}(result)
	return <-result
}

func executeLogin(email string, password string) {
	document := js.Global().Get("document")
	wrapperdiv := document.Call("getElementById", "wrapper")
	logindiv := document.Call("getElementById", "login")
	loginstatus := document.Call("getElementById", "login_status")
	go func() {
		loginDTO := dto.LoginRequest{Email: email, Password: password}
		fmt.Printf("loginDTO: %v", loginDTO)
		payload, err := json.Marshal(loginDTO)
		if err != nil {
			log.Fatal(err)
		}
		request, err := http.NewRequest("PUT", baseurl+"/login", bytes.NewBuffer(payload))
		jar, _ := cookiejar.New(nil)

		// Create a new HTTP client with the cookie jar
		client := &http.Client{
			Jar: jar,
		}
		response, err := client.Do(request)
		if err != nil {
			loginstatus.Set("innerHTML", err.Error())
		} else {
			fmt.Println("Cookies: ")
			for _, cookie := range jar.Cookies(response.Request.URL) {
				fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
			}
			for name, values := range response.Header {
				for _, value := range values {
					fmt.Printf("%s: %s\n", name, value)
				}
			}
			if response.StatusCode == 200 {
				loginstatus.Set("innerHTML", "login success")
				// document.Set("carpooltoken", response.Cookies()[0].Value)
				wrapperdiv.Call("removeChild", logindiv)
			} else { // No errors but no 200 either
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					loginstatus.Set("value", err.Error())
				} else {
					var errMsg errorMessage
					json.Unmarshal(body, &errMsg)
					loginstatus.Set("innerHTML", errMsg.Error)
				}

			}

		}
	}()
}
