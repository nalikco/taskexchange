package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"regexp"
	"taskexchange"
	"taskexchange/pkg/repository"
	"time"
)

const hashSalt = "jIdYUOExRflDd5zLnawESdryvjj5pEDNsiTdYD4C97agltuyDcSsnFtjJUVwdUdV"
const usernameRegexp = "^[a-zA-Z0-9]+$"
const firstNameRegexp = "^[a-zA-Zа-яА-Я]+$"
const lastNameRegexp = "^[a-zA-Zа-яА-Я]+$"
const signingKey = "rRRCregcbAbk21yICfoySs4kVjicuAsC8Rf9zt2CZbSSiCUvKDvjipnLclRAYPXL"
const tokenTTL = 31 * 24 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Users
}

func NewAuthService(repo repository.Users) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user taskexchange.User) (taskexchange.User, error) {
	user.Password = generatePasswordHash(user.Password)

	if match := regexp.MustCompile(usernameRegexp).Match([]byte(user.Username)); match == false {
		return taskexchange.User{}, errors.New("not valid username")
	}

	if match := regexp.MustCompile(firstNameRegexp).Match([]byte(user.FirstName)); match == false {
		return taskexchange.User{}, errors.New("not valid first name")
	}

	if match := regexp.MustCompile(lastNameRegexp).Match([]byte(user.LastName)); match == false {
		return taskexchange.User{}, errors.New("not valid last name")
	}

	userByEmail, err := s.repo.GetByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return taskexchange.User{}, err
	}
	if userByEmail.Email == user.Email {
		return taskexchange.User{}, errors.New("email is already taken")
	}

	userByUsername, err := s.repo.GetByUsername(user.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return taskexchange.User{}, err
	}
	if userByUsername.Username == user.Username {
		return taskexchange.User{}, errors.New("username is already taken")
	}

	if user.Type != 1 && user.Type != 2 {
		return taskexchange.User{}, errors.New("not valid type")
	}

	return s.repo.Create(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetByEmailAndPassword(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) UpdateOnline(user taskexchange.User) error {
	user.LastOnline = time.Now()
	return s.repo.Update(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(hashSalt)))
}
