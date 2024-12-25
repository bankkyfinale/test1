package handler

import (
	"fmt"
	"net/http"
)

// HelloHandler handles requests to the root endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
