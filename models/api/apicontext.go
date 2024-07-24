package apimodels

import "github.com/akrck02/valhalla-core-sdk/models/database"

type ApiContext struct {
	Trazability Trazability
	Database    database.Database
	Request     Request
	Response    interface{}
}
