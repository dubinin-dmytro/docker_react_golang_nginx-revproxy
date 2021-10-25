package main

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	rs := Response{
		Result: true,
		Data:   data,
	}
	makeResponse(w, rs, http.StatusOK)
}

func ErrorResponse(w http.ResponseWriter, err error) {
	rs := Response{
		Result:  false,
		Message: err.Error(),
	}

	makeResponse(w, rs, http.StatusInternalServerError)
}

func makeResponse(w http.ResponseWriter, rs Response, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(rs)
}
