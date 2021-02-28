package auth

import "mekar/model"

type IAuthRepo interface {
	ReadAccountByEmail(string) (*model.Account, error)
	CreateAccount(account *model.Account, id string) error
}