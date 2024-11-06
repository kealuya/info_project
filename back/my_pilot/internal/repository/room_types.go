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

func InsertRoomTypes(roomTypes []RoomType) {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	for _, roomType := range roomTypes {
		_, err := session.Insert(&roomType)
		if err != nil {
			session.Rollback()
			//common.ErrorHandler(err)
			logs.Error("roomType insert failed.", fmt.Sprintf("%+v", roomType))
		}
	}

	errSessionCommit := session.Commit()
	common.ErrorHandler(errSessionCommit)
	return
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
