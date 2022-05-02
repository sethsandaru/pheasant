package services

import (
	"fmt"
	"pheasant-api/app/helper"
	"pheasant-api/app/models"
	"pheasant-api/app/requests"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CheckAuthentication(loginBody requests.LoginBody) (string, error)
	ValidateToken(token string) (*jwt.Token, error)

	Register(email string, password string, fullName string) (*models.User, error)
}

type authServiceParams struct {
	secretKey []byte
	userModel models.UserModel
}

const tokenTtl = 30 // minutes
const bcryptPasswordCost = 14

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func GetAuthService() AuthService {
	return &authServiceParams{
		secretKey: getJwtKey(),
		userModel: models.GetUserModel(),
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
func (service *authServiceParams) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})
}

func (service *authServiceParams) Register(email string, password string, fullName string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptPasswordCost)
	if err != nil {
		return nil, err
	}

	return service.userModel.CreateUser(&models.User{
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
	})
}

func getJwtKey() []byte {
	return []byte(helper.GetEnv("APP_KEY", "demo-jwt-key"))
}

func createJwtToken(user *models.User) string {
	tokenExpirationTime := time.Now().Add(tokenTtl * time.Minute)
	claims := &Claims{
		Email: user.Email,
		Name:  user.FullName,
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
