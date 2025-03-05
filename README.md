# Библиотека для Хеширования Паролей (hashpassword)
Библиотека hashpassword предназначена для безопасного хеширования паролей с использованием криптографически стойкого алгоритма SHA-512. Она включает генерацию случайной соли для каждого пароля, что значительно повышает безопасность хранения данных. Библиотека также предоставляет функцию для проверки совпадения паролей.

## Как использовать?
```
func main() {
	// Пример использования
	password := "my_secure_password"

	// Хэшируем пароль
	hashedPassword, salt, err := hashpassword.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	fmt.Println("Hashed Password:", hashedPassword)
	fmt.Println("Salt:", salt)

	// Проверяем пароль
	match := hashpassword.DoPasswordsMatch(hashedPassword, "my_secure_password", salt)
	fmt.Println("Passwords match:", match) // true

	match = hashpassword.DoPasswordsMatch(hashedPassword, "wrong_password", salt)
	fmt.Println("Passwords match:", match) // false
}
```
