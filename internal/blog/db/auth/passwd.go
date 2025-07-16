package auth

import (
	"crypto/rand"
	"crypto/subtle"

	"golang.org/x/crypto/argon2"
)

var (
	iterations = uint32(1<<4)
	memory = uint32(1<<12)
	threads = uint8(2)
	length = uint32(64)
)

func HashPassword(password string) ([]byte, []byte, error) {
	salt := make([]byte, 32)
    _, err := rand.Read(salt)
    if (err != nil) {
        return nil, nil, err
    }

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, threads, length)

	return hash, salt, nil
}

func CheckPassword(password string, hash []byte, salt []byte) bool {
	adhoc := argon2.IDKey([]byte(password), salt, iterations, memory, threads, length)

    if len(hash) != len(adhoc) {
		return false
    }
	
	return subtle.ConstantTimeCompare(hash, adhoc) == 1
}
