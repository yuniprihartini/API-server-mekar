package auth

import "mekar/model"

type IAuthUsecase interface {
	Login(account *model.Account) (*model.Account, error)
	Register(account *model.Account) (*model.Account, error)
}