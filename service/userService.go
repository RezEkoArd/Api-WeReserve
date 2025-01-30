package service

import (
	"errors"
	"os"
	"time"
	"wereserve/models"
	"wereserve/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) SignUp(name, email, password string, role models.Role) (*models.User, error) {
	//Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	//Jika role tidak tersedia, use default 
	if role == "" {
		role = models.RoleUser
	}

	//create user
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hash),
		Role:     role,
	}

	// save to db
	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (s *UserService) Login(email, password string) (string, error) {
	//? Cari User Berdasarkan emial
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	//? Verifikasi password decode hashing password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid creadentials")
	}

	//? create JWT TOKEN
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24 *  7).Unix(),
	})

	//Sign in token dengan  secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return tokenString, nil

}