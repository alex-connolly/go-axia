package encoding

// Protocol ...
type Protocol interface {
	Encode([]byte) []byte
	Decode([]byte) []byte
}
