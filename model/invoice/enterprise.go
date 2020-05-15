package invoice

type Enterprise struct {
	Id                         int    `xorm:"not null pk autoincr comment('主键') INT(10)"`
	TaxpayerNum                string `xorm:"comment('纳税识别号') VARCHAR(64)"`
	EnterpriseName             string `xorm:"comment('企业名称') VARCHAR(64)"`
	LegalPersonnAme            string `xorm:"comment('法人名称') VARCHAR(64)"`
	ContactsName               string `xorm:"comment('联系人名称') VARCHAR(64)"`
	ContactsEmail              string `xorm:"comment('联系人邮箱') VARCHAR(64)"`
	ContactsPhone              string `xorm:"comment('联系人手机') VARCHAR(11)"`
	RegionCode                 string `xorm:"comment('地区编码') VARCHAR(2)"`
	CityName                   string `xorm:"comment('市（地区）名称') VARCHAR(32)"`
	EnterpriseAddress          string `xorm:"comment('详细地址') VARCHAR(64)"`
	TaxRegistrationCertificate string `xorm:"comment('税务登记照片') VARCHAR(64)"`
}
