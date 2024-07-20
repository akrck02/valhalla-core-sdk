package systemmodels

type ValhallaContext struct {
	Launcher    Launcher
	Trazability Trazability
	Database    Database
	Request     Request
	Response    Response
}
