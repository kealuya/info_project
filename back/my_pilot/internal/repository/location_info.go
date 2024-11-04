package repository

import (
	"my_pilot/common"
	"time"
)

type LocationInfo struct {
	Id        int64     `xorm:"pk autoincr 'id'"`
	CountryId int       `xorm:"'country_id' not null"`
	CountryCn string    `xorm:"'country_cn' varchar(100)"`
	CountryEn string    `xorm:"'country_en' varchar(100)"`
	StateId   int       `xorm:"'state_id'"`
	StateCn   string    `xorm:"'state_cn' varchar(100)"`
	StateEn   string    `xorm:"'state_en' varchar(100)"`
	CityId    int       `xorm:"'city_id'"`
	CityCn    string    `xorm:"'city_cn' varchar(100)"`
	CityEn    string    `xorm:"'city_en' varchar(100)"`
	Created   time.Time `xorm:"created"`
	Updated   time.Time `xorm:"updated"`
}

// GetDistinctStatesByCountry 获取指定国家的去重省份列表
func GetDistinctStatesByCountry(countryId int) ([]LocationInfo, error) {
	var locationInfo []LocationInfo
	err := dbEngine.Distinct("state_id", "state_cn", "state_en").
		Table("location_info").
		Where("country_id = ?", countryId).
		Find(&locationInfo)
	return locationInfo, err
}

func InsertLocations(locations []LocationInfo) {

	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	for _, loc := range locations {
		_, err := session.Insert(&loc)
		if err != nil {
			session.Rollback()
			common.ErrorHandler(errSessionBegin)
		}
	}
	errSessionCommit := session.Commit()
	common.ErrorHandler(errSessionCommit)
	return
}
