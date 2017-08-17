package chain

import (
	"axia/go-axia/axia"
	"ender/ansible/chainable"
	"time"
)

// A Block is the primary storage component of a blockchain.
// It has three consistent (implementation-agnostic) interfaces:
// Author: return the creator of a block
// Parent: return the hash of the preceding block
// Transitions: return the state changes performed during this block
type Block interface {
	Author() axia.Address
	Parent() axia.Hash
	Transitions() []*chainable.Chainable

	// Fields related to networking
	ReceivedAt() time.Time
	ReceivedFrom() interface{}
}

// AxiaBlock is the default Block struct
type AxiaBlock struct {
    header Header
	author      axia.Address
	parent      axia.Hash
	transitions []*Chainable
}

func (b *AxiaBlock) Author() axia.Address {
	return b.author
}

func (b *AxiaBlock) Parent() axia.Hash {
	return b.parent
}

func (b *AxiaBlock) Transitions() []*Chainable {
	return b.transitions
}

func NewAxiaBlock(header *Header, cns []*Chainable, uncles []*Header, receipts []*Receipt) *Block {
    b := &AxiaBlock{
        header: header
    }

}
