package main


import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello,this version 1!"))
	})
	http.HandleFunc("/bye", sayBye)

	log.Println("Starting server...v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Body)
	w.Write([]byte("Bye bye this version 1!"))
}
