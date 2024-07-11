package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"github.com/akrck02/valhalla-core-sdk/http"
	devicemodels "github.com/akrck02/valhalla-core-sdk/models/device"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/golang-jwt/jwt/v5"
)

const OTP_CHARS = "1234567890"

// Generate a new auth token
//
// [param] user | models.User | The user
// [param] device | models.Device | The device
//
// [return] string | The token --> error if something went wrong
func GenerateAuthToken(user *usersmodels.User, device *devicemodels.Device, secret string) (string, *systemmodels.Error) {

	now := getCurrentMillis()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"device":    device.UserAgent + "-" + device.Address,
		"username":  user.Username,
		"email":     user.Email,
		"timestamp": now,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", &systemmodels.Error{
			Status:  http.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   valerror.CANNOT_GENERATE_AUTH_TOKEN,
			Message: "Error generating auth token",
		}
	}

	return tokenString, nil
}

// Decrypt a token
//
// [param] token | string | The token
//
// [return] *jwt.Token | The token --> error if something went wrong
func DecryptToken(token string, secret string) (*jwt.Token, *systemmodels.Error) {

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_UNAUTHORIZED,
			Error:   valerror.INVALID_TOKEN,
			Message: "Invalid token",
		}
	}

	return parsed, nil

}

// Encrypt a string using sha256
//
// [param] text | string | The text to encrypt
//
// [return] string | The encrypted text
func EncryptSha256(text string) string {
	plainText := []byte(text)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

// Generate a validation code
//
// [param] text | string | The text to encrypt
//
// [return] string | The encrypted text
func GenerateValidationCode(text string) (string, *systemmodels.Error) {

	// Generate a random string
	randomString, err := GenerateOTP(20)

	if err != nil {
		return "", err
	}

	return randomString + EncryptSha256(text), nil
}

// Generate a random string
//
// [param] length | int | The length of the string
//
// [return] string | The random string --> error if something went wrong
func GenerateOTP(length int) (string, *systemmodels.Error) {

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", &systemmodels.Error{
			Status:  http.HTTP_STATUS_INTERNAL_SERVER_ERROR,
			Error:   valerror.CANNOT_CREATE_VALIDATION_CODE,
			Message: "Error generating OTP",
		}

	}

	otpCharsLength := len(OTP_CHARS)
	for i := 0; i < length; i++ {
		buffer[i] = OTP_CHARS[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
