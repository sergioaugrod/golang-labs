package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", time.Now())
		})

	http.ListenAndServe(":8090", nil)
}
