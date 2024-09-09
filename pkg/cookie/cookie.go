package cookie

import (
	"net/http"
	"time"
)

// SetCookie - функция для создания cookie
func SetCookie(w http.ResponseWriter, name, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Expires:  time.Now().Add(1000 * time.Hour),
		HttpOnly: false,
		Secure:   false,
		MaxAge:   3600 * 1000,
	}

	http.SetCookie(w, &cookie)
}

// GetCookie - функция для получения cookie
func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)

	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
