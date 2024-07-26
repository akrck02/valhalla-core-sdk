package projectmodels

import "go.mongodb.org/mongo-driver/bson"

type Project struct {
	CreationDate *int64   `bson:"creationdate,omitempty"`
	LastUpdate   *int64   `bson:"updatedate,omitempty"`
	Id           string   `bson:"_id,omitempty"`
	Name         string   `bson:"name,omitempty"`
	Description  string   `bson:"description,omitempty"`
	Owner        string   `bson:"owner,omitempty"`
	Teams        []string `bson:"teams,omitempty"`
	Wikis        []string `bson:"wikis,omitempty"`
	Notes        []string `bson:"notes,omitempty"`
	Tasks        []string `bson:"tasks,omitempty"`
}

func (p *Project) Default() {
	p.CreationDate = nil
	p.LastUpdate = nil
	p.Id = ""
	p.Name = ""
	p.Description = ""
	p.Owner = ""
	p.Teams = []string{}
	p.Wikis = []string{}
	p.Notes = []string{}
	p.Tasks = []string{}
}

func (p *Project) Clone() *Project {
	return &Project{
		CreationDate: p.CreationDate,
		LastUpdate:   p.LastUpdate,
		Id:           p.Id,
		Name:         p.Name,
		Description:  p.Description,
		Owner:        p.Owner,
		Teams:        p.Teams,
		Wikis:        p.Wikis,
		Notes:        p.Notes,
		Tasks:        p.Tasks,
	}
}

func (p *Project) Bson(hideId bool) bson.M {

	bsonMap := bson.M{}

	if p.Name != "" {
		bsonMap["name"] = p.Name
	}

	if p.Description != "" {
		bsonMap["description"] = p.Description
	}

	if p.Owner != "" {
		bsonMap["owner"] = p.Owner
	}

	if len(p.Teams) > 0 {
		bsonMap["teams"] = p.Teams
	}

	if len(p.Wikis) > 0 {
		bsonMap["wikis"] = p.Wikis
	}

	if len(p.Notes) > 0 {
		bsonMap["notes"] = p.Notes
	}

	if len(p.Tasks) > 0 {
		bsonMap["tasks"] = p.Tasks
	}

	if !hideId && p.Id != "" {
		bsonMap["_id"] = p.Id
	}

	if p.LastUpdate != nil {
		bsonMap["updatedate"] = p.LastUpdate
	}

	if p.CreationDate != nil {
		bsonMap["creationdate"] = p.CreationDate
	}

	return bsonMap
}
