package apimodels

import (
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
)

type Request struct {
	Authorization string            `json:"authorization"`
	IP            string            `json:"ip"`
	UserAgent     string            `json:"userAgent"`
	User          *usersmodels.User `json:"user"`
	Params        interface{}       `json:"params"`
	Body          interface{}       `json:"body"`
	Headers       interface{}       `json:"headers"`
	Files         interface{}       `json:"files"`
	Queries       interface{}       `json:"queries"`
}
