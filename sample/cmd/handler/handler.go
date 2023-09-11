package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{ \"health\": \"ok\" }")
}

func EcoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(b))
}
