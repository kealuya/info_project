package controller

import (
	"github.com/gin-gonic/gin"
	"my_pilot/internal/handler"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
}

func SaveLocationInfo(context *gin.Context) {

	err := handler.SaveLocationInfo()
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    nil,
	})
	return
}

func SaveHotelInfo(context *gin.Context) {

	err := handler.SaveHotelInfo()
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    nil,
	})
	return
}

func SaveHotelStaticInfo(context *gin.Context) {

	err := handler.SaveHotelDetailInfo()
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    nil,
	})
	return
}
