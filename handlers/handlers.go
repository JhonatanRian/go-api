package handlers

import (
	"github.com/JhonatanRian/go-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handlers")
	db = config.GetSQLite()
}
