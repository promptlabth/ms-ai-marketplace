package app

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    *T     `json:"data"`
}
