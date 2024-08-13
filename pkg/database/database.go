package database

import (
	"github.com/aebalz/go-gin-gone/configs"
	"gorm.io/gorm"
)

type Database interface {
	Connect(config configs.Config, isDebug bool) error
	Close() error
	GetDB() *gorm.DB
}
