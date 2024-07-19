package projectmodels

type Project struct {
	CreationDate *int64   `bson:"creationdate,omitempty"`
	LastUpdate   *int64   `bson:"updatedate,omitempty"`
	ID           string   `bson:"_id,omitempty"`
	Name         string   `bson:"name,omitempty"`
	Description  string   `bson:"description,omitempty"`
	Owner        string   `bson:"owner,omitempty"`
	Teams        []string `bson:"teams,omitempty"`
	Wikis        []string `bson:"wikis,omitempty"`
	Notes        []string `bson:"notes,omitempty"`
	Tasks        []string `bson:"tasks,omitempty"`
}
