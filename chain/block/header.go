package chain

import (
	"axia/go-axia/axia"
	"math/big"
)

// Header for a block
type Header interface {
	ParentHash() axia.Hash
	Coinbase() axia.Address
	Root() axia.Hash
	Time() *big.Int
}
