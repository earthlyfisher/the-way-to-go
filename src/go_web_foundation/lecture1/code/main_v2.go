package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye2)

	log.Println("Starting server...v2")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

func sayBye2(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Body)
	w.Write([]byte("Bye bye this version 2!"))
}

type myHandler struct{}

//实现接口Handler
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello v2, the request URL is: " + r.URL.String()))
}
