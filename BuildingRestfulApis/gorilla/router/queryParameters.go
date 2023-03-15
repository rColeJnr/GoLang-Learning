package router

import (
	"fmt"
	"net/http"
)

func QParamsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch q params as a map
	qParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got param id:%s!\n", qParams["id"][0])
	fmt.Fprintf(w, "Got param category:%s!\n", qParams["id"][0])
}
