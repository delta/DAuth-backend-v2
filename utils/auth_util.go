package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	config "github.com/delta/DAuth-backend-v2/config/impl"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID int64) (string, error) {
	config := config.New()
	jwtSecret := config.Get("JWT_SECRET")
	tokenLifespan, err := strconv.Atoi(config.Get("TokenHourLifeSpan"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(userToken string) (int64, error) {
	if userToken != "" {
		config := config.New()
		jwtSecret := config.Get("JWT_SECRET")
		token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return 0, errors.New("Unexpected signing method")
			}
			return []byte(jwtSecret), nil

		})
		if err != nil {
			return 0, err
		}
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return 0, errors.New("Invalid JWT")
			}

			userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["userID"]), 10, 32)
			if err != nil {
				return 0, err
			}
			return int64(userID), nil
		}
		return 0, err
	}
	return 0, errors.New("No Token")
}

func IsValidPhoneNumber(phoneNumber string) bool {
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

	return re.Find([]byte(phoneNumber)) != nil
}
