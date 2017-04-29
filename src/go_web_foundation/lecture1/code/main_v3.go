package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//新建一个Server结构体实例
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: time.Second * 2,
	}

	quit := make(chan os.Signal)
	//此处发生中断时将信号添加到chan中
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler3{})
	mux.HandleFunc("/bye", sayBye3)

	//另起一个线程判断chan中信号状态，并作中断后续的处理
	go func() {
		<-quit

		if err := server.Close(); err != nil {

		}
	}()

	server.Handler = mux
	log.Print("Starting server... v3")
	err := server.ListenAndServe()
	if err != nil {
		/*if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}*/
	}
}

func sayBye3(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Body)
	w.Write([]byte("Bye bye this version 2!"))
}

type myHandler3 struct{}

//实现接口Handler
func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello v2, the request URL is: " + r.URL.String()))
}
