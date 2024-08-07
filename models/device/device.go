package devicemodels

type Device struct {
	CreationDate *int64 `bson:"creationdate,omitempty"`
	LastUpdate   *int64 `bson:"updatedate,omitempty"`
	User         string `bson:"user,omitempty"`
	Address      string `bson:"address,omitempty"`
	UserAgent    string `bson:"useragent,omitempty"`
	Token        string `bson:"token,omitempty"`
}
