package apimodels

import usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"

type Trazability struct {
	Endpoint  Endpoint          `json:"endpoint"`
	Timestamp *int64            `json:"timestamp"`
	Launcher  Launcher          `json:"launcher"`
	User      *usersmodels.User `json:"user"`
}
