package axia

const (
	hashLength = 30
)

// Hash ...
type Hash [hashLength]byte

// HashString ...
func HashString(str string) (h Hash) {
	h.SetString(str)
	return h
}

// HashBytes ...
func HashBytes(bytes []byte) (h Hash) {
	h.SetBytes(bytes)
	return h
}

// SetBytes ...
func (h Hash) SetBytes(bytes []byte) {
	if len(bytes) > hashLength {
		bytes = bytes[len(bytes)-hashLength:]
	}
	copy(h[hashLength-len(bytes):], bytes)
}

// SetString ...
func (h Hash) SetString(str string) {
	h.SetBytes([]byte(str))
}
