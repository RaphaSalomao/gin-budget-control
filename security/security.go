package security

import (
	"errors"
	"os"
	"time"

	"github.com/RaphaSalomao/gin-budget-control/database"
	"github.com/RaphaSalomao/gin-budget-control/model/entity"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var key []byte

func LoadKey() {
	key = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tknString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tknString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, KeyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func GetLoggedUserFromToken(tokenString string) (*entity.User, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	email := token.Claims.(jwt.MapClaims)["email"]
	var user entity.User
	tx := database.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid token")
		} else {
			return nil, tx.Error
		}
	}
	return &user, nil
}

func KeyFunc(t *jwt.Token) (interface{}, error) {
	email := t.Claims.(jwt.MapClaims)["email"]
	var user entity.User
	tx := database.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid token")
		} else {
			return nil, tx.Error
		}
	}
	return []byte(key), nil
}
