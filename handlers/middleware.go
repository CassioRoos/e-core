package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func MethodHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && !strings.Contains(r.URL.Path, "health")  {
			message := fmt.Sprintf("Method %s not allowed. Only `%s` is supported", r.Method, http.MethodPost)
			http.Error(w, message, http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
