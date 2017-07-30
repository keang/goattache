package middlewares

import (
	"net/http"

	"github.com/keang/goattache/utils"
)

func Authorize(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if secret != "" {
			expectedHMAC := utils.SignHMAC(
				secret,
				r.FormValue("uuid")+r.FormValue("expiration"),
			)
			if expectedHMAC != r.Form.Get("hmac") {
				w.Header().Set("X-Exception", "Authorization failed")
				http.Error(w, "Authorization failed", http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
