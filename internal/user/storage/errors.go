package storage

import "errors"

var (
	ErrOrderNotFound       = errors.New("no orders with such id")
	ErrDuplicateFeatureTag = errors.New("such pair of feature and tag already exists")
	ErrDuplicateItem       = errors.New("duplicate item")
)
