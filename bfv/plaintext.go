package bfv

import (
	"github.com/ldsec/lattigo/ring"
)

// Plaintext is a BigPoly of degree 0.
type Plaintext struct {
	*bfvElement
	value *ring.Poly
}

// NewPlaintext creates a new plaintext from the target context.
func NewPlaintext(params *Parameters) *Plaintext {

	if !params.isValid {
		panic("cannot NewPlaintext : params not valid (check if they where generated properly)")
	}

	plaintext := &Plaintext{newBfvElement(params, 0), nil}
	plaintext.value = plaintext.bfvElement.value[0]
	plaintext.isNTT = false
	return plaintext
}
