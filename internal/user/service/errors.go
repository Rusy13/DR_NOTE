package service

import "errors"

var (
	ErrOrdersIsInactive = errors.New("this order is inactive")
)
