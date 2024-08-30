package entities

type Model struct {
	ID int64 `db:"id" json:"id"`
}

type IEntity interface {
	TableName() string
}

type User struct {
	Model
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func (user *User) TableName() string {
	return "users"
}
