package model

type QueryCouponsScopeListResp struct {
	CategoryIds []uint64                 `json:"categoryIds"`
	Products    *CouponsScopeProductList `json:"products"`
}

type CouponsScopeProductList struct {
	Total                uint64                 `json:"total"`
	CouponsScopeProducts []*CouponsScopeProduct `json:"list"`
}

type CouponsScopeProduct struct {
	ProductId          uint64 `json:"productId"`
	MasterName         string `json:"masterName"`
	CategoryId         uint64 `json:"categoryId"`
	CategoryName       string `json:"categoryName"`
	ParentCategoryName string `json:"parentCategoryName"`
	Price              uint64 `json:"price"`
}

type QueryCouponCodeListInfoResp struct {
	CouponCodeList []*CouponCode `json:"couponCodeList"`
}

type CouponCode struct {
	CouponsCodeId  uint64 `json:"couponsCodeId"`
	CouponsId      uint64 `json:"couponsId"`
	DiscountAmount uint32 `json:"discountAmount"`
	DiscountRate   uint32 `json:"discountRate"`
	Title          string `json:"title"`
	LimitAmount    uint32 `json:"limitAmount"`
	ScopeType      uint32 `json:"scopeType"`
	Label          string `json:"label"`
	Remark         string `json:"remark"`
	BeginTime      string `json:"beginTime"`
	EndTime        string `json:"endTime"`
}


type QueryCouponBaseInfoByCouponIdListResp struct {
	CouponList []*Coupon `json:"couponList"`
}

type Coupon struct {
	CouponsId		uint64 `json:"couponsId"`
	Title          		string `json:"title"`
	Label          		string `json:"label"`
	DiscountType   		uint32 `json:"discountType"`
	MinDiscountAmount  	uint32 `json:"minDiscountAmount"`
	MaxDiscountAmount   	uint32 `json:"maxDiscountAmount"`
	MinDiscountRate   	uint32 `json:"minDiscountRate"`
	MaxDiscountRate   	uint32 `json:"maxDiscountRate"`
	LimitAmount    		uint32 `json:"limitAmount"`
	MaximumPerUser    	uint32 `json:"maximumPerUser"`
	BeginTime      		uint32 `json:"beginTime"`
	EndTime        		uint32 `json:"endTime"`
	ScopeType      		uint32 `json:"scopeType"`
	ValidityPeriod      	uint32 `json:"validityPeriod"`
	TotalCount        	uint32 `json:"totalCount"`
}

type QueryUserCouponBySkuListResp struct {
	CouponCount uint32		`json:"couponCount"`
	CouponList []*CouponCode 	`json:"couponList"`
}