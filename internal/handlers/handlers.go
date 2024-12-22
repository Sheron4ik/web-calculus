package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sheron4ik/web-calculus/pkg/calculus"
)

type Body struct {
	Expr string `json:"expression"`
}

func CalculusHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	Body := Body{}
	err := json.NewDecoder(r.Body).Decode(&Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ERROR: %s", "internal server error")
		return
	}

	result, err := calculus.Calc(Body.Expr)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "ERROR: %s", err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Result: %f", result)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalculusHandler)
	return LoggingMiddleware(mux)
}
