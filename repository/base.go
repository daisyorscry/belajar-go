package repository

import (
	"context"

	"gorm.io/gorm"
)

type Todo interface {
	GetAll(ctx context.Context)
}

type Connection struct {
	Database *gorm.DB
}
