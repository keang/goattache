package main

import (
	"image"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/keang/goattache/store"
	"github.com/keang/goattache/testutils"
)

func TestDownloadResize(t *testing.T) {
	assert := testutils.Assert{t}
	rr := httptest.NewRecorder()
	g := Goattache{Store: store.Disk{"fixtures"}}
	handler := http.HandlerFunc(g.DownloadHandler)
	path := "/view/images/16x16/download.png"
	u, err := url.Parse(path)
	if err != nil {
		t.Error(err)
	}
	req := httptest.NewRequest(
		"GET",
		u.EscapedPath(),
		nil,
	)
	handler.ServeHTTP(rr, req)
	result := rr.Result()
	assert.Equal(result.Header.Get("Content-Type"), "image/png")
	conf, format, e := image.DecodeConfig(rr.Body)
	assert.Equal(e, nil)
	assert.Equal(conf.Width, 16)
	assert.Equal(conf.Height, 16)
	assert.Equal(format, "png")
}
