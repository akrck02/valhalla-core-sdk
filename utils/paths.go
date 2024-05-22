package utils

import (
	"os"
)

const DEFAULT_PROFILE_PICTURE_EXTENSION = ".jpg"

func GetProfilePicturePath(username string, path string) string {

	var picpath string = path + username

	if username != "" {
		picpath += DEFAULT_PROFILE_PICTURE_EXTENSION
	}

	return picpath
}

func ExistsDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func Ping() string {
	return "Pong!"
}
