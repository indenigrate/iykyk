package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handlerRequest(w http.ResponseWriter, r *http.Request) {
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
	respondWithText(w, 200, response+"\n")
}

func getInput(r *http.Request) (string, error) {
	input := chi.URLParam(r, "input")
	return input, nil
}

func (apiCfg *apiConfig) handlerChatRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 400, "Error reading body")
		return
	}
	input := string(body)
	fmt.Printf("%+v\n", input)
	response, err := callGroqAPI(input)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	respondWithText(w, 200, response+"\n")
}
