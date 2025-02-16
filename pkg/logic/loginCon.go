package logic

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}

var DBConn *Handler

func LogicDB(db *gorm.DB) {
	DBConn = &Handler{
		DB: db,
	}
}
