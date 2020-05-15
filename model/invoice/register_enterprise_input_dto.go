package invoice

//注册企业传入实体
type RegisterEntity struct {
	Id                         int16  `json:"-" binding:"-"`
	TaxpayerNum                string `json:"taxpayerNum" binding:"required"`
	EnterpriseName             string `json:"enterpriseName" binding:"required"`
	LegalPersonName            string `json:"legalPersonName" binding:"required"`
	ContactsName               string `json:"contactsName" binding:"required"`
	ContactsEmail              string `json:"contactsEmail" binding:"required"`
	ContactsPhone              string `json:"contactsPhone" binding:"required"`
	CityName                   string `json:"cityName" binding:"required"`
	EnterpriseAddress          string `json:"enterpriseAddress" binding:"required"`
	TaxRegistrationCertificate string `json:"taxRegistrationCertificate" binding:"required"`
}

type RegisterEnterpriseInputDTO struct {
	Common
	RegisterEntity `json:"content"`
}
