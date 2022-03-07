package tron_address_conversion

import (
	"strings"
	"testing"
)

func TestDefaultConversion(t *testing.T) {
	cv := NewConversion()
	baseAddr, err := cv.HexToBase58(strings.ToLower("0xa614f803B6FD780986A42c78Ec9c7f77e6DeD13C"))
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log(baseAddr)
	}

	hexAddr, err := cv.Base58ToHex("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(hexAddr)
	}
}
