package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World71!")

}
func hi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi")

}

func main() {

	port := os.Getenv("PORT")
	fmt.Println(port)
	http.HandleFunc("/", hello)
	http.HandleFunc("/hi", hi)
	http.ListenAndServe(":"+port, nil)
}
