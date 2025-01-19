package database

import (
	"context"

	"gorm.io/gorm"
)

var DBConn IDatabase

type IDatabase interface {
	Ping() error
	Close() error
	WithContext(ctx context.Context) *gorm.DB
	Begin() IDatabase
	Commit() error
	Rollback() error
	Client() IDatabase
}

type ITransactor interface {
	Begin() IDatabase
	Commit() error
	Rollback() error
}
