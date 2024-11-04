package repository

import (
	"time"
)

type HotelStaticInfo struct {
	Id                int64     `xorm:"pk autoincr 'id'"`
	HotelId           int64     `xorm:"'hotel_id' not null"`
	ThemeType         string    `xorm:"'theme_type' varchar(100)"`
	HotelNameCn       string    `xorm:"'hotel_name_cn' varchar(200) not null"`
	HotelNameEn       string    `xorm:"'hotel_name_en' varchar(400)"`
	Star              string    `xorm:"'star' varchar(50)"`
	AddressCn         string    `xorm:"'address_cn' varchar(500)"`
	AddressEn         string    `xorm:"'address_en' varchar(1000)"`
	CountryId         int       `xorm:"'country_id'"`
	StateId           int       `xorm:"'state_id'"`
	CityId            int       `xorm:"'city_id'"`
	Area              int       `xorm:"'area'"`
	BusinessCircle    int       `xorm:"'business_circle'"`
	Phone             string    `xorm:"'phone' varchar(100)"`
	Longitude         string    `xorm:"'longitude' varchar(50)"`
	Latitude          string    `xorm:"'latitude' varchar(50)"`
	Fax               string    `xorm:"'fax' varchar(100)"`
	Email             string    `xorm:"'email' varchar(200)"`
	PostCode          string    `xorm:"'post_code' varchar(50)"`
	CheckPolicy       string    `xorm:"'check_policy' text"`
	ChildrenPolicy    string    `xorm:"'children_policy' text"`
	PetPolicy         string    `xorm:"'pet_policy' text"`
	EstablishmentDate string    `xorm:"'establishment_date' varchar(50)"`
	RenovationDate    string    `xorm:"'renovation_date' varchar(50)"`
	HotelGroup        string    `xorm:"'hotel_group' varchar(200)"`
	HotelBrand        string    `xorm:"'hotel_brand' varchar(200)"`
	Facilities        string    `xorm:"'facilities' text"`
	CardType          string    `xorm:"'card_type' varchar(500)"`
	IntroduceCn       string    `xorm:"'introduce_cn' text"`
	IntroduceEn       string    `xorm:"'introduce_en' text"`
	SellType          int       `xorm:"'sell_type' default(0)"`
	Status            int       `xorm:"'status' default(1)"`
	Created           time.Time `xorm:"created"`
	Updated           time.Time `xorm:"updated"`
}

// InsertHotelStaticInfo 插入单个酒店静态信息
func InsertHotelStaticInfo(hotel *HotelStaticInfo) error {
	_, err := dbEngine.Insert(hotel)
	return err
}

// BatchInsertHotelStaticInfo 批量插入酒店静态信息
func BatchInsertHotelStaticInfo(hotels []HotelStaticInfo) error {
	session := dbEngine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	for _, hotel := range hotels {
		if _, err := session.Insert(&hotel); err != nil {
			session.Rollback()
			return err
		}
	}

	return session.Commit()
}

// GetHotelStaticInfo 根据酒店ID获取静态信息
func GetHotelStaticInfo(hotelId int64) (*HotelStaticInfo, bool, error) {
	var hotel HotelStaticInfo
	has, err := dbEngine.Where("hotel_id = ? AND status = 1", hotelId).Get(&hotel)
	return &hotel, has, err
}

// GetHotelsByCity 根据城市获取酒店列表
func GetHotelsByCity(cityId int) ([]HotelStaticInfo, error) {
	var hotels []HotelStaticInfo
	err := dbEngine.Where("city_id = ? AND status = 1", cityId).Find(&hotels)
	return hotels, err
}

// UpdateHotelStaticInfo 更新酒店静态信息
func UpdateHotelStaticInfo(hotel *HotelStaticInfo) error {
	_, err := dbEngine.ID(hotel.Id).Update(hotel)
	return err
}

// UpdateHotelStatus 更新酒店状态
func UpdateHotelStatus(hotelId int64, status int) error {
	_, err := dbEngine.Table("hotel_static_info").
		Where("hotel_id = ?", hotelId).
		Update(map[string]interface{}{"status": status})
	return err
}

// GetHotelsByGroup 根据集团查询酒店列表
func GetHotelsByGroup(hotelGroup string) ([]HotelStaticInfo, error) {
	var hotels []HotelStaticInfo
	err := dbEngine.Where("hotel_group = ? AND status = 1", hotelGroup).Find(&hotels)
	return hotels, err
}

// GetHotSaleHotels 获取热销酒店列表
func GetHotSaleHotels() ([]HotelStaticInfo, error) {
	var hotels []HotelStaticInfo
	err := dbEngine.Where("sell_type = 1 AND status = 1").Find(&hotels)
	return hotels, err
}

// GetHotelsByLocation 按地理位置范围查询酒店
func GetHotelsByLocation(cityId int, areaIds []int) ([]HotelStaticInfo, error) {
	var hotels []HotelStaticInfo
	err := dbEngine.Where("city_id = ? AND area IN (?) AND status = 1", cityId, areaIds).Find(&hotels)
	return hotels, err
}

// SearchHotels 搜索酒店
func SearchHotels(keyword string) ([]HotelStaticInfo, error) {
	var hotels []HotelStaticInfo
	err := dbEngine.Where("status = 1").
		And("(hotel_name_cn LIKE ? OR hotel_name_en LIKE ? OR address_cn LIKE ?)",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Find(&hotels)
	return hotels, err
}
