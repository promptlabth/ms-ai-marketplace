package app

import "net/http"

type Context interface {
	Bind(v any) error
	Param(key string) string
	OK(v any)
	BadRequest(err error)
	StoreError(err error)
	WriteHeader(key, value string)
	Request() *http.Request
	Next()
	AbortWithStatus(code int)
}

type HandlerFunc func(Context)
