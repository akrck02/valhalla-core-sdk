package apimodels

import (
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
)

type Endpoint struct {
	Path     string           `json:"path"`
	Method   int              `json:"method"`
	Listener EndpointListener `json:"-"`
	Checks   EndpointCheck    `json:"-"`
	Secured  bool             `json:"secured"`
	Database bool             `json:"-"`
}

type EndpointCheck func(context *systemmodels.ValhallaContext) *systemmodels.Error
type EndpointListener func(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error)

func EndpointFrom(path string, method int, listener EndpointListener, checks EndpointCheck, secured bool, database bool) Endpoint {
	return Endpoint{
		Path:     path,
		Method:   method,
		Listener: listener,
		Checks:   checks,
		Secured:  secured,
		Database: database,
	}
}
