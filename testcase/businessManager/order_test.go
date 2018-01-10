package testcase

import (
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
	//"maizuo.com/aura/aura-sit/common"
	//"encoding/json"
	//"maizuo.com/aura/aura-sit/model"
)



func Test_Order(t *testing.T)  {
	//Convey("Given 数据库为空", t, func() {
	//	common.ExecuteSQLScriptFile("./../../config/manager-order_clean.sql")
	//	Convey("When 根据订单id列表拉取包裹信息1 接口", func() {
	//		jsonStr := `{
	//						"orderIdList":["20171130"]
	//					}`
	//		requestAPI := "/api/v1/order/queryOrderPackageList"
	//		result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
	//
	//		Convey("Then 根据订单id列表拉取包裹信息1 判定", func() {
	//			So(result.Status, ShouldEqual, 0)
	//			resultData := result.Data.(map[string]interface{})
	//			josnbyte, _ := json.Marshal(resultData)
	//			queryOrderPackageListResp := &model.QueryOrderPackageListResp{}
	//			json.Unmarshal(josnbyte, &queryOrderPackageListResp)
	//			So(len(queryOrderPackageListResp.OrderList), ShouldEqual, 0)
	//
	//		})
	//
	//	})
	//
	//
	//})
	//
	//
	//Convey("Given 数据库数据已初始化", t, func() {
	//	common.ExecuteSQLScriptFile("./../../config/manager-order_clean.sql")
	//	common.ExecuteSQLScriptFile("./../../config/manager-order.sql")
	//	Convey("When 根据订单id列表拉取包裹信息2 接口", func() {
	//		jsonStr := `{
	//						"orderIdList":["20171130"]
	//					}`
	//		requestAPI := "/api/v1/order/queryOrderPackageList"
	//		result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
	//
	//		Convey("Then 根据订单id列表拉取包裹信息2 判定", func() {
	//			So(result.Status, ShouldEqual, 0)
	//			resultData := result.Data.(map[string]interface{})
	//			josnbyte, _ := json.Marshal(resultData)
	//			queryOrderPackageListResp := &model.QueryOrderPackageListResp{}
	//			json.Unmarshal(josnbyte, &queryOrderPackageListResp)
	//			So(len(queryOrderPackageListResp.OrderList), ShouldEqual, 1)
	//			for _, v := range queryOrderPackageListResp.OrderList {
	//				if v.OrderId == "20171130" {
	//					So(len(v.PackageList), ShouldEqual, 2)
	//					for _, packageInfo := range v.PackageList {
	//						So(packageInfo.ExpressCompanyName, ShouldEqual, "圆通速递")
	//					}
	//				}
	//			}
	//		})
	//
	//	})
	//
	//
	//	Convey("When 拉取单个订单v2接口说明1 接口", func() {
	//
	//		requestAPI := "/api/v2/queryOrder?orderId=20171130"
	//		result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)
	//
	//		Convey("Then 拉取单个订单v2接口说明1 判定", func() {
	//			So(result.Status, ShouldEqual, 0)
	//			resultData := result.Data.(map[string]interface{})
	//			josnbyte, _ := json.Marshal(resultData)
	//			order := &model.Order{}
	//			json.Unmarshal(josnbyte, &order)
	//			So(order.Id, ShouldEqual, "20171130")
	//			So(len(order.PackageList), ShouldEqual, 2)
	//			for _, packageInfo := range order.PackageList {
	//				So(packageInfo.ExpressCompanyName, ShouldEqual, "圆通速递")
	//			}
	//		})
	//
	//	})
	//
	//})

}

