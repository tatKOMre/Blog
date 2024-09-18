package middleware

type Middleware struct {
	Signkey []byte
}

func New(key []byte) *Middleware {
	return &Middleware{
}

/*
Добавь middleware для авторизации

type HandlerFunc func(w http.ResponseWriter, r *http.Request, act *token.Claims)

func (m *Middleware) Auth(f HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Тут получение токена из кукисов
		// и его проверка
		f(w, r, act)
	}
}
*/
