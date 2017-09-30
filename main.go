package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func main() {
	/*
		http.HandleFunc("/", sayhelloName)       // set router
		err := http.ListenAndServe(":9090", nil) // set listen port
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	*/
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":9090", nil)

}
