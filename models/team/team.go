package teammodels

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Team struct {
	CreationDate *int64   `bson:"creationdate,omitempty"`
	LastUpdate   *int64   `bson:"updatedate,omitempty"`
	Name         string   `bson:"name,omitempty"`
	Description  string   `bson:"description,omitempty"`
	ProfilePic   string   `bson:"profilepic,omitempty"`
	Projects     []string `bson:"projects,omitempty"`
	Owner        string   `bson:"owner,omitempty"`
	Members      []string `bson:"members,omitempty"`
	ID           string   `bson:"_id,omitempty"`
}

func (t *Team) Clone() *Team {
	return &Team{
		CreationDate: t.CreationDate,
		LastUpdate:   t.LastUpdate,
		Name:         t.Name,
		Description:  t.Description,
		ProfilePic:   t.Description,
		Projects:     t.Projects,
		Owner:        t.Owner,
		Members:      t.Members,
		ID:           t.ID,
	}
}

func (t *Team) PurgedBson(hideID bool) bson.M {

	purgedBson := bson.M{}

	if t.Name != "" {
		purgedBson["name"] = t.Name
	}

	if t.Description != "" {
		purgedBson["description"] = t.Description
	}

	if t.ProfilePic != "" {
		purgedBson["profilepic"] = t.ProfilePic
	}

	if t.Owner != "" {
		purgedBson["owner"] = t.Owner
	}

	if !hideID && t.ID != "" {
		purgedBson["_id"] = t.ID
	}

	return purgedBson
}
