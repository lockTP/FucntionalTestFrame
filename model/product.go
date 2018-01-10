package model

type ProductListReq struct {
	SkuCode           string   `json:"skuCode"`
	SkuId             int      `json:"skuId"`
	CategoryID        []uint32 `json:"categoryId"`
	DisplaySalesCount int      `json:"displaySalesCount"`
	PageNumber        int      `json:"pageNumber"`
	PageSize          int      `json:"pageSize"`
	ProductID         int      `json:"productId"`
	ProductName       string   `json:"productName"`
	SortKey           int      `json:"sortKey"`
	SortType          int      `json:"sortType"`
	SupplierID        int      `json:"supplierId"`
	ProductSource     int      `json:"productSource"`
	ProductStatus     int      `json:"productStatus"`
	Brand             string   `json:"brand"`
}

type ProductListFormngReq struct {
	SkuCode           string   `json:"skuCode"`
	SkuId             int      `json:"skuId"`
	CategoryID        []uint32 `json:"categoryId"`
	DisplaySalesCount int      `json:"displaySalesCount"`
	PageNumber        int      `json:"pageNumber"`
	PageSize          int      `json:"pageSize"`
	ProductID         int      `json:"productId"`
	ProductName       string   `json:"productName"`
	SortKey           int      `json:"sortKey"`
	SortType          int      `json:"sortType"`
	SupplierID        uint32   `json:"supplierId"`
	SupplierIDs       []uint32 `json:"supplierIds"`
	ProductSource     int      `json:"productSource"`
	ProductStatus     int      `json:"productStatus"`
	Brand             string   `json:"brand"`
}

type QueryProductAllSalesChannelInfoResp struct {
	List []ProductAllSalesChannelInfoList `json:"list"`
}

type ProductAllSalesChannelInfoList struct {
	SalesChannelId int          `json:"salesChannelId"`
	SkuList        []SkuChannel `json:"skuList"`
}

type SkuChannel struct {
	SkuSalesChannelId       int `json:"skuSalesChannelId"`
	SkuId                   int `json:"skuId"`
	Price                   int `json:"price"`
	Inventory               int `json:"inventory"`
	AvailableInventory      int `json:"availableInventory"`
	TotalAvailableInventory int `json:"totalAvailableInventory"`
}

type QuerySalesChannelProductInfoListResp struct {
	Total           int               `json:"total"`
	ProductInfoList []ProductInfoList `json:"productInfoList"`
}

type ProductInfoList struct {
	ProductId   int       `json:"productId"`
	ProductName string    `json:"productName"`
	SkuList     []SkuInfo `json:"skuList"`
}

type SkuInfo struct {
	SkuId                    int    `json:"skuId"`
	SkuName                  string `json:"skuName"`
	Price                    int    `json:"price"`
	Inventory                int    `json:"inventory"`
	AvailableInventory       int    `json:"availableInventory"`
	OriginPrice              int    `json:"originPrice"`
	OriginInventory          int    `json:"originInventory"`
	OriginAvailableInventory int    `json:"originAvailableInventory"`
	TotalAvailableInventory  int    `json:"totalAvailableInventory"`
}

type QueryAttrList struct {
	Total int        `json:"total"`
	List  []AttrList `json:"list"`
}

type AttrList struct {
	Total         int            `json:"total"`
	AttrId        int            `json:"attrId"`
	AttrName      string         `json:"attrName"`
	AttrStyle     int            `json:"attrStyle"`
	Remark        string         `json:"remark"`
	AttrType      []uint64       `json:"attrType"`
	AttrValueList []AttrValueObj `json:"attrValueList"`
}

type AttrValueObj struct {
	AttrValueId      int    `json:"attrValueId"`
	AttrValue        string `json:"attrValue"`
	AttrValueDisplay string `json:"attrValueDisplay"`
}

type QueryCategoryOptionList struct {
	Total int                    `json:"total"`
	List  []QueryOptionValueList `json:"list"`
}

type QueryOptionValueList struct {
	OptionId        int           `json:"optionId"`
	OptionName      string        `json:"optionName"`
	OptionValueList []OptionValue `json:"optionValueList"`
}

type OptionValue struct {
	OptionValueId int    `json:"optionValueId"`
	OptionValue   string `json:"optionValue"`
}

type FrontCategory struct {
	Id       int `json:"id"`
	ParentId int `json:"parentId"`
	Name     int `json:"name"`
	Order    int `json:"order"`
	Image    int `json:"image"`
}

type QueryProductInfoList struct {
	List []ProductInfo `json:"list"`
}

type ProductInfo struct {
	ProductId          int               `json:"productId"`
	Id                 int               `json:"id"`
	MasterName         string            `json:"masterName"`
	SlaveName          string            `json:"slaveName"`
	SupplierId         int               `json:"supplierId"`
	ProductSource      int               `json:"productSource"`
	DefaultSkuId       int               `json:"defaultSkuId"`
	SalesCount         int               `json:"salesCount"`
	AttrTypeList       []int             `json:"attrTypeList"`
	IsVirtualProduct   int               `json:"isVirtualProduct"`
	CategoryName       string            `json:"categoryName"`
	ParentCategoryName string            `json:"parentCategoryName"`
	SkuList            []SkuBusinessInfo `json:"skuList"`
}

type SkuBusinessInfo struct {
	Id                 int    `json:"id"`
	MasterName         string `json:"masterName"`
	SlaveName          string `json:"slaveName"`
	Image              string `json:"image"`
	Price              int    `json:"price"`
	Score              int    `json:"score"`
	Inventory          int    `json:"inventory"`
	AvailableInventory int    `json:"availableInventory"`
	TotalInventory     int    `json:"totalInventory"`
	SalesCount         int    `json:"salesCount"`
	MarketPrice        int    `json:"marketPrice"`
}

type ProductDetailInfo struct {
	Id               int                      `json:"id"`
	MasterName       string                   `json:"masterName"`
	SlaveName        string                   `json:"slaveName"`
	SupplierId       int                      `json:"supplierId"`
	ProductSource    int                      `json:"productSource"`
	DefaultSkuId     int                      `json:"defaultSkuId"`
	SalesCount       int                      `json:"salesCount"`
	MinPrice         int                      `json:"minPrice"`
	MaxPrice         int                      `json:"maxPrice"`
	AttrTypeList     []int                    `json:"attrTypeList"`
	IsVirtualProduct int                      `json:"isVirtualProduct"`
	ProductSkuList   []ProductSkuInfo         `json:"skuList"`
	Options          []ProductOptionValueList `json:"options"`
}

type ProductOptionValueList struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Item []string `json:"item"`
}

type ProductSkuInfo struct {
	Id                 int                  `json:"id"`
	MasterName         string               `json:"masterName"`
	SlaveName          string               `json:"slaveName"`
	Image              []string             `json:"image"`
	Price              int                  `json:"price"`
	Score              int                  `json:"score"`
	Inventory          int                  `json:"inventory"`
	AvailableInventory int                  `json:"availableInventory"`
	SalesCount         int                  `json:"salesCount"`
	MarketPrice        int                  `json:"marketPrice"`
	Options            []ProductOptionValue `json:"options"`
}

type ProductOptionValue struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type QueryProductSkuList struct {
	QueryProductSku []QueryProductSku `json:"list"`
}

type QueryProductSku struct {
	ProductBaseInfo ProductBaseInfo     `json:"productBaseInfo"`
	Sku             SkuBusinessInfo     `json:"sku"`
	SkuList         []SkuBusinessInfo   `json:"skuList"`
	ProductAttrs    []*QueryProductAttr `json:"productAttrs"`
}

type QueryProductAttr struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

type ProductBaseInfo struct {
	Id               int    `json:"id"`
	MasterName       string `json:"masterName"`
	SlaveName        string `json:"slaveName"`
	SupplierId       int    `json:"supplierId"`
	ProductSource    int    `json:"productSource"`
	AttrTypeList     []int  `json:"attrTypeList"`
	IsVirtualProduct int    `json:"isVirtualProduct"`
}

type QueryMultiProductAttrListByAttrTypeId struct {
	ProductList []ProductAttrList `json:"productList"`
}

type ProductAttrList struct {
	ProductId  int         `json:"id"`
	AttrTypeId int         `json:"attrTypeId"`
	AttrList   []AttrValue `json:"attrList"`
}

type AttrValue struct {
	AttrId    int    `json:"attrId"`
	AttrName  string `json:"attrName"`
	AttrValue string `json:"attrValue"`
}

type QueryAttrTypeList struct {
	Total        int        `json:"total"`
	AttrTypeList []AttrType `json:"list"`
}

type AttrType struct {
	AttrTypeId   int    `json:"attrTypeId"`
	AttrTypeName string `json:"attrTypeName"`
	Remark       string `json:"remark"`
}

type QueryAttrsByPIDAdnAttrRange struct {
	AttrValues []AttrValue `json:"attrs"`
}

type MasterSkuList struct {
	MasterSkuList []*MasterSku `json:"masterSkuList"`
}

type MasterSku struct {
	MasterSkuId uint64      `json:"masterSkuId"`
	SkuCounts   []*SkuCount `json:"skuList"`
}

type SkuCount struct {
	SkuId         uint64 `json:"skuId"`
	ProductId     uint64 `json:"productId"`
	MasterName    string `json:"masterName"`
	SlaveName     string `json:"slaveName"`
	SkuCode       string `json:"skuCode"`
	SkuMasterName string `json:"skuMasterName"`
	SkuSlaveName  string `json:"skuSlaveName"`
	Image         string `json:"image"`
	Count         uint32 `json:"skuCount"`
}


type QueryProductInfobySkuidlistResp struct {
	SkuList []*QueryProductInfobySkuidlistRespData `json:"list"`
}

type QueryProductInfobySkuidlistRespData struct {
	ProductId     uint64 `json:"productId"`
	SkuId         uint64 `json:"skuId"`
	MasterName    string `json:"masterName"`
	SlaveName     string `json:"slaveName"`
	SkuCode       string `json:"skuCode"`
	SkuMasterName string `json:"skuMasterName"`
	SkuSlaveName  string `json:"skuSlaveName"`
	Image         string `json:"image"`
}
