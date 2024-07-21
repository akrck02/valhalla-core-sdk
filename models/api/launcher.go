package apimodels

type Launcher struct {
	Id           string       `json:"id"`
	LauncherType LauncherType `json:"type"`
}
