package models

import (
	"math/big"
	"os"
)

const ALTCHA_ALG = "SHA-256"
var ALTCHA_HMAC_KEY string = os.Getenv("ALTCHA_SECRET")
var ALTCHA_NUM_RANGE = [2]*big.Int{big.NewInt(1e3), big.NewInt(1e5)}

type Challenge struct {
	Algorithm string  `json:"algorithm"`
	Challenge string  `json:"challenge"`
	Salt      string  `json:"salt"`
	Signature string  `json:"signature"`
	Number    big.Int `json:"number"`
}

