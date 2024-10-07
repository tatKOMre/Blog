package service

import "errors"

var (
	errNotPermission = errors.New("Недостаточно прав")
	wrongpass        = errors.New("Неверный пароль")
)
