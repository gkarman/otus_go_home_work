package domain

import "errors"

var (
	ErrEntityAlreadyExists = errors.New("entity already exists")
	ErrEntityNotFound      = errors.New("entity not found")
)
