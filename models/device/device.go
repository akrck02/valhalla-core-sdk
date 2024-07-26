package devicemodels

import "go.mongodb.org/mongo-driver/bson"

type Device struct {
	Id           string `bson:"_id,omitempty"`
	CreationDate *int64 `bson:"creationdate,omitempty"`
	LastUpdate   *int64 `bson:"updatedate,omitempty"`
	User         string `bson:"user,omitempty"`
	Address      string `bson:"address,omitempty"`
	UserAgent    string `bson:"useragent,omitempty"`
	Token        string `bson:"token,omitempty"`
}

func (d *Device) Default() {
}

func (d *Device) Clone() *Device {
	return &Device{
		Id:           d.Id,
		CreationDate: d.CreationDate,
		LastUpdate:   d.LastUpdate,
		User:         d.User,
		Address:      d.Address,
		UserAgent:    d.UserAgent,
		Token:        d.Token,
	}
}

func (d *Device) Bson() bson.M {

	bsonMap := bson.M{}

	if d.User != "" {
		bsonMap["user"] = d.User
	}

	if d.Address != "" {
		bsonMap["address"] = d.Address
	}

	if d.UserAgent != "" {
		bsonMap["useragent"] = d.UserAgent
	}

	if d.Token != "" {
		bsonMap["token"] = d.Token
	}

	if d.CreationDate != nil {
		bsonMap["creationdate"] = d.CreationDate
	}

	if d.LastUpdate != nil {
		bsonMap["updatedate"] = d.LastUpdate
	}

	return bsonMap
}
