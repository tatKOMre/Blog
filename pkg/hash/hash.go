package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

// Hash - функция для хэширования строки с помощью SHA-1 алгоритма
func Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
