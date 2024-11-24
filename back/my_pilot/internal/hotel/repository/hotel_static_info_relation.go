package repository

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"my_pilot/common"
	"time"
)

type RoomType struct {
	Id              int64     `xorm:"pk autoincr 'id'"`
	HotelId         int       `xorm:"'hotel_id' not null"`
	RoomTypeId      string    `xorm:"'room_type_id' not null"`
	RoomTypeCn      string    `xorm:"'room_type_cn' varchar(400) not null"`
	RoomTypeEn      string    `xorm:"'room_type_en' varchar(200)"`
	BasisRoomId     int       `xorm:"'basis_room_id' default(-1)"`
	BasisRoomCn     string    `xorm:"'basis_room_cn' varchar(100)"`
	BedType         string    `xorm:"'bed_type' varchar(50)"`
	Maximize        int       `xorm:"'maximize' default(2)"`
	Acreage         string    `xorm:"'acreage' varchar(50)"`
	BedWidth        string    `xorm:"'bed_width' varchar(50)"`
	FloorDistribute string    `xorm:"'floor_distribute' varchar(200)"`
	Facilities      string    `xorm:"'facilities' text"`
	ExtraBedState   int       `xorm:"'extra_bed_state'"`
	BedCount        int       `xorm:"'bed_count' default(0)"`
	Status          int       `xorm:"'status' default(1)"`
	Created         time.Time `xorm:"created"`
	Updated         time.Time `xorm:"updated"`
}

// RateType 价格类型信息
type RateType struct {
	Id         int64     `xorm:"pk autoincr 'id'"`
	HotelId    int       `xorm:"'hotel_id' not null"`
	RateTypeId int       `xorm:"'rate_type_id' not null"`
	RateTypeCn string    `xorm:"'rate_type_cn' varchar(100) not null"`
	RateTypeEn string    `xorm:"'rate_type_en' varchar(100)"`
	Status     int       `xorm:"'status' default(1)"`
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
}

// Image 图片信息
type Image struct {
	Id          int64     `xorm:"pk autoincr 'id'"`
	HotelId     int       `xorm:"'hotel_id' not null"`
	ImageId     int       `xorm:"'image_id' not null"`
	Type        string    `xorm:"'type' varchar(50) not null"`
	RoomTypeIds string    `xorm:"'room_type_ids' varchar(500)"`
	ThumbUrl    string    `xorm:"'thumb_url' varchar(500) not null"`
	ImageUrl    string    `xorm:"'image_url' varchar(500) not null"`
	ImageLogo   int       `xorm:"'image_logo' default(0)"`
	ImageSize   int       `xorm:"'image_size' default(0)"`
	Status      int       `xorm:"'status' default(1)"`
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
}

func InsertRoomTypes(roomTypes []RoomType) error {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	if errSessionBegin != nil {
		return fmt.Errorf("roomType insert failed.%w", errSessionBegin)
	}
	for _, roomType := range roomTypes {
		_, err := session.Insert(&roomType)
		if err != nil {
			//session.Rollback()
			//common.ErrorHandler(err)
			logs.Error(fmt.Errorf("roomType insert failed.%w - roomType:%+v", err, roomType))
			return err
		}
	}

	errSessionCommit := session.Commit()
	if errSessionCommit != nil {
		return fmt.Errorf("roomType insert failed.%w", errSessionCommit)
	}
	return nil
}

// InsertRateTypes 批量插入价格类型
func InsertRateTypes(rateTypes []RateType) error {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	for _, rateType := range rateTypes {
		_, err := session.Insert(&rateType)
		if err != nil {
			//session.Rollback()
			logs.Error(fmt.Errorf("rateType insert failed.%w - rateTypes:%+v", err, rateTypes))
			continue
		}
	}

	errSessionCommit := session.Commit()
	return fmt.Errorf("rateTypes insert failed.%w", errSessionCommit)

}

// InsertImages 批量插入图片信息
func InsertImages(images []Image) error {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	for _, image := range images {
		_, err := session.Insert(&image)
		if err != nil {
			//session.Rollback()
			logs.Error(fmt.Errorf("image insert failed.%w - rateTypes:%+v", err, image))
			continue
		}
	}

	errSessionCommit := session.Commit()
	return fmt.Errorf("images insert failed.%w", errSessionCommit)
}

// UpdateRoomTypesStatus 批量更新房型状态
func UpdateRoomTypesStatus(hotelId int64, roomTypeIds []int, status int) {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	_, err := session.Table("room_types").
		Where("hotel_id = ? AND room_type_id IN (?)", hotelId, roomTypeIds).
		Update(map[string]interface{}{"status": status})

	if err != nil {
		session.Rollback()
		common.ErrorHandler(err)
		return
	}

	errSessionCommit := session.Commit()
	common.ErrorHandler(errSessionCommit)
	return
}

// GetRoomTypesByHotelId 查询酒店的所有房型
func GetRoomTypesByHotelId(hotelId int64) []RoomType {
	var roomTypes []RoomType
	err := dbEngine.Where("hotel_id = ? AND status = 1", hotelId).Find(&roomTypes)
	common.ErrorHandler(err)
	return roomTypes
}

// GetRoomTypeById 根据房型ID查询
func GetRoomTypeById(hotelId int64, roomTypeId int) (*RoomType, bool) {
	var roomType RoomType
	has, err := dbEngine.Where("hotel_id = ? AND room_type_id = ?", hotelId, roomTypeId).Get(&roomType)
	common.ErrorHandler(err)
	return &roomType, has
}
