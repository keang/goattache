package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/keang/goattache/utils"
)

func TestNoAuthorization(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)

	secret := os.Getenv("ATTACHE_SECRET_KEY")
	if err := os.Setenv("ATTACHE_SECRET_KEY", ""); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/upload", nil)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Error("Wrong status code.")
	}

	if err := os.Setenv("ATTACHE_SECRET_KEY", secret); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidSignature(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)

	// signed with "wrong#{key}"
	hmac := utils.SignHMAC("wrong"+os.Getenv("ATTACHE_SECRET_KEY"), "random1501349484")
	values := url.Values{"expiration": {"1501349484"}, "uuid": {"random"}, "hmac": {hmac}}
	req := httptest.NewRequest("POST",
		"/upload",
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

func TestValidSignature(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)
	// valid hmac
	hmac := utils.SignHMAC(os.Getenv("ATTACHE_SECRET_KEY"), "random1501349484")
	values := url.Values{"expiration": {"1501349484"}, "uuid": {"random"}, "hmac": {hmac}}
	req := httptest.NewRequest("POST",
		"/upload",
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
}
