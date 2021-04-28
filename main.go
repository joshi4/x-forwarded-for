package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Fatalln(http.ListenAndServe(":8888", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		hs := req.Header.Get("X-Forwarded-For")
		log.Println("X-Forwarded-For headers:", hs)
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, hs)
	})))
}
