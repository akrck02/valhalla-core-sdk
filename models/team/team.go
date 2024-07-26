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
	Id           string   `bson:"_id,omitempty"`
}

func (t *Team) Default() {
	t.CreationDate = nil
	t.LastUpdate = nil
	t.Name = ""
	t.Description = ""
	t.ProfilePic = ""
	t.Projects = []string{}
	t.Owner = ""
	t.Members = []string{}
	t.Id = ""
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
		Id:           t.Id,
	}
}

func (t *Team) Bson(hideID bool) bson.M {

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

	if !hideID && t.Id != "" {
		purgedBson["_id"] = t.Id
	}

	if len(t.Projects) > 0 {
		purgedBson["projects"] = t.Projects
	}

	if len(t.Members) > 0 {
		purgedBson["members"] = t.Members
	}

	return purgedBson
}
