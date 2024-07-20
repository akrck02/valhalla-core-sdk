package apimodels

type Trazability struct {
	Endpoint  Endpoint `json:"endpoint"`
	Timestamp string   `json:"timestamp"`
}
