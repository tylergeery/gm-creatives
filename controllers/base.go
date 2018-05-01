package controllers

import (
	"encoding/json"
	"net/http"
)

type object interface {
	validate() error
}

func reject(w http.ResponseWriter, obj object) bool {
	err := obj.validate()
	if err != nil {
		emitErrorAsJSON(w, err)
		return true
	}

	return false
}

func emitErrorAsJSON(w http.ResponseWriter, err error) {
	response := map[string]string{
		"error": err.Error(),
	}

	w.WriteHeader(http.StatusBadRequest)
	emitAsJSON(w, response)
}

func emitAsJSON(w http.ResponseWriter, j interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(j)
}
