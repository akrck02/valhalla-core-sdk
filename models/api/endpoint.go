package apimodels

type Endpoint struct {
	Path     string           `json:"path"`
	Method   int              `json:"method"`
	Listener EndpointListener `json:"-"`
	Checks   EndpointCheck    `json:"-"`
	Secured  bool             `json:"secured"`
	Database bool             `json:"-"`
}

type EndpointCheck func(context *ApiContext) *Error
type EndpointListener func(context *ApiContext) (*Response, *Error)

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
