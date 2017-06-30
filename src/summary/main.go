package main

import (
	"summary/utils"
	"summary/controller"
	"net/http"
	"log"
)

func main() {
	startServer()
}

func startServer() {
	logger := utils.NewFileLogger()
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to page ....."))
	})
	mux.Handle("/customer", controller.NewCustomerHandler())

	logger.Info("Start server..........\n")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
