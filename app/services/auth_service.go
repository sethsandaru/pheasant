package services

import (
	"errors"
	"fmt"
	"pheasant-api/app/helper"
	"pheasant-api/app/jobs"
	"pheasant-api/app/models"
	"pheasant-api/app/requests"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CheckAuthentication(loginBody requests.LoginBody) (string, error)
	ValidateToken(token string) (*jwt.Token, error, *Claims)

	Register(email string, password string, fullName string) (*models.User, error)

	ForgotPassword(email string) bool

	IsResetPasswordTokenStillValid(token string) bool

	ResetPassword(token string, newPassword string) error
}

type authServiceParams struct {
	secretKey                []byte
	userModel                models.UserModel
	forgotPasswordTokenModel models.ForgotPasswordTokenModel
}

const tokenTtl = 30 // minutes
const bcryptPasswordCost = 14

type Claims struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

func GetAuthService() AuthService {
	return &authServiceParams{
		secretKey:                getJwtKey(),
		userModel:                models.GetUserModel(),
		forgotPasswordTokenModel: models.GetForgotPasswordTokenModel(),
	}
}

// CheckAuthentication will handle the check and return the JWT token on success
func (service *authServiceParams) CheckAuthentication(loginBody requests.LoginBody) (string, error) {
	user, err := service.userModel.GetUserByEmail(loginBody.Email)
	if err != nil {
		return "", err
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password))
	if err != nil {
		return "", err
	}

	// create token
	return createJwtToken(user), nil
}

// ValidateToken check the availability of the Token
func (service *authServiceParams) ValidateToken(token string) (*jwt.Token, error, *Claims) {
	claims := &Claims{}

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})

	return jwtToken, err, claims
}

func (service *authServiceParams) Register(email string, password string, fullName string) (*models.User, error) {
	user, err := service.userModel.GetUserByEmail(email)
	if err != nil || user != nil {
		return nil, errors.New("Email is not available to register")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptPasswordCost)
	if err != nil {
		return nil, errors.New("Internal server error, please try again")
	}

	return service.userModel.Create(&models.User{
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
	})
}

func (service *authServiceParams) ForgotPassword(email string) bool {
	user, err := service.userModel.GetUserByEmail(email)
	if err != nil {
		return false
	}

	// send email to user via Queue Job
	return jobs.InitForgotPasswordJob().Dispatch(*user) == nil
}

func (service *authServiceParams) IsResetPasswordTokenStillValid(token string) bool {
	return service.forgotPasswordTokenModel.IsTokenStillValid(token)
}

func (service *authServiceParams) ResetPassword(token string, newPassword string) error {
	forgotPasswordToken := service.forgotPasswordTokenModel.FindByToken(token)
	if forgotPasswordToken == nil {
		return errors.New("Token is invalid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcryptPasswordCost)
	if err != nil {
		return errors.New("Internal server error, please try again")
	}

	user := forgotPasswordToken.User
	user.Password = string(hashedPassword)

	_, err = service.userModel.Update(&user)
	if err != nil {
		return errors.New("Internal server error, please try again")
	}

	// remove this token
	service.forgotPasswordTokenModel.DeleteByToken(token)

	return nil
}

func getJwtKey() []byte {
	return []byte(helper.GetEnv("APP_KEY", "demo-jwt-key"))
}

func createJwtToken(user *models.User) string {
	tokenExpirationTime := time.Now().Add(tokenTtl * time.Minute)
	claims := &Claims{
		UserId: user.ID,
		Email:  user.Email,
		Name:   user.FullName,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: tokenExpirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(getJwtKey())
	if err != nil {
		return ""
	}

	return tokenString
}
