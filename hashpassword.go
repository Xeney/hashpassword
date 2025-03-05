package hashpassword

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

const (
	saltSize = 16 // Размер соли в байтах
)

// HashPassword возвращает хэш пароля с использованием соли.
// Возвращает хэш в виде hex-строки и соль.
func HashPassword(password string) (string, []byte, error) {
	if password == "" {
		return "", nil, errors.New("password cannot be empty")
	}

	// Генерация соли
	salt := GenerateRandomSalt()

	// Создание хэша с использованием SHA-512
	hash := sha512.New()
	hash.Write([]byte(password))
	hash.Write(salt)
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	return hashedPassword, salt, nil
}

// DoPasswordsMatch проверяет, совпадает ли текущий пароль с хэшированным паролем.
func DoPasswordsMatch(hashedPassword, currPassword string, salt []byte) bool {
	if hashedPassword == "" || currPassword == "" || len(salt) == 0 {
		return false
	}

	// Хэшируем текущий пароль с той же солью
	hash := sha512.New()
	hash.Write([]byte(currPassword))
	hash.Write(salt)
	currPasswordHash := hex.EncodeToString(hash.Sum(nil))

	// Сравниваем хэши
	return hashedPassword == currPasswordHash
}

// GenerateRandomSalt генерирует случайную соль заданного размера.
func GenerateRandomSalt() []byte {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err) // В реальном приложении лучше обработать ошибку
	}
	return salt
}
