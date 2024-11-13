package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserId string

func (id UserId) String() string {
	return string(id)
}

type Email string

func (e Email) String() string {
	return string(e)
}

type User struct {
	ID             UserId
	Email          Email
	HashedPassword string // захэшированный
}

func NewUser() *User {
	return &User{
		ID: UserId(uuid.New().String()),
	}
}

type Token string

func (t Token) String() string {
	return string(t)
}

func NewToken(user *User) (Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	duration := time.Hour * 24
	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	// Подписываем токен, используя секретный ключ приложения
	secretKey := "auth_service"
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return Token(tokenString), nil
}

type AuthUser struct {
	User  *User
	Token Token
}

func NewAuthUser() *AuthUser {
	return &AuthUser{}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
