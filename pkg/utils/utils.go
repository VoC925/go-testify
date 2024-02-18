package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// получение hash пароля SHA-256
func PasswordToHash(password string) string {
	hashPassword := sha256.Sum256([]byte(password))
	res := hex.EncodeToString(hashPassword[:])
	return res
}
