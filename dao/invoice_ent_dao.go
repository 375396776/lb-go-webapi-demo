package dao

import (
	_ "github.com/go-sql-driver/mysql" // 引入没有具体用  所以前面加个_ 来表示只引入init
	"github.com/go-xorm/xorm"
	"invoice/model/invoice"
	"log"
)

//定义orm引擎
var x *xorm.Engine

const DriverName = "mysql"
const MasterDataSourceName = "root:qwe123QWE@tcp(114.215.223.220:3306)/libin?charset=utf8"

//创建orm引擎
func init() {
	var err error
	x, err = xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(invoice.Enterprise)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	//日志打印SQL
	x.ShowSQL(true)
	//设置连接池的空闲数大小
	x.SetMaxIdleConns(5)
	//设置最大打开连接数
	x.SetMaxOpenConns(5)
}

//增
func Insert(registerEnterpriseInputDTO *invoice.RegisterEnterpriseInputDTO) (int64, bool) {
	ent := new(invoice.Enterprise)
	ent.TaxpayerNum = registerEnterpriseInputDTO.TaxpayerNum
	ent.CityName = registerEnterpriseInputDTO.CityName
	ent.ContactsEmail = registerEnterpriseInputDTO.ContactsEmail
	ent.ContactsName = registerEnterpriseInputDTO.ContactsName
	ent.ContactsPhone = registerEnterpriseInputDTO.ContactsPhone
	ent.EnterpriseAddress = registerEnterpriseInputDTO.EnterpriseAddress
	ent.EnterpriseName = registerEnterpriseInputDTO.EnterpriseName
	ent.LegalPersonnAme = registerEnterpriseInputDTO.LegalPersonName
	ent.RegionCode = "01"
	ent.TaxRegistrationCertificate = registerEnterpriseInputDTO.TaxRegistrationCertificate
	ent2 := ent
	ent2.RegionCode = "02"
	affected, err := x.Insert(ent)
	if err != nil {
		return affected, false
	}
	return affected, true
}

//删
func DelById(id int) bool {
	user := &invoice.Enterprise{Id: id}
	_, err := x.Delete(user)
	if err != nil {
		return false
	}
	return true
}

//改
func UpdateById(id int, ent *invoice.Enterprise) bool {
	_, err := x.ID(id).Update(ent)
	if err != nil {
		return false
	}
	//if affected == 0 {
	//	return true
	//}
	return true
}

//查
func GetById(id int) *invoice.Enterprise {
	ent := &invoice.Enterprise{Id: id}
	is, _ := x.Get(ent)
	if !is {
		return nil
	}
	return ent
}
