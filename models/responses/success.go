package responses

type SuccessResponse[T any] struct {
  Data T `json:"data"`
  Message string `json:"message"`
}
