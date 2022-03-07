package tron_address_conversion

import "github.com/ethereum/go-ethereum/common"

type Conversion interface {
	Base58ToHex(s string) (string, error)
	HexToBase58(s string) (string, error)
}

type DefaultConversion struct{}

func NewConversion() *DefaultConversion {
	return new(DefaultConversion)
}

func (c DefaultConversion) Base58ToHex(s string) (string, error) {
	addr, err := DecodeCheck(s)
	if err != nil {
		return "", err
	}
	return common.BytesToAddress(addr[len(addr)-20:]).String(), nil
}

func (c DefaultConversion) HexToBase58(s string) (string, error) {
	addr, err := HexToAddress(s)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}
