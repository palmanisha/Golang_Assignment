package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Number ..
type Number struct {
	A int `json:"a"`
	B int `json:"b"`
}

// DecodeFactroial ..
func DecodeFactroial(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var number Number
	json.Unmarshal(reqBody, &number)
	ch := make(chan int)
	ch1 := make(chan int)
	go func() {
		ch <- Factroial(number.A)

	}()

	go func() {
		ch1 <- Factroial(number.B)
	}()
	a := <-ch
	b := <-ch1
	json.NewEncoder(w).Encode(a * b)

}

// Factroial ...
func Factroial(number int) int {

	fact := 1

	for i := 1; i <= number; i++ {

		fact = fact * i

	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Factorial of %#v is %#v ", number, fact)
	return fact

}

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var number Number
		json.Unmarshal(reqBody, &number)
		if number.A <= 0 || number.B <= 0 {
			http.Error(w, " a & b not exist ", 400)
			return
		}

		fmt.Println("number is printing :", string(reqBody))
		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		next.ServeHTTP(w, r)
	})

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	factHandler := http.HandlerFunc(DecodeFactroial)

	http.Handle("/factroial", middlewareOne(factHandler))
	log.Fatal(http.ListenAndServe(":8989", nil))
}
