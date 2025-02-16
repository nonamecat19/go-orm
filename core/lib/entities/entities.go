package entities

type Model struct {
	ID int64 `db:"id" json:"id"`
}

type IEntity interface {
	Info() string
}
