package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s", address)
}
func thankshandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/thanks" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported ", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "This website is Developed By MR. Arvind Kumar rana!")
	fmt.Fprintf(w, "Web Server is written by Shoib Ansari")
}

func main() {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/thanks", thankshandler)
	fmt.Printf("Starting server at 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
