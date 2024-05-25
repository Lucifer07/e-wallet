package util

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Lucifer07/e-wallet/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type HelperInf interface {
	HashPassword(pwd string, cost int) ([]byte, error)
	CheckPassword(pwd string, hash []byte) (bool, error)
	CreateAndSign(user entity.User) (string, error)
}
type HelperImpl struct {
}

func (h *HelperImpl) HashPassword(pwd string, cost int) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)

	if err != nil {

		return nil, err

	}

	return hash, nil

}
func (h *HelperImpl) CheckPassword(pwd string, hash []byte) (bool, error) {

	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))

	if err != nil {

		return false, err

	}

	return true, nil

}

func (h *HelperImpl) CreateAndSign(user entity.User) (string, error) {
	godotenv.Load()
	app := os.Getenv("APP_NAME")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"iss":   app,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
		"iat":   time.Now(),
		"email": user.Email,
	})
	signatur := os.Getenv("JWT_SIGNATURE_KEY")
	signed, err := token.SignedString([]byte(signatur))

	if err != nil {

		return "", err

	}

	return signed, nil

}
func CheckClaim(ctx context.Context) (*entity.User, error) {
	data := ctx.Value("data").(map[string]string)
	if len(data) == 0 {
		return nil, ErrorUnauthorized
	}
	userId, ok := data["id"]
	if !ok && userId == "" {
		return nil, ErrorBadRequest
	}
	userEmail, ok := data["email"]
	if !ok && userEmail == "" {
		return nil, ErrorBadRequest
	}
	id, _ := strconv.Atoi(userId)
	return &entity.User{
		Id:    id,
		Email: userEmail,
	}, nil
}

func ParseAndVerify(signed string) (jwt.MapClaims, error) {
	app := os.Getenv("APP_NAME")
	signatur := os.Getenv("JWT_SIGNATURE_KEY")
	token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
		return []byte(signatur), nil

	}, jwt.WithIssuer(app),

		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),

		jwt.WithExpirationRequired(),
	)

	if err != nil {

		return nil, err

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		return claims, nil

	} else {

		return nil, fmt.Errorf("unknown claims")

	}

}

func RandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), nil
}
