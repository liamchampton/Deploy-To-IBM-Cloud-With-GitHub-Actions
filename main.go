package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create the route handler listening on '/'
	http.HandleFunc("/", Home)
	fmt.Println("Starting server on port 8080")

	// Start the sever
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Assign the 'msg' variable with a string value
	msg := "Hello everybody!"

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf(msg)))
}
