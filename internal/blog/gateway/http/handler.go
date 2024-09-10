package http

import (
	"tatKOM/internal/blog"
)

type Handler struct {
	Service blog.Service
	SignKey []byte
}

func New(service blog.Service, key []byte) *Handler {
	return &Handler{
		Service: service,
		SignKey: key,
	}
}

/*
Напишу тут, ибо во всех функциях gateway слоя надо пофиксить

Вместо работы с json надо переделать все на работу с html,
слои сервиса и репозитория готовы.

Возможно придется объединить некоторые функции в одну
например, для странички с постом нужно получить сам пост и комменты,
тоесть функции GetCommentsFor и GetPost совместить в одну функцию,
тут надо продумать какие страницы будут на сайте,
и какие данные надо получать из бд для каждой из них

Работа с html формами

Сначала:
r.ParseForm()

Потом получаешь значение по названию поля
r.FormValue("fuck")
*/
