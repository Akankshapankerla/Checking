package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

/*func sum(w http.ResponseWriter, r *http.Request) {
	x, y := 10, 20

	sum_of_values := x + y

	fmt.Fprintf(w, "Sum of 10 and 20 is %d", sum_of_values)
}*/

func main() {
	fmt.Println("Server is runining on port 8080")

	/*http.Handle("/",
		http.HandlerFunc(
			index,
		),
	)*/

	http.HandleFunc("/sum", index)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

//Create a basic HTTP server that listens on a specific port and responds with "Hello, World!" to any request.
