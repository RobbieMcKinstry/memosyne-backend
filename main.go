package main

import (
	"log"
	. "net/http"
)

func main() {
	log.Println("Hello Memosyne")
	HandleFunc("/", HandleSessionRead)
	log.Println("Now listening on port 8080")
	ListenAndServe(":8080", nil)
}

func HandleSessionRead(w ResponseWriter, r *Request) {
	w.Write([]byte("Hello, Memosyne!"))
}
