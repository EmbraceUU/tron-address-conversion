package conversion

import (
	"strings"
	"testing"
)

func TestDefaultConversion(t *testing.T) {
	cv := NewConversion()
	baseAddr, err := cv.HexToBase58(strings.ToLower("0x0000000000000000000000000000000000000000"))
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log(baseAddr)
	}

	hexAddr, err := cv.Base58ToHex("T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(hexAddr)
	}
}
