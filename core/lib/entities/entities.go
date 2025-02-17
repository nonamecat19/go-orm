package entities

type Model struct {
	ID int64 `db:"id" json:"id"`
	//CreatedAt time.Time `db:"created_at" json:"created_at"`
	//UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type IEntity interface {
	Info() string
}
