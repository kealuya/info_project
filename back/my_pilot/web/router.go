package web

import (
	"github.com/gin-gonic/gin"
	"my_pilot/web/controller"
	"net/http/pprof"
)

// SetupRouter 初始化路由
func SetupRouter(r *gin.Engine) *gin.Engine {

	// 可以在这里添加全局中间件
	// r.Use(middleware.Logger())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	// API 版本分组
	v1 := r.Group("/api/v1")
	{
		system := v1.Group("/system_init")
		{
			system.GET("/save_location_info", controller.SaveLocationInfo)
			system.GET("/save_hotel_info", controller.SaveHotelInfo)
			system.GET("/save_hotel_static_info", controller.SaveHotelStaticInfo)
			system.GET("/search_hotel_static_info", controller.SearchHotelStaticInfo)
			system.POST("/query_hotel_rate_plan", controller.QueryHotelRatePlan)
			system.POST("/query_hotel_order_price", controller.QueryOrderPrice)
			system.POST("/create_hotel_order", controller.CreateOrder)
		}
	}

	// 注册debug路由组
	debug := r.Group("/debug/pprof")
	{
		debug.GET("/", gin.WrapF(pprof.Index))
		debug.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		debug.GET("/profile", gin.WrapF(pprof.Profile))
		debug.POST("/symbol", gin.WrapF(pprof.Symbol))
		debug.GET("/symbol", gin.WrapF(pprof.Symbol))
		debug.GET("/trace", gin.WrapF(pprof.Trace))
		debug.GET("/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
		debug.GET("/block", gin.WrapF(pprof.Handler("block").ServeHTTP))
		debug.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
		debug.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
		debug.GET("/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
		debug.GET("/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
	}

	return r
}
