package altcha

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
)

func CreateALTCHA(salt string, number big.Int) (*models.Challenge, error) {
	if salt == "" {
		saltBytes := make([]byte, 12)
		_, err := rand.Read(saltBytes)
		if err != nil {
			return nil, err
		}
		salt = hex.EncodeToString(saltBytes)
	}

  randomNum := big.NewInt(0)

  if number.Cmp(big.NewInt(0)) == 0 {
    var err error
    // Generate random number
    randomNum, err = RandomInt(
      *models.ALTCHA_NUM_RANGE[0], 
      *models.ALTCHA_NUM_RANGE[1],
    )
    if err != nil { return nil, err }
  } else {
    randomNum = &number
  }
  
  randomStr := salt + randomNum.String()
	h := sha256.New()
	h.Write([]byte(randomStr))
	challenge := hex.EncodeToString(h.Sum(nil))

	mac := hmac.New(sha256.New, []byte(models.ALTCHA_HMAC_KEY))
	mac.Write([]byte(challenge))
	signature := hex.EncodeToString(mac.Sum(nil))

	return &models.Challenge{
		Algorithm: models.ALTCHA_ALG,
		Challenge: challenge,
		Salt: salt,
		Signature: signature,
		Number: *randomNum,
	}, nil
}

func VerifyALTCHA(payload models.SubscriberWithChallenge) (bool, error) {
  decodedStr, err := base64.StdEncoding.DecodeString(*payload.Altcha)
  var challenge models.Challenge
  err = json.Unmarshal([]byte(decodedStr), &challenge)
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Error occurred when parsing the JSON challenge.",
    )
    return false, err
  }

	check, err := CreateALTCHA(challenge.Salt, challenge.Number)
	if err != nil {
    logs.Output(
      logs.ERROR,
      "Error occurred when creating an Altcha challenge.",
    )
		return false, err
	}

	return challenge.Algorithm == check.Algorithm &&
		challenge.Challenge == check.Challenge &&
		challenge.Signature == check.Signature, nil
}

