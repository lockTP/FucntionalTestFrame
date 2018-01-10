package model

type AddUserAddress struct {
	UserId uint64		`json:"userId"`
	Name string		`json:"name"`
	CountryId uint32 	`json:"countryId"`
	ProvinceId uint32	`json:"provinceId"`
	CityId uint32		`json:"cityId"`
	DistrictId uint32	`json:"districtId"`
	Address string		`json:"address"`
	PostCode string		`json:"postCode"`
	Phone 	string		`json:"phone"`
	IsDefault uint32	`json:"isDefault"`
	IdCard string		`json:"idCard"`
}

type UpdateUserAddress struct {
	Id uint64		`json:"id"`
	Name string		`json:"name"`
	CountryId uint32 	`json:"countryId"`
	ProvinceId uint32	`json:"provinceId"`
	CityId uint32		`json:"cityId"`
	DistrictId uint32	`json:"districtId"`
	Address string		`json:"address"`
	PostCode string		`json:"postCode"`
	Phone 	string		`json:"phone"`
	IsDefault uint32	`json:"isDefault"`
	IdCard string		`json:"idCard"`
}

type AddUser struct {
	RegistType uint32	`json:"registType"`
	OpenId string		`json:"openId"`
	UUID string		`json:"uuId"`
	OpenType uint32		`json:"openType"`
	AccountName string	`json:"accountName"`
	Mobile string		`json:"mobile"`
	NickName string		`json:"nickName"`
	HeadIcon string		`json:"headIcon"`
	Gender uint32		`json:"gender"`
	Mail string		`json:"mail"`
	Birthday string		`json:"birthday"`
	SmsId uint32		`json:"smsId"`
	SmsCode string		`json:"smsCode"`
	Password string		`json:"password"`
}
