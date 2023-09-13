package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raanfefu/basajaun/admission/cmd/k8s"
)

func Mutation(w http.ResponseWriter, r *http.Request) {
	e := k8s.DecodeAdmissionReview(r)
	w.Header().Set("Content-Type", "application/json")
	admissionReviewOut := PodMutationHandler(&e)
	ojson, _ := json.MarshalIndent(admissionReviewOut, "", "   ")
	rBody := string(ojson)
	fmt.Fprintf(w, rBody)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{ \"health\": \"ok\" }")
}
