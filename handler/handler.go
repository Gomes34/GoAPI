package handler

import (
	"github.com/Gomes34/GoAPI/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitalizeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
