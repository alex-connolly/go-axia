package axia

import "testing"

func TestHashString(t *testing.T) {
	s := "hello"
	HashString(s)
	s = "thisstringislongerthanthehaslengthbecauseitis"
	HashString(s)
}

func TestHashBytes(t *testing.T) {
	s := "hello"
	HashBytes([]byte(s))

}
