package handlers

import (
	"fmt"
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}
