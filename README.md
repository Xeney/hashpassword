# hashpassword

## Как использовать?
```
package main

import (
	"fmt"
	"your_project/hashpassword"
)

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
