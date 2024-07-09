package usersmodels

type AuthLogin struct {
	Email     string `bson:"email,omitempty"`
	AuthToken string `bson:"auth_token,omitempty"`
}
