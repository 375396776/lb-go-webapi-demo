package v1

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"invoice/common/response"
	"invoice/dao"
	"invoice/log"
	"invoice/model/invoice"
	"invoice/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	//测试Url
	testUrl = "https://devcustomer.51fengmi.net/api/stores/1/follow"
	//注册企业
	registerEnterpriseUrl = "http://fpkj.testnw.vpiaotong.cn/tp/openapi/register.pt"
	//新增测试记录
	insertEnterprise = "http://localhost:8089/invoices/register2"
)

type signStruct struct {
	Content      string
	PlatformCode string
	SignType     string
	Format       string
	Timestamp    string
	Version      string
	SerialNo     string
	Sign         string
}

//注册企业
func RegisterEnterprise(ctx *gin.Context) {
	var registerEnterpriseInputDTO invoice.RegisterEnterpriseInputDTO
	if err := ctx.ShouldBindJSON(&registerEnterpriseInputDTO); err != nil {
		response.FailResult(400, "请求参数缺失", ctx)
		return
	}
	log.Info.Printf("invoice#RegisterEnterprise 传入参数：%+v\n", registerEnterpriseInputDTO)
	signStruct := signStruct{}
	signStruct.Version = registerEnterpriseInputDTO.Version
	signStruct.Timestamp = registerEnterpriseInputDTO.Timestamp
	signStruct.Format = registerEnterpriseInputDTO.Format
	signStruct.SignType = registerEnterpriseInputDTO.SignType
	signStruct.PlatformCode = registerEnterpriseInputDTO.PlatformCode
	signStruct.SerialNo = registerEnterpriseInputDTO.SerialNo
	key := []byte("123456789012345678901234")                                     //定义加密密钥，必须是24byte
	contentString, err := json.Marshal(registerEnterpriseInputDTO.RegisterEntity) //待加密明文
	signStruct.Content = base64.StdEncoding.EncodeToString(util.ThriDESEnCrypt(contentString, key))
	ksortStringResult := ksort(signStruct)
	sign := util.GetSign([]byte(ksortStringResult), "/Users/libin/lbdata/private.pem")
	signStruct.Sign = base64.StdEncoding.EncodeToString(sign)
	marshal, err := json.Marshal(signStruct)
	log.Info.Printf("invoice#RegisterEnterprise 参数：%s\n", marshal)
	if registerEnterpriseInputDTO.SerialNo == "" {
		response.FailResult(400, "SerialNo不能为空", ctx)
		return
	}
	if err != nil {
		log.Error.Println("error = ", err)
	}
	res, err := http.Post(insertEnterprise, "application/json;charset=utf-8", bytes.NewBuffer(marshal))
	if err != nil {
		log.Error.Println("error = ", err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error.Println("error = ", err)
	}
	defer func() {
		err := recover() //recover内置函数可以捕获到异常
		if err != nil {  //nil是err的零值
			log.Error.Println("error = ", err)
			//TODO 这里可以发送邮件通知
		}
	}()
	response.SuccessResult(string(content), ctx)
}

func Register(ctx *gin.Context) {
	body := ctx.Request.Body
	all, err := ioutil.ReadAll(body)
	log.Info.Printf("invoice#Register 传入body：%s\n", all)
	if err != nil {
		log.Error.Println("error")
	}
	signStruct := signStruct{}
	_ = json.Unmarshal(all, &signStruct)
	signTemp := signStruct.Sign
	signStruct.Sign = ""

	key := []byte("123456789012345678901234") //定义密钥，必须是24byte
	ksortStringResult := ksort(signStruct)
	acceptmsg := []byte(ksortStringResult)

	acceptsign, err := base64.StdEncoding.DecodeString(signTemp) //接受到的签名
	if err != nil {
		log.Error.Println("error = ", err)
	}
	//验证签名
	result := util.VerifySign(acceptmsg, acceptsign, "/Users/libin/lbdata/public.pem")
	if !result {
		response.FailResult(500, "发生了某种错误,签名失败", ctx)
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(signStruct.Content)
	if err != nil {
		log.Error.Println("error = ", err)
	}
	entity := invoice.RegisterEntity{}
	crypt := util.ThriDESDeCrypt(decodeString, key)
	err = json.Unmarshal(crypt, &entity)
	if err != nil {
		log.Error.Println("error = ", err)
	}
	if entity.EnterpriseName == "" {
		response.FailResult(400, "EnterpriseName不能为空", ctx)
		return
	}
	dto := invoice.RegisterEnterpriseInputDTO{}
	dto.EnterpriseName = entity.EnterpriseName
	dto.SerialNo = signStruct.SerialNo
	insert, b := dao.Insert(&dto)
	if b {
		log.Info.Printf("插入记录条数：%d\n\n", insert)
		response.SuccessResultWithEmptyData(ctx)
		return
	}
	response.FailResult(500, "发生了某种错误", ctx)
}

func ksort(signStruct signStruct) string {
	commonMap := util.Struct2Map(signStruct)
	commonStringMap := make(map[string]string)
	for key, value := range commonMap {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		commonStringMap[strKey] = strValue
	}
	ksortStringResult := util.Ksort(commonStringMap)
	return ksortStringResult
}

//测试get
func Test(ctx *gin.Context) {
	res, err := http.Get("http://test.zhihuishangjie.cn/f-pay/getWXToken?appid=YXA63SZi3QfoTRiasCtXdj4KSQ")
	if err != nil {
		log.Error.Println("error = ", err)
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	response.SuccessResult(string(body), ctx)
}

func GetEnterpriseById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	enterprise := dao.GetById(id)
	response.SuccessResult(enterprise, ctx)
}

func UpdateEnterpriseById(ctx *gin.Context) {
	body := ctx.Request.Body
	all, err := ioutil.ReadAll(body)
	log.Info.Printf("invoice#register 传入body：%s\n", all)
	if err != nil {
		log.Error.Println("error = ", err)
	}
	ent := invoice.Enterprise{}
	_ = json.Unmarshal(all, &ent)
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	updateResult := dao.UpdateById(id, &ent)
	if !updateResult {
		response.FailResult(500, "发生了某种错误", ctx)
	}
	response.SuccessResultWithEmptyData(ctx)
}

func DelById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	delResult := dao.DelById(id)
	if !delResult {
		response.FailResult(500, "发生了某种错误", ctx)
	}
	response.SuccessResultWithEmptyData(ctx)
}
