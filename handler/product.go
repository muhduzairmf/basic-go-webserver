package handler

import (
	"fmt"
	"net/http"
	"strconv"
)


func ProductHandler(w http.ResponseWriter, r *http.Request) {
	products := map[int]string{
		1: "chocolate cake",
		2: "red velvet cake",
		3: "peanut cookies",
		4: "fruits cookies",
	}
	// Make an example list of product
	
	id_queryParams := r.URL.Query().Get("id")
	// This is how to get URL query params like
	// http://localhost:8080/product?id=4

	id, err := strconv.Atoi(id_queryParams)
	// Convert id from string to integer
	
	if err != nil || id < 1 || id > len(products) {
		fmt.Fprint(w, "<h1 style='color: red'>Id product is not valid</h1>")
		return
	}
	// Validate the query params if the convertion produce an error
	// or the id number is not valid
	
	fmt.Fprintf(w, "<h1>Product Page</h1><br><h3>Product id : %v</h3><h3>Product name : %v</h3>", id, products[id])
}
