package entities

type Task struct {
	Model
	Title       string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func (user *Task) TableName() string {
	return "tasks"
}
