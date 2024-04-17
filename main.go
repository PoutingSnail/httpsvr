package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(res http.ResponseWriter, req *http.Request) {
	if req != nil && req.Header != nil {
		for key, value := range req.Header {
			fmt.Printf("key: %s, value: %s \n", key, value)
			res.Header().Set(key, value)
		}
	}
	res.WriteHeader(403)
	io.WriteString(res, "ok")
}
