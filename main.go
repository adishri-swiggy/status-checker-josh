package main

import (
	"fmt"
	"net/http"

	"github.com/status-checker-josh/handler"
)

func main() {
	fmt.Println("Starting Server at port : 8080")
	http.HandleFunc("/", handler.Handler)
	error := http.ListenAndServe("localhost:8080", nil)
	//handling error
	if error != nil {
		fmt.Printf("Error found %v", error)
		return
	}
}
