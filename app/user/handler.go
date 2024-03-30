package user

import "github.com/promptlabth/ms-orch-user-service/app"

type usecase interface {
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewUser(c app.Context) {
	c.OK(map[string]string{
		"hello": "world",
	})
}
