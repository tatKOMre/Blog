package service

import "tatKOM/internal/blog"

type Service struct {
	Repository blog.Repository // интерфейс репозитория
	SignKey    []byte          // ключ шифрования jwt токенов
}

func New(repo blog.Repository, key []byte) *Service {
	return &Service{
		Repository: repo,
		SignKey:    key,
	}
}
