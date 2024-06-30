package systemmodels

import "context"

type Database struct {
	Name    string `json:"name"`
	Context context.Context
}
