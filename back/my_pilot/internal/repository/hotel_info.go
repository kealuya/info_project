package repository

import (
	"my_pilot/common"
	"time"
)

type HotelInfo struct {
	Id                  int64     `xorm:"pk autoincr 'id'"`
	HotelId             int       `xorm:"'hotel_id' not null"`
	CountryId           int       `xorm:"'country_id' not null"`
	StateId             int       `xorm:"'state_id'"`
	CityId              int       `xorm:"'city_id'"`
	Star                int       `xorm:"'star'"`
	HotelNameCn         string    `xorm:"'hotel_name_cn' varchar(200)"`
	HotelNameEn         string    `xorm:"'hotel_name_en' varchar(200)"`
	AddressCn           string    `xorm:"'address_cn' varchar(500)"`
	AddressEn           string    `xorm:"'address_en' varchar(500)"`
	Phone               string    `xorm:"'phone' varchar(50)"`
	Longitude           string    `xorm:"'longitude' varchar(20)"`
	Latitude            string    `xorm:"'latitude' varchar(20)"`
	InstantConfirmation int       `xorm:"'instant_confirmation'"`
	SellType            int       `xorm:"'sell_type'"`
	UpdateTime          time.Time `xorm:"'update_time'"`
	Status              int       `xorm:"'status'"`
	Created             time.Time `xorm:"created"`
	Updated             time.Time `xorm:"updated"`
}

func InsertHotels(hotels []HotelInfo) {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	for _, hotel := range hotels {
		_, err := session.Insert(&hotel)
		if err != nil {
			session.Rollback()
			common.ErrorHandler(err)
			return
		}
	}

	errSessionCommit := session.Commit()
	common.ErrorHandler(errSessionCommit)
	return
}

// GetHotelsAll 分页获取酒店信息
func GetHotelsAll(start, limit int) []HotelInfo {
	var hotels []HotelInfo
	err := dbEngine.Limit(limit, start).Find(&hotels)
	common.ErrorHandler(err)
	return hotels
}

// CountHotels 计算全部酒店数量
func CountHotels() int64 {
	var hotels []HotelInfo
	count, err := dbEngine.Count(&hotels)
	common.ErrorHandler(err)
	return count
}
