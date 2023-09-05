package services

import (
	"errors"
	"goREST/db"
	"goREST/models"
	"goREST/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func SignUp(userInput models.UserRequest) (string, error) {
	// Generate password
	password, err := utils.GeneratePasswordHash([]byte(userInput.Password))
	if err != nil {
		return "", err
	}

	// Create user object
	var user = models.User{
		Email:    userInput.Email,
		Password: string(password),
		ID:       uuid.NewString(),
	}
	// Write to db
	dbErr := db.DBConn.Create(&user).Error
	if dbErr != nil {
		sqlError := dbErr.(*mysql.MySQLError)
		// Duplicate Key
		if sqlError.Number == 1062 {
			return "", sqlError
		}
	}

	// Generate Token
	token, err := utils.GenerateAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(userInput models.UserRequest) (string, error) {
	var user models.User
	// Retrieve user object
	result := db.DBConn.First(&user, "email = ?", userInput.Email)
	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	// Check password
	checkEquivalance := utils.ComparePasswordHash(user.Password, userInput.Password)
	if checkEquivalance != nil {
		return "", errors.New("credentials incorrect")
	}
	// Generate access token
	token, _ := utils.GenerateAccessToken()
	return token, nil
}
