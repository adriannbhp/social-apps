package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte(`{"status": "ok"}`))

	data := map[string]string{
		"status":  "OK",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		// error
		writeJSONError(w, http.StatusInternalServerError, "err.Error()")
	}
}
