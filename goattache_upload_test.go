package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	"github.com/keang/goattache/middlewares"
	"github.com/keang/goattache/store"
	"github.com/keang/goattache/testutils"
	"github.com/keang/goattache/utils"
)

func TestUploadHandler(t *testing.T) {
	assert := testutils.Assert{t}
	rr := httptest.NewRecorder()
	g := Goattache{Store: store.Disk{"tmp"}}
	filename := "images/Exãmpl%e _1.234 _20.png"
	secret := "topsecret"
	handler := middlewares.Authorize(secret, http.HandlerFunc(g.UploadHandler))

	expiration := fmt.Sprintf("%v", time.Now().Add(3*time.Hour).Unix())
	hmac := utils.SignHMAC(secret, "random"+expiration)
	queryParams := url.Values{"expiration": {expiration},
		"uuid": {"random"},
		"hmac": {hmac},
		"file": {filename},
	}
	image, err := ioutil.ReadFile("fixtures/" + filename)
	if err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(
		"POST",
		"/upload?"+queryParams.Encode(),
		bytes.NewBuffer(image),
	)

	handler.ServeHTTP(rr, req)

	res := store.SavedFile{}
	jsonBytes, _ := ioutil.ReadAll(rr.Body)
	err = json.Unmarshal(jsonBytes, &res)
	assert.Nil(err, t)
	assert.Equal(filepath.Base(res.Path), "Exãmpl_e _1.234 _20.png")
	assert.Equal(res.Geometry, "20x16")
	assert.Equal(res.ContentType, "image/png")
	assert.Equal(res.Size, int64(2212))
	assert.Equal(rr.Code, http.StatusCreated)
}
