package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)

}
