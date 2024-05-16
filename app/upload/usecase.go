package upload

import (
	"context"
	"mime/multipart"
)

type gcs interface {
	Upload(context.Context, multipart.File, string) error
}

type Usecase struct {
	gcs gcs
}

func NewUsecase(g gcs) *Usecase {
	return &Usecase{
		gcs: g,
	}
}

func (u *Usecase) Uploadfile(ctx context.Context, file multipart.File, filename string) (*string, error) {
	m := "error"
	_, err := m, u.gcs.Upload(ctx, file, filename)

	return &m, err
}
