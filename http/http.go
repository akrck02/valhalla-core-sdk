package http

type HttpMethod int

const (
	HTTP_METHOD_GET     = 0
	HTTP_METHOD_POST    = 1
	HTTP_METHOD_PUT     = 2
	HTTP_METHOD_DELETE  = 3
	HTTP_METHOD_PATCH   = 4
	HTTP_METHOD_HEAD    = 5
	HTTP_METHOD_OPTIONS = 6
	HTTP_METHOD_TRACE   = 7
	HTTP_METHOD_CONNECT = 8
)

type HttpStatusCode int

const (
	HTTP_STATUS_OK                         = 200
	HTTP_STATUS_CREATED                    = 201
	HTTP_STATUS_ACCEPTED                   = 202
	HTTP_STATUS_NO_CONTENT                 = 204
	HTTP_STATUS_NO_CHANGE                  = 304
	HTTP_STATUS_BAD_REQUEST                = 400
	HTTP_STATUS_UNAUTHORIZED               = 401
	HTTP_STATUS_FORBIDDEN                  = 403
	HTTP_STATUS_NOT_FOUND                  = 404
	HTTP_STATUS_METHOD_NOT_ALLOWED         = 405
	HTTP_STATUS_NOT_ACCEPTABLE             = 406
	HTTP_STATUS_CONFLICT                   = 409
	HTTP_STATUS_INTERNAL_SERVER_ERROR      = 500
	HTTP_STATUS_NOT_IMPLEMENTED            = 501
	HTTP_STATUS_BAD_GATEWAY                = 502
	HTTP_STATUS_SERVICE_UNAVAILABLE        = 503
	HTTP_STATUS_GATEWAY_TIMEOUT            = 504
	HTTP_STATUS_HTTP_VERSION_NOT_SUPPORTED = 505
)
