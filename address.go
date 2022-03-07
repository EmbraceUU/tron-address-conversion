package tron_address_conversion

type Address []byte

func (a Address) String() string {
	return EncodeCheck(a)
}
