package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("method", r.Method, r.URL.String())
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Could not read body")
	} else {
		fmt.Println("body", string(b))
	}
	//w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"Hello\": \"friend\"}")
}

func main() {
	http.HandleFunc("/", respond)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
