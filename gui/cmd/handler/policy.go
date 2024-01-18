package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raanfefu/basajaun/gui/cmd/policy"
	"github.com/raanfefu/basajaun/gui/pkg"
)

func PostPolicy(w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)
	var bodyPolicy pkg.Rule
	json.NewDecoder(r.Body).Decode(&bodyPolicy)
	response := policy.PostPolicy(bodyPolicy)
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(response)
	fmt.Fprintf(w, "%s", string(b))
}
