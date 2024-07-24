package apierror

type ApiError int

// 0 --> 999 | GENERAL ERRORS
const (
	UnexpectedError ApiError = iota
	AccessDenied
	NotImplemented
	InvalidRequest
	DatabaseError
	InvalidObjectId
	NotEnoughtPermissions
	InvalidEmail
	ShortEmail
	NoAtEmail
	NoDotEmail
	NoAlphanumericPassword
	NoMayusMinusPassword
	InvalidToken
	CannotCreateValidationCode
	InvalidValidationCode
	InvalidPassword
	ShortPassword
	NoSpecialCharactersPassword
	ShortName
	LongName
	ShortDescription
	LongDescription
	UpdateParametersEqual
	UnsupportedOperation
)

// 1000 --> 1999 | USER ERRORS
const (
	UserAlreadyExists ApiError = iota + 1000
	UserNotupdated
	UserNotFound
	UserNotDeleted
	EmptyUsername
	EmptyUserPassword
	EmptyUserEmail
	UserAlreadyValidated
	UserNotAuthorized
)

// 2000 --> 2999 | PROJECT ERRORS
const (
	EmptyProjectName ApiError = iota + 2000
	EmptyProjectDescription
	EmptyProjectOwner
	ProjectAlreadyExists
	ProjectNotFound
)

// 3000 --> 3999 | TEAM ERRORS
const (
	TeamAlreadyExists ApiError = iota + 3000
	EmptyTeamName
	EmptyOwner
	OwnerNotFound
	TeamNotFound
	EmptyTeamDescription
	EmptyMember
	TeamEmpty
	UserIsAlreadyMember
	TeamSearchError
	UserIsNotMember
)

// 4000 --> 4999 | DEVICE ERRORS
const (
	DeviceNotFound ApiError = iota + 5000
	CannotGenerateAuthToken
)
