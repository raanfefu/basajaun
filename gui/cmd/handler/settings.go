package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raanfefu/basajaun/gui/cmd/settings"
)

func GetSettings(w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)
	response := settings.GetSettings()
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(response)
	fmt.Fprintf(w, "%s", string(b))
}
