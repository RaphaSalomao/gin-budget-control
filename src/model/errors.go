package model

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidUserPassword = errors.New("invalid user/password")
	ErrMonthDuplicated     = errors.New("already created in current month")
)
