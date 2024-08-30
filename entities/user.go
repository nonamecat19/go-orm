package entities

type User struct {
	Model
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func (user *User) TableName() string {
	return "users"
}
