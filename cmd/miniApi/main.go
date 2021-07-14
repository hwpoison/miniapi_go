package main

import (
	"log"
	"net/http"
	"miniApi/pkg/server"
	"time"
)

func main() {
	log.Println("Initiaizing server...")
	s := server.New()
	addr := "127.0.0.1:8082"
	server := &http.Server{
		Addr				:addr,
		Handler 			:s.Router(),
		ReadTimeout			: 10 * time.Second,
		WriteTimeout		: 10 * time.Second,
		MaxHeaderBytes		: 1 << 20,
	}
	log.Println(addr)
	log.Fatal(server.ListenAndServe())
}