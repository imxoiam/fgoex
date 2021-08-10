package main

import (
	"fmt"
	"net/http"
)
func main(){
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", http.FileServer(http.Dir("static")))
}
