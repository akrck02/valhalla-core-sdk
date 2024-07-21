package utils

import (
	"net/http"
	"strings"

	apierror "github.com/akrck02/valhalla-core-sdk/error"

	"github.com/akrck02/valhalla-core-sdk/lang"
	apimodels "github.com/akrck02/valhalla-core-sdk/models/api"
)

const MINIMUM_CHARACTERS_FOR_PASSWORD = 16
const MINIMUM_CHARACTERS_FOR_EMAIL = 5

func ValidateEmail(email string) *apimodels.Error {

	if len(email) < MINIMUM_CHARACTERS_FOR_EMAIL {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.ShortEmail,
			Message: "Email must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_EMAIL) + " characters",
		}
	}

	if !strings.Contains(email, "@") {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.NoAtEmail,
			Message: "Email must have one @",
		}
	}

	if !strings.Contains(email, ".") {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.NoDotEmail,
			Message: "Email must have at least one .",
		}
	}

	return nil
}

func ValidatePassword(password string) *apimodels.Error {

	if len(password) < MINIMUM_CHARACTERS_FOR_PASSWORD {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.ShortPassword,
			Message: "Password must have at least " + lang.Int2String(MINIMUM_CHARACTERS_FOR_PASSWORD) + " characters",
		}
	}

	if !ContainsSpecialCharacters(password) {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.NoSpecialCharactersPassword,
			Message: "Password must have at least one special character",
		}
	}

	if IsLowerCase(password) {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.NoMayusMinusPassword,
			Message: "Password must have at least one uppercase character",
		}
	}

	if IsUpperCase(password) {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.NoMayusMinusPassword,
			Message: "Password must have at least one lowercase character",
		}
	}

	if !ContainsNumbers(password) {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.NoAlphanumericPassword,
			Message: "Password must have at least one number",
		}
	}

	return nil
}

func ValidateName(name string) *apimodels.Error {

	if len(name) < 2 {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.ShortName,
			Message: "Name must have at least 2 characters",
		}
	}

	if len(name) > 50 {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.LongName,
			Message: "Name must have at most 50 characters",
		}
	}

	return nil
}

func ValidateDescription(description string) *apimodels.Error {

	if len(description) < 2 {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.ShortDescription,
			Message: "Description must have at least 2 characters",
		}
	}

	if len(description) > 500 {
		return &apimodels.Error{
			Status:  http.StatusBadRequest,
			Error:   apierror.LongDescription,
			Message: "Description must have at most 500 characters",
		}
	}

	return nil
}
