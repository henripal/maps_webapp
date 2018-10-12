package main

import (
	"io"
	"net/http"

	"github.com/rs/cors"
)

func data(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "WHAT IS THE TEMPLATE????")
}

func main() {
	http.HandleFunc("/data", data)
	http.ListenAndServe(":8888", cors.Default().Handler(http.DefaultServeMux))
}
