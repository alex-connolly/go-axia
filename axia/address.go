package axia

import (
	"encoding/hex"
	"ender/ans/crypto/sha3"
	"fmt"
	"math/big"
)

const (
	addressLength = 20
)

// Address ...
type Address [addressLength]byte

// ToBig ...
func (a Address) ToBig() *big.Int {
	return new(big.Int).SetBytes(a[:])
}

// BytesToAddress ...
func BytesToAddress(b []byte) (a Address) {
	a.SetBytes(b)
	return a
}

// StringToAddress string ==> Address
func StringToAddress(s string) Address { return BytesToAddress([]byte(s)) }

// BigToAddress big.Int ==> Address
func BigToAddress(b *big.Int) Address { return BytesToAddress(b.Bytes()) }

// HexToAddress string ==> Address
func HexToAddress(s string) Address { return BytesToAddress(FromHex(s)) }

// IsHexAddress verifies whether a string can represent a valid hex-encoded
// Ethereum address or not.
func IsHexAddress(s string) bool {
	if len(s) == 2+2*addressLength && IsHex(s) {
		return true
	}
	if len(s) == 2*addressLength && IsHex("0x"+s) {
		return true
	}
	return false
}

// Str form of Address
func (a Address) Str() string { return string(a[:]) }

// Bytes form of Address
func (a Address) Bytes() []byte { return a[:] }

// Big Integer form of Address
func (a Address) Big() *big.Int { return new(big.Int).SetBytes(a[:]) }

// Hash form of Address
func (a Address) Hash() Hash { return BytesToHash(a[:]) }

// Hex returns an EIP55-compliant hex string representation of the address.
func (a Address) Hex() string {
	unchecksummed := hex.EncodeToString(a[:])
	sha := sha3.NewKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}

// String implements the stringer interface and is used also by the logger.
func (a Address) String() string {
	return a.Hex()
}

// Format implements fmt.Formatter, forcing the byte slice to be formatted as is,
// without going through the stringer interface used for logging.
func (a Address) Format(s fmt.State, c rune) {
	fmt.Fprintf(s, "%"+string(c), a[:])
}

// SetBytes of Address to the value of b. If b is larger than len(a) it will panic
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-addressLength:]
	}
	copy(a[addressLength-len(b):], b)
}

// Set string `s` to a. If s is larger than len(a) it will panic
func (a *Address) SetString(s string) { a.SetBytes([]byte(s)) }

// Set a to another address
func (a *Address) Set(other Address) {
	for i, v := range other {
		a[i] = v
	}
}

// MarshalText returns the hex representation of a.
func (a Address) MarshalText() ([]byte, error) {
	return hexutil.Bytes(a[:]).MarshalText()
}

// UnmarshalText parses a hash in hex syntax.
func (a *Address) UnmarshalText(input []byte) error {
	return hexutil.UnmarshalFixedText("Address", input, a[:])
}

// UnmarshalJSON parses a hash in hex syntax.
func (a *Address) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(addressT, input, a[:])
}
