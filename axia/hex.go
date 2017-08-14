package axia

import "encoding/hex"

// ToHex coverts byte array ==> hex string
func ToHex(bytes []byte) string {
	hex := BytesToHex(bytes)
	/* EVM: Prefer output of "0x0" instead of "0x"
	if len(hex) == 0 {
		hex = "0"
	} */
	return "0x" + hex
}

// FromHex coverts hex string ==> byte array
func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		return HexToBytes(s)
	}
	return nil
}

// BytesToHex converts byte array ==> hex string
func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// HexToBytes converts hex string ==> byte array
func HexToBytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

// HexToBytesFixed ...
func HexToBytesFixed(str string, flen int) []byte {
	h, _ := hex.DecodeString(str)
	if len(h) == flen {
		return h
	} else if len(h) > flen {
		return h[len(h)-flen:]
	} else {
		hh := make([]byte, flen)
		copy(hh[flen-len(h):flen], h[:])
		return hh
	}
}

// RightPadBytes with zeroes
func RightPadBytes(slice []byte, l int) []byte {
	if l <= len(slice) {
		return slice
	}

	padded := make([]byte, l)
	copy(padded, slice)

	return padded
}

// LeftPadBytes with zeroes
func LeftPadBytes(slice []byte, l int) []byte {
	if l <= len(slice) {
		return slice
	}
	padded := make([]byte, l)
	copy(padded[l-len(slice):], slice)
	return padded
}
