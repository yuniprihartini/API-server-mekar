package model

import "strings"

type Account struct{
	AccountId string            `json:"accountId"`
	Email     string            `json:"email"`
	Password  string            `json:"password"`
	Token     string            `json:"token"`
	Errors    map[string]string `json:"errors"`
}

func (account *Account) Validate() bool {
	account.Errors = make(map[string]string)
	if strings.TrimSpace(account.Email) == "" {
		account.Errors["Email"] = "Input your email here"
	}
	if strings.TrimSpace(account.Password) == "" {
		account.Errors["Password"] = "Please enter your password here"
	}
	return len(account.Errors) == 0
}