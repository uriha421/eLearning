package defs

//the definition of data model

//requests
type UserCredential struct {
	Username string      `json:"user_name"`
	Pwd string           `json:"pwd"`
}
