package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/keang/goattache/utils"
)

func dummyHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pass")
	})
}

func TestNoAuthorization(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := Authorize("", dummyHandler())

	req := httptest.NewRequest("POST", "/example", nil)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("Wrong status code")
	}
	if hException := rr.Header().Get("X-Exception"); hException != "" {
		t.Error("Wrong X-Exception header.")
	}
	if body, _ := ioutil.ReadAll(rr.Body); string(body) != "pass" {
		t.Error("Wrong response body")
	}
}

func TestInvalidSignature(t *testing.T) {
	rr := httptest.NewRecorder()
	secret := "secretkey"
	handler := Authorize(secret, dummyHandler())

	hmac := utils.SignHMAC("wrongkey", "random1501349484")
	values := url.Values{"expiration": {"1501349484"}, "uuid": {"random"}, "hmac": {hmac}}
	req := httptest.NewRequest("POST",
		"/example",
		strings.NewReader(values.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Error("Wrong status code.")
	}
	if hException := rr.Header().Get("X-Exception"); hException != "Authorization failed" {
		t.Error("Wrong X-Exception header.")
	}
}

func TestValidAuthorization(t *testing.T) {
	rr := httptest.NewRecorder()
	secret := "secretkey"
	handler := Authorize(secret, dummyHandler())

	// valid hmac
	hmac := utils.SignHMAC(secret, "random1501349484")
	values := url.Values{"expiration": {"1501349484"}, "uuid": {"random"}, "hmac": {hmac}}
	req := httptest.NewRequest("POST",
		"/example",
		strings.NewReader(values.Encode()),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("Wrong status code.")
	}
	if hException := rr.Header().Get("X-Exception"); hException != "" {
		t.Error("Wrong X-Exception header.")
	}
	if body, _ := ioutil.ReadAll(rr.Body); string(body) != "pass" {
		t.Error("Wrong response body")
	}
}
