package bfv

// Ciphertext is a *ring.Poly array representing a polynomial of degree > 0 where coefficients are in R_Q.
type Ciphertext struct {
	*bfvElement
}

// NewCiphertext creates a new ciphertext parameterized by degree, level and scale.
func NewCiphertext(params *Parameters, degree uint64) (ciphertext *Ciphertext) {

	if !params.isValid {
		panic("cannot NewCiphertext : params not valid (check if they where generated properly)")
	}

	return &Ciphertext{newBfvElement(params, degree)}
}

// NewCiphertextRandom generates a new uniformely distributed ciphertext of degree, level and scale.
func NewCiphertextRandom(params *Parameters, degree uint64) (ciphertext *Ciphertext) {

	if !params.isValid {
		panic("cannot NewCiphertextRandom : params not valid (check if they where generated properly)")
	}

	return &Ciphertext{newBfvElementRandom(params, degree)}
}
