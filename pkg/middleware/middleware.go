package middleware

type Middleware struct {
	Signkey []byte
}

func New(key []byte) *Middleware {
	return &Middleware{
		Signkey: key,
	}
}
