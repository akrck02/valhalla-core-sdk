package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	apierror "github.com/akrck02/valhalla-core-sdk/error"

	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
	devicemodels "github.com/akrck02/valhalla-core-sdk/models/device"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"

	"github.com/golang-jwt/jwt/v5"
)

const OTP_CHARS = "1234567890"

func GenerateAuthToken(user *usersmodels.User, device *devicemodels.Device, secret string) (string, *apimodels.Error) {

	now := GetCurrentMillis()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"device":    device.UserAgent + "-" + device.Address,
		"username":  user.Username,
		"email":     user.Email,
		"timestamp": now,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", &apimodels.Error{
			Status:  http.StatusInternalServerError,
			Error:   apierror.CannotGenerateAuthToken,
			Message: "Error generating auth token",
		}
	}

	return tokenString, nil
}

func DecryptToken(token string, secret string) (*jwt.Token, *apimodels.Error) {

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, &apimodels.Error{
			Status:  http.StatusUnauthorized,
			Error:   apierror.InvalidToken,
			Message: "Invalid token",
		}
	}

	return parsed, nil

}

func EncryptSha256(text string) string {
	plainText := []byte(text)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func GenerateValidationCode(text string) (string, *apimodels.Error) {

	// Generate a random string
	randomString, err := GenerateOTP(20)

	if err != nil {
		return "", err
	}

	return randomString + EncryptSha256(text), nil
}

func GenerateOTP(length int) (string, *apimodels.Error) {

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", &apimodels.Error{
			Status:  http.StatusInternalServerError,
			Error:   apierror.CannotCreateValidationCode,
			Message: "Error generating OTP",
		}

	}

	otpCharsLength := len(OTP_CHARS)
	for i := 0; i < length; i++ {
		buffer[i] = OTP_CHARS[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
