package model

type Order struct {
	Id            string         `json:"id"`
	PayOrderId    string         `json:"payOrderId"`
	UserId        uint64         `json:"userId"`
	Name          string         `json:"name"`
	Address       string         `json:"address"`
	PostCode      string         `json:"postCode"`
	Phone         string         `json:"phone"`
	ExpressAmount uint32         `json:"expressAmount"`
	SupplierName  string         `json:"supplierName"`
	PackageList   []*PackageInfo `json:"packageList"`
}

type QueryOrderPackageListResp struct {
	OrderList []*OrderPackage `json:"orderList"`
}

type OrderPackage struct {
	OrderId     string         `json:"orderId"`
	PackageList []*PackageInfo `json:"packageList"`
}

type PackageInfo struct {
	PackageId          uint64        `json:"packageId"`
	PackageType        uint32        `json:"packageType"`
	ExpressCompanyId   uint32        `json:"expressCompanyId"`
	ExpressCompanyName string        `json:"expressCompanyName"`
	ExpressId          string        `json:"expressId"`
	SkuList            []*SkuPackage `json:"skuList"`
}

type SkuPackage struct {
	SkuId    uint64 `json:"skuId"`
	SkuCount uint32 `json:"skuCount"`
}

type QueryUserCartResp struct {
	Status int        `json:"status"`
	Msg    string     `json:"msg"`
	Data   []UserCart `json:"data"`
}

type UserCart struct {
	SkuId       int    `json:"skuId"`
	Count       int    `json:"count"`
	Price       int    `json:"price"`
	Score       int    `json:"score"`
	ProductName string `json:"productName"`
	SkuName     string `json:"skuName"`
	Image       string `json:"image"`
	RelationId  int    `json:"relationId"`
	PromotionId int    `json:"promotionId"`
	CreatedAt   int    `json:"createdAt"`
}

type QueryUserOrderResp struct {
	Total     int         `json:"total"`
	UserOrder []UserOrder `json:"list"`
}

type UserOrder struct {
	PayOrderId        string     `json:payOrderId`
	OrderId           string     `json:orderId`
	EndPayTime        int        `json:endPayTime`
	TotalPrice        int        `json:totalPrice`
	TotalScore        int        `json:totalScore`
	ExternalPayAmount int        `json:externalPayAmount`
	ExpressAmount     int        `json:expressAmount`
	OrderStatus       int        `json:orderStatus`
	PayStatus         int        `json:payStatus`
	VirtualProduct    int        `json:virtualProduct`
	CreatedAt         int        `json:createdAt`
	SupplierId        int        `json:supplierId`
	PackageType       int        `json:packageType`
	SkuList           []OrderSku `json:skuList`
}

type OrderSku struct {
	Id            int    `json:id`
	ProductId     int    `json:productId`
	ProductName   string `json:productName`
	Name          string `json:name`
	Image         string `json:image`
	Count         int    `json:count`
	Price         int    `json:price`
	FinalPrice    int    `json:finalPrice`
	MarketPrice   int    `json:marketPrice`
	Score         int    `json:score`
	ProductSource int    `json:productSource`
	AttrTypeList  []int  `json:attrTypeList`
}

type QueryBillOrderBalanceListResp struct {
	Total         int            `json:"total"`
	OrderBalances []OrderBalance `json:"list"`
}

type OrderBalance struct {
	OrderId     string `json:"orderId"`
	ServiceId   string `json:"serviceId"`
	TypeName    string `json:"typeName"`
	Amount      int    `json:"amount"`
	Remark      string `json:"remark"`
	Source      int    `json:"source"`
	Attachment  string `json:"attachment"`
	CreatedTime int    `json:"createdTime"`
}
