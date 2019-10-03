package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/famasoon/gowhois/whois"
	"github.com/gorilla/mux"
)

type WhoisResponse struct {
	Info string `json:"info"`
}

type WhoisController struct{}

func (wc WhoisController) GetWhois(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domain := vars["domain"]

	info, err := whois.Whois(domain)
	if err != nil {
		respondJSON(w, err, http.StatusInternalServerError)
	}
	res := WhoisResponse{
		Info: info,
	}

	respondJSON(w, res, http.StatusOK)
}

func (wc WhoisController) GetRoot(w http.ResponseWriter, r *http.Request) {
	res := WhoisResponse{
		Info: "OK!",
	}
	respondJSON(w, res, http.StatusOK)
}

func respondJSON(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset-utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	wc := WhoisController{}
	subRouter.HandleFunc("/", wc.GetRoot)
	subRouter.HandleFunc("/whois/{domain}", wc.GetWhois)

	http.Handle("/", router)
	if err := http.ListenAndServe(":80", nil); err != nil {
		os.Exit(1)
	}
}
