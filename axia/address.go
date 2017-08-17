package axia

import (
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

// Str form of Address
func (a Address) Str() string { return string(a[:]) }

// Bytes form of Address
func (a Address) Bytes() []byte { return a[:] }

// Big Integer form of Address
func (a Address) Big() *big.Int { return new(big.Int).SetBytes(a[:]) }

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
