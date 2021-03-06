package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/keang/goattache/testutils"
	"github.com/keang/goattache/utils"
)

func dummyHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pass")
	})
}

func TestNoAuthorization(t *testing.T) {
	assert := testutils.Assert{t}
	rr := httptest.NewRecorder()
	handler := Authorize("", dummyHandler())

	req := httptest.NewRequest("POST", "/example", nil)
	handler.ServeHTTP(rr, req)
	assert.Equal(rr.Code, http.StatusOK)
	assert.Equal(rr.Header().Get("X-Exception"), "")
	body, _ := ioutil.ReadAll(rr.Body)
	assert.Equal(string(body), "pass")
}

func TestInvalidSignature(t *testing.T) {
	assert := testutils.Assert{t}
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
	assert.Equal(rr.Code, http.StatusUnauthorized)
	assert.Equal(rr.Header().Get("X-Exception"), "Authorization failed")
}

func TestValidAuthorization(t *testing.T) {
	assert := testutils.Assert{t}
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
	assert.Equal(rr.Code, http.StatusOK)
	assert.Equal(rr.Header().Get("X-Exception"), "")
	body, _ := ioutil.ReadAll(rr.Body)
	assert.Equal(string(body), "pass")
}
