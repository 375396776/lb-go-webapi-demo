package invoice

//注册企业传入实体
type invoiceBlueEntity struct {
	TaxpayerNum        string `json:"taxpayerNum"`
	InvoiceReqSerialNo string `json:"invoiceReqSerialNo"`
	BuyerName          string `json:"buyerName"`
	BuyerTaxpayerNum   string `json:"buyerTaxpayerNum"`
	BuyerAddress       string `json:"buyerAddress"`
	BuyerTel           string `json:"buyerTel"`
	BuyerBankName      string `json:"buyerBankName"`
	BuyerBankAccount   string `json:"buyerBankAccount"`
	SellerAddress      string `json:"sellerAddress"`
	SellerTel          string `json:"sellerTel"`
	SellerBankName     string `json:"sellerBankName"`
	SellerBankAccount  string `json:"sellerBankAccount"`
	CasherName         string `json:"casherName"`
	ReviewerName       string `json:"reviewerName"`
	DrawerName         string `json:"drawerName"`
	TakerName          string `json:"takerName"`
	TakerTel           string `json:"takerTel"`
	TakerEmail         string `json:"takerEmail"`
	ItemList           []item `json:"itemList"`
}

type item struct {
	GoodsName             string `json:"goodsName"`
	TaxClassificationCode string `json:"taxClassificationCode"`
	SpecificationModel    string `json:"specificationModel"`
	MeteringUnit          string `json:"meteringUnit"`
	Quantity              string `json:"quantity"`
	UnitPrice             string `json:"unitPrice"`
	InvoiceAmount         string `json:"invoiceAmount"`
	TaxRateValue          string `json:"taxRateValue"`
}
type BlueInputDTO struct {
	Common
	invoiceBlueEntity `json:"content"`
}
