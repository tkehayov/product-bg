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

func main() {

	port := os.Getenv("PORT")
	fmt.Println(port)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
