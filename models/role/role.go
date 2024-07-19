package rolemodels

type Role struct {
	CreationDate *int64 `bson:"creationdate,omitempty"`
	LastUpdate   *int64 `bson:"updatedate,omitempty"`
	Name         string `bson:"name,omitempty"`
	Description  string `bson:"description,omitempty"`
	Team         string `bson:"team,omitempty"`
	Permissions  []int  `bson:"permissions,omitempty"`
}
