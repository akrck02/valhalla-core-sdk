package apimodels

type Message struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
