package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/t"
	"my_pilot/internal/handler"
	"my_pilot/internal/handler/batch"
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

func BatchSaveHotelStaticInfo(context *gin.Context) {

	err := batch.Main()
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

func SearchHotelStaticInfo(context *gin.Context) {

	hotelId, ok := context.GetQuery("hotelId")
	if !ok {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint("hotelId不能为空"),
		})
		return
	}

	queryHotelDetailResponse, err := handler.SearchHotelStaticInfo(t.ParseInt(hotelId))
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), hotelId),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    queryHotelDetailResponse,
	})
	return
}

func QueryHotelRatePlan(context *gin.Context) {

	qr := handler.QueryRatePlanRequest{}

	err := context.ShouldBindJSON(&qr)
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), fmt.Sprintf("%v", qr)),
		})
		return
	}

	queryRatePlanResult, err := handler.QueryHotelRatePlan(qr)
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), fmt.Sprintf("%v", qr)),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    queryRatePlanResult,
	})
	return
}

func QueryOrderPrice(context *gin.Context) {

	qr := handler.QueryOrderPriceRequest{}

	err := context.ShouldBindJSON(&qr)
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), fmt.Sprintf("%v", qr)),
		})
		return
	}

	queryOrderPriceResult, err := handler.QueryOrderPrice(qr)
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), fmt.Sprintf("%v", qr)),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    queryOrderPriceResult,
	})
	return
}

func CreateOrder(context *gin.Context) {

	qr := handler.CreateOrderRequest{}

	err := context.ShouldBindJSON(&qr)
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), fmt.Sprintf("%v", qr)),
		})
		return
	}

	createOrderResult, err := handler.CreateOrder(qr)
	if err != nil {
		context.JSON(200, Response{
			Success: false,
			Msg:     fmt.Sprint(err.Error(), fmt.Sprintf("%v", qr)),
		})
		return
	}

	context.JSON(200, Response{
		Success: true,
		Data:    createOrderResult,
	})
	return
}
