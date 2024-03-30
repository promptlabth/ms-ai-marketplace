package user

type storage interface {
}

type domain interface {
}

type Usecase struct {
	storage storage
	domain  domain
}

func NewUsecase(s storage, d domain) *Usecase {
	return &Usecase{storage: s, domain: d}
}

func (u *Usecase) NewUser() error {
	return nil
}
