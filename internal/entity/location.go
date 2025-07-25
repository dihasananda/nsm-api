package entity

type Province struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProvinceID int    `json:"province_id"`
}

type SubDistrict struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	CityID int    `json:"city_id"`
}