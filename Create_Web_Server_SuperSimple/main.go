package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Parseform() Err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	//	^^ tell golang to check out static directory and now golang with know
	//  that it has to look for index.html
	http.Handle("/", fileServer)
	// ^^ use Handle function to handling root route and send that
	// to the fileServer
	http.HandleFunc("/form", formHandler)   // show from.html
	http.HandleFunc("/hello", HelloHandler) // print hello

	fmt.Printf("Starting Server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
