package apimodels

type Request struct {
	Authorization string      `json:"authorization"`
	IP            string      `json:"ip"`
	UserAgent     string      `json:"userAgent"`
	Params        interface{} `json:"params"`
	Body          interface{} `json:"body"`
	Headers       interface{} `json:"headers"`
	Files         interface{} `json:"files"`
	Queries       interface{} `json:"queries"`
}
