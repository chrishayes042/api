package model

type User struct {
	UserId    int32  `json:"UserId"`
	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

type Users []User
