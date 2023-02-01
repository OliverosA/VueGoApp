package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileserver := http.FileServer(http.Dir("dist/"))

	http.Handle("/", http.StripPrefix("/", fileserver))

	fmt.Println("Serving on port: 8080")
	http.ListenAndServe(":8080", nil)

}
