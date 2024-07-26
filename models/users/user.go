package usersmodels

import (
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	CreationDate   *int64 `bson:"creationdate,omitempty"`
	LastUpdate     *int64 `bson:"updatedate,omitempty"`
	Email          string `bson:"email,omitempty"`
	Password       string `bson:"password,omitempty"`
	Username       string `bson:"username,omitempty"`
	Validated      bool   `bson:"validated"`
	ValidationCode string `bson:"validation_code,omitempty"`
	ProfilePic     string `bson:"profile_pic,omitempty"`
	Id             string `bson:"_id,omitempty"`
}

func (u *User) Default() {
	u.CreationDate = nil
	u.LastUpdate = nil
	u.Email = ""
	u.Password = ""
	u.Username = ""
	u.Validated = false
	u.ValidationCode = ""
	u.ProfilePic = ""
	u.Id = ""
}

func (u *User) Clone() *User {
	return &User{
		CreationDate:   u.CreationDate,
		LastUpdate:     u.LastUpdate,
		Email:          u.Email,
		Password:       u.Password,
		Username:       u.Username,
		Validated:      u.Validated,
		ValidationCode: u.ValidationCode,
		ProfilePic:     u.ProfilePic,
		Id:             u.Id,
	}
}

func (u *User) Bson(hideId bool) bson.M {

	bsonMap := bson.M{}

	if u.Email != "" {
		bsonMap["email"] = u.Email
	}

	if u.Password != "" {
		bsonMap["password"] = u.Password
	}

	if u.Username != "" {
		bsonMap["username"] = u.Username
	}

	if u.Validated {
		bsonMap["validated"] = u.Validated
	}

	if u.ValidationCode != "" {
		bsonMap["validation_code"] = u.ValidationCode
	}

	if u.ProfilePic != "" {
		bsonMap["profile_pic"] = u.ProfilePic
	}

	if !hideId && u.Id != "" {
		bsonMap["_id"] = u.Id
	}

	if u.LastUpdate != nil {
		bsonMap["updatedate"] = u.LastUpdate
	}

	if u.CreationDate != nil {
		bsonMap["creationdate"] = u.CreationDate
	}

	return bsonMap
}
