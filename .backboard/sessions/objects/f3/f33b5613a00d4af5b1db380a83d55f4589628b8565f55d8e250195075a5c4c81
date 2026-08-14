package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiHandler)

	fmt.Println("Server running at http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		fmt.Println("Server error:", err)
		return
	}
}