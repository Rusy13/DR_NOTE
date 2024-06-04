package service

import "errors"

var (
	ErrOrdersIsInactive = errors.New("this user is inactive")
)
