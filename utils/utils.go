package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

func SignHMAC(key, value string) string {
	sig := hmac.New(sha1.New, []byte(key))
	sig.Write([]byte(value))
	return hex.EncodeToString(sig.Sum(nil))
}

func Assert(description string, truth bool, t *testing.T) {
	if !truth {
		t.Errorf("%v: failed", description)
	}
}
