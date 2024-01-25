package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creaing opening: %v", err.Error())
		return
	}
}
