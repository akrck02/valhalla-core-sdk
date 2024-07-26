package rolemodels

type Role struct {
	Id           string `bson:"_id,omitempty"`
	CreationDate *int64 `bson:"creationdate,omitempty"`
	LastUpdate   *int64 `bson:"updatedate,omitempty"`
	Name         string `bson:"name,omitempty"`
	Description  string `bson:"description,omitempty"`
	Team         string `bson:"team,omitempty"`
	Permissions  []int  `bson:"permissions,omitempty"`
}

func (r *Role) Default() {

	if r.Permissions == nil {
		r.Permissions = []int{}
	}
}

func (r *Role) Clone() *Role {
	return &Role{
		Id:           r.Id,
		CreationDate: r.CreationDate,
		LastUpdate:   r.LastUpdate,
		Name:         r.Name,
		Description:  r.Description,
		Team:         r.Team,
		Permissions:  r.Permissions,
	}
}

func (r *Role) Bson(hideId bool) map[string]interface{} {

	purgedBson := map[string]interface{}{}

	if r.Name != "" {
		purgedBson["name"] = r.Name
	}

	if r.Description != "" {
		purgedBson["description"] = r.Description
	}

	if r.Team != "" {
		purgedBson["team"] = r.Team
	}

	if len(r.Permissions) > 0 {
		purgedBson["permissions"] = r.Permissions
	}

	return purgedBson
}
