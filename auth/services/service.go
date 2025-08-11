package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"microblog-api/auth"
	"microblog-api/auth/util"
	"microblog-api/models"
	"time"
)

type UserService struct {
	repo         auth.Repository
	passwordSalt string
	signingKey   string
	expire       time.Duration
}

func NewUserService(
	repo auth.Repository,
	passwordSalt string,
	signingKey string,
	expire time.Duration) *UserService {
	return &UserService{
		repo:         repo,
		passwordSalt: passwordSalt,
		signingKey:   signingKey,
		expire:       expire,
	}
}

func (s *UserService) Signup(username, password string) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	err = s.repo.Create(&models.User{
		Id:       id.String(),
		Username: username,
		Password: util.GeneratePasswordHash(password, s.passwordSalt),
		Role:     "auth",
	})
	return err
}

func (s *UserService) Signin(username, password string) (string, error) {
	user, err := s.repo.Get(username, password)
	if err != nil {
		return "", err
	}
	return util.GenerateToken(user.Id, user.Role, s.signingKey)
}

func (s *UserService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return s.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*auth.UserClaims); ok && token.Valid {
		return claims.Id, nil
	}

	return "", errors.New("invalid token")
}
