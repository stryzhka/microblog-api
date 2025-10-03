package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"microblog-api/auth"
	"microblog-api/auth/util"
	"microblog-api/models"
	"microblog-api/profile"
	"regexp"
	"strings"
	"time"
)

type UserService struct {
	repo           auth.Repository
	profileService profile.Service
	passwordSalt   string
	signingKey     string
	expire         time.Duration
}

func NewUserService(
	repo auth.Repository,
	profileService profile.Service,
	passwordSalt string,
	signingKey string,
	expire time.Duration) *UserService {
	return &UserService{
		repo:           repo,
		profileService: profileService,
		passwordSalt:   passwordSalt,
		signingKey:     signingKey,
		expire:         expire,
	}
}

func (s *UserService) Signup(username, password string) error {
	valid := regexp.MustCompile("^[a-zA-Z0-9]+$")
	if !valid.MatchString(username) || strings.TrimSpace(password) == "" {
		return auth.ErrValidation
	}
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
	if err == nil {
		err = s.profileService.Create(id.String(), username, "")
	}

	return err
}

func (s *UserService) Signin(username, password string) (string, error) {
	user, err := s.repo.Get(username, util.GeneratePasswordHash(password, s.passwordSalt))
	if err != nil {
		return "", err
	}
	return util.GenerateToken(user.Id, user.Role, s.signingKey)
}

func (s *UserService) ParseToken(accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.signingKey), nil
	})

	if err != nil {
		//fmt.Println(err.Error())
		return nil, auth.ErrInvalidToken
	}

	if claims, ok := token.Claims.(*auth.UserClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidToken
}
