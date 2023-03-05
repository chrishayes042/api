package model

type User struct {
	UserId    int32  `json:"UserId" db:"id"`
	UserName  string `json:"UserName" db:"user_name"`
	FirstName string `json:"FirstName" db:"first_name"`
	LastName  string `json:"LastName" db:"last_name"`
	Email     string `json:"Email" db:"email"`
	CreatedOn string `db:"created_on"`
}

type Users []User
