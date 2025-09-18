package repository

import (
	"context"

	"gorm.io/gorm"

	"belajar-go/model"
)

type Todo interface {
	GetAll(ctx context.Context) ([]model.Todo, error)
	GetById(ctx context.Context, id int) (*model.Todo, error)
	CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}

// abaikan
type Connection struct {
	Database *gorm.DB
}

// struct dari gorm kita conversi ke struct kita

func NewTodo(Database *gorm.DB) Todo {
	return &Connection{
		Database: Database,
	}
}

// abaikan

func (c *Connection) GetAll(ctx context.Context) ([]model.Todo, error) {
	var todo []model.Todo
	if err := c.Database.WithContext(ctx).Find(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (c *Connection) GetById(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo

	if err := c.Database.WithContext(ctx).First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil

}

func (c *Connection) CreateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {

	if err := c.Database.WithContext(ctx).Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (c *Connection) UpdateTodo(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	if err := c.Database.WithContext(ctx).Save(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
func (c *Connection) DeleteTodo(ctx context.Context, id int) error {

	return c.Database.WithContext(ctx).Delete(model.Todo{}, id).Error
}
