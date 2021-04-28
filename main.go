package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Fatalln(http.ListenAndServe(":8888", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		hs := req.Header.Get("X-Forwarded-For")
		log.Println("X-Forwarded-For headers:", hs)
		rw.WriteHeader(http.StatusOK)
		io.WriteString(rw, hs)
		for k, v := range req.Header {
			io.WriteString(rw, fmt.Sprintf("%s:%s\n", k, strings.Join(v, ",")))
		}
	})))
}
