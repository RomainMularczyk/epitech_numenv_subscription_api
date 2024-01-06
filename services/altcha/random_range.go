package altcha

import (
	cryptoRand "crypto/rand"
	"math/big"
	mathRand "math/rand"
	"numenv_subscription_api/errors/logs"
	"time"
)

// Gets a random int from a range
func RandomInt(first big.Int, second big.Int) (*big.Int, error) {
  mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
  diff := new(big.Int).Sub(&second, &first)
  randomNum, err := cryptoRand.Int(cryptoRand.Reader, diff)
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not generate random number.",
    )
    return nil, err
  }
  return randomNum, nil
}

