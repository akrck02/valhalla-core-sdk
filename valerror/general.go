package valerror

type General int

// 0 --> 999
const (
	UNEXPECTED_ERROR               = 0
	ACCESS_DENIED                  = 1
	NOT_IMPLEMENTED                = 2
	INVALID_REQUEST                = 3
	DATABASE_ERROR                 = 4
	INVALID_OBJECT_ID              = 5
	UPDATE_ERROR                   = 6
	INSERT_ERROR                   = 7
	DELETE_ERROR                   = 8
	NOT_ENOUGHT_PERMISSIONS        = 9
	INVALID_EMAIL                  = 10
	SHORT_EMAIL                    = 11
	NO_AT_EMAIL                    = 12
	NO_DOT_EMAIL                   = 13
	NO_ALPHANUMERIC_PASSWORD       = 14
	NO_UPPER_LOWER_PASSWORD        = 15
	INVALID_TOKEN                  = 16
	CANNOT_CREATE_VALIDATION_CODE  = 17
	INVALID_VALIDATION_CODE        = 18
	INVALID_PASSWORD               = 19
	SHORT_PASSWORD                 = 20
	NO_SPECIAL_CHARACTERS_PASSWORD = 21
	NO_MAYUS_MINUS_PASSWORD        = 22
	SHORT_NAME                     = 23
	LONG_NAME                      = 24
	SHORT_DESCRIPTION              = 25
	LONG_DESCRIPTION               = 26
)
