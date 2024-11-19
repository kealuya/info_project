package batch

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"my_pilot/internal/repository"
)

func getHotelInfoFromDb() chan []repository.HotelInfo {

	taskChan := make(chan []repository.HotelInfo, 10)

	const saveHotelDetailInfoPageSize = 200 // 每次从数据库获取的酒店数量
	start := 0
	totalProcessed := 0
	go func() {
		defer close(taskChan)
		for {
			// 分页获取酒店信息
			var hotels []repository.HotelInfo
			func() {
				defer func() {
					if err := recover(); err != nil {
						logs.Error(fmt.Sprintf("getHotelInfoFromDb error ::%v,start:%v,limit:%v", err, start, saveHotelDetailInfoPageSize))
					}
				}()
				hotels = repository.GetHotelsAll(start, saveHotelDetailInfoPageSize)
			}()

			// 如果没有更多数据，退出循环
			if len(hotels) == 0 {
				break
			}

			// 将酒店数据发送到任务通道
			taskChan <- hotels

			totalProcessed += len(hotels)
			logs.Info(fmt.Sprintf("已发送 %d 条酒店记录到处理队列\n", totalProcessed))

			// 更新起始位置
			start += len(hotels)
		}
	}()

	return taskChan

}
