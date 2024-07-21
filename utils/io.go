package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func SaveFile(filepath string, content []byte) error {

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()
	n, err := file.Write(content)

	if err != nil {
		return err
	}

	if n != len(content) {
		return errors.New("error writing file")
	}

	return nil
}

func ReadFile(filepath string) ([]byte, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func ParseJson[T interface{}](body io.Reader, object *T) error {

	err := json.NewDecoder(body).Decode(object)
	if err != nil {
		return err
	}

	return nil
}
