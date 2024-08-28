// adaptor.go

package user

import (
	firebase "firebase.google.com/go/v4"
)

type UserAdaptor struct {
	firebase *firebase.App
}

func NewUserAdaptor(
	firebase *firebase.App) *UserAdaptor {
	return &UserAdaptor{
		firebase: firebase,
	}
}
