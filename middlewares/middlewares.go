package middlewares

import (
	"log"
	"net/http"

	"github.com/keang/goattache/utils"
)

func Authorize(secret string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if secret != "" {
			r.ParseForm()
			expectedHMAC := utils.SignHMAC(secret,
				r.Form.Get("uuid")+r.Form.Get("expiration"),
			)
			if expectedHMAC != r.Form.Get("hmac") {
				w.Header().Set("X-Exception", "Authorization failed")
				http.Error(w, "Authorization failed", 401)
				log.Printf("%+v", 401)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
