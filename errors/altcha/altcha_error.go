package altchaError

type AltchaNotMatchingError struct {
  Message string
}

func (err AltchaNotMatchingError) Error() string {
  return err.Message
}

