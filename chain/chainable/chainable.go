package chainable

// A Chainable is a piece of data which can be bundled into a block
type Chainable interface {
	Sign()
}
