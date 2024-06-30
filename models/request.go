package models

import (
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
)

type Request struct {
	Authorization string            `json:"authorization"`
	IP            string            `json:"ip"`
	UserAgent     string            `json:"userAgent"`
	User          *usersmodels.User `json:"user"`
}
