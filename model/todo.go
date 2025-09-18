package model

import "time"

type Todo struct {
	ID          int       `gorm:"primaryKey;column:id"`
	Title       string    `gorm:"column:title;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:text"`
	IsCompleted bool      `gorm:"column:is_completed;default:false"`
	DueDate     time.Time `gorm:"column:due_date"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Todo) TableName() string {
	return "todos"
}
