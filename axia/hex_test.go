package axia

import (
	"testing"

	"github.com/end-r/goutil"
)

func TestRightPadBytesShort(t *testing.T) {
	s := []byte("111")
	s = RightPadBytes(s, 10)
	goutil.Assert(t, len(s) == 10, "incorrect padding length")
}

func TestLeftPadBytesShort(t *testing.T) {
	s := []byte("111")
	s = LeftPadBytes(s, 10)
	goutil.Assert(t, len(s) == 10, "incorrect padding length")
}

func TestRightPadBytesEqual(t *testing.T) {
	s := []byte("1111111111")
	s = RightPadBytes(s, 10)
	goutil.Assert(t, len(s) == 10, "incorrect padding length")
}

func TestLeftPadBytesEqual(t *testing.T) {
	s := []byte("1111111111")
	s = LeftPadBytes(s, 10)
	goutil.Assert(t, len(s) == 10, "incorrect padding length")
}

func TestRightPadBytesLong(t *testing.T) {
	s := []byte("1111111111AAAAA")
	s = RightPadBytes(s, 10)
	goutil.Assert(t, len(s) == 15, "incorrect padding length")
}

func TestLeftPadBytesLong(t *testing.T) {
	s := []byte("1111111111AAAAA")
	s = LeftPadBytes(s, 10)
	goutil.Assert(t, len(s) == 15, "incorrect padding length")
}
