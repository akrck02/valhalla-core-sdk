package apimodels

import (
	apierror "github.com/akrck02/valhalla-core-sdk/error"
)

type Error struct {
	Status  int               `json:"status,omitempty"`
	Error   apierror.ApiError `json:"error,omitempty"`
	Message string            `json:"message,omitempty"`
}
