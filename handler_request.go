package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handlerRequest(w http.ResponseWriter, r *http.Request) {
	// get id
	input, err := getInput(r)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	// input = strings.ReplaceAll(input, "+", " ")
	response, err := callGroqAPI(input)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	respondWithJSON(w, 200, fmt.Sprintf("%v", response))
}

func getInput(r *http.Request) (string, error) {
	input := chi.URLParam(r, "input")
	return input, nil
}
