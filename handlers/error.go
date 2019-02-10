package handlers

import (
	"net/http"
	"encoding/json"
)

func sendError(res http.ResponseWriter, err error, code int) {
	marshaledError, err := json.Marshal(err.Error())
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(code)
	res.Write(marshaledError)
}
