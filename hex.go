package tron_address_conversion

import "encoding/hex"

func HexToAddress(s string) (Address, error) {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}

	if s[:2] != "41" {
		s = "41" + s[:]
	}

	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}
