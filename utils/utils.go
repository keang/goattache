package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"
)

func SignHMAC(key, value string) string {
	sig := hmac.New(sha1.New, []byte(key))
	sig.Write([]byte(value))
	return hex.EncodeToString(sig.Sum(nil))
}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandString(n int) string {
	b := make([]byte, n)
	randSrc := rand.NewSource(time.Now().UnixNano())
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
