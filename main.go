package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/famasoon/gowhois/whois"
	"github.com/gorilla/mux"
	"https://github.com/uber-go/zap"
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

func logging(logger *zap.logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResoponseWriter, r *http.Request)) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = newRequestID()
			}

			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			logger.Info("", zap.String("requestID", requestID), zap.String("METHOD", r.Method), zap.String("Address", r.RemoteAddr), zap.String("UserAgent", r.UserAgent()))
			next.ServeHTTP(w, r.WithContext(ctx))
		}
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
