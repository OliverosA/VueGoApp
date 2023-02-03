package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	portFlag := flag.String("port", "8080", "Server Port Default 8080")
	flag.Parse()

	fileserver := http.FileServer(http.Dir("dist/"))

	http.Handle("/", http.StripPrefix("/", fileserver))

	fmt.Println("Serving on port: " + *portFlag)
	http.ListenAndServe(":"+*portFlag, nil)

}
