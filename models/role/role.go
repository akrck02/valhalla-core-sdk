package rolemodels

type Role struct {
	Name         string `bson:"name,omitempty"`
	Description  string `bson:"description,omitempty"`
	Team         string `bson:"team,omitempty"`
	Permissions  []int  `bson:"permissions,omitempty"`
	CreationDate string `bson:"creationdate,omitempty"`
	LastUpdate   string `bson:"updatedate,omitempty"`
}
