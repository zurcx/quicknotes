package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("teste")

	http.ListenAndServe(":5000", nil)
}
