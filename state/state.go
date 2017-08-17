package state

import "axia/go-axia/axia"

// A State stores the entirity of the current chain position.
// It is the product of all previous Transtition Chainables.
type State interface {
	Get([]byte) []byte
	Delete([]byte)
}

type StateDB struct {
}

func (s *StateDB) Get([]byte) []byte {

}

func GetByAddress(s State, a axia.Address) []byte {
	return s.Get(a.Bytes())
}
