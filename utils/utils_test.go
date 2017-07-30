package utils

import (
	"testing"
)

func TestSignHMAC(t *testing.T) {
	sig := SignHMAC("abc", "abcde")
	Assert("matching signature", sig == "c13c92744b69681405b3af8ee2115adc0b3a7efb", t)
	sig = SignHMAC("secretkey", "f30d9d10-9d79-41c0-8ffe-874319ea39ca1501368164")
	Assert("mathes signature", sig == "548347834c75737ab8c1e0775412126294f30703", t)
}
