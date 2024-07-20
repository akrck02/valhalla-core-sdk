package apimodels

import "github.com/akrck02/valhalla-core-sdk/models/database"

type APIContext struct {
	Launcher    Launcher
	Trazability Trazability
	Database    database.Database
	Request     Request
	Response    Response
}
