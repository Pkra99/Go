package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)
const PORT = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Hello world!")
	})

	log.Printf("Server listening on :%d\n", PORT)
	if err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil); err != nil{
		log.Fatalf("Server error: %v", err)
	}
}
