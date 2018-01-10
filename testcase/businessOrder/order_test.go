package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"fmt"
	"encoding/json"
	"maizuo.com/aura/aura-sit/model"
)



func Test_Order(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func(){
		common.ExecuteSQLScriptFile("./../../config/business-order_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/business-order.sql")

		Convey("When 批量添加购物车中sku 接口", func() {

			jsonStr := `{
							"userId": 1001,
							"type":0,
							"skuList": [
								{
									"skuId": 1,
									"count": 8
								},
								{
									"skuId": 2,
									"count": 12
								}
							]
						}`
			requestAPI := "/api/v1/:appId/cartBatAdd"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_ORDER)

			Convey("Then 批量添加购物车中sku 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count, skuCount, price int
				var productName string
				sql := "SELECT count(1) FROM shopping_cart WHERE user_id = ?"
				err := common.DB.QueryRow(sql, 1001).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 2)
				sql = "SELECT product_name, sku_count, price FROM shopping_cart WHERE user_id = ? AND sku_id = ?"
				err = common.DB.QueryRow(sql, 1001, 2).Scan(&productName, &skuCount, &price)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(productName, ShouldEqual, "sit测试商品2")
				So(skuCount, ShouldEqual, 12)
				So(price, ShouldEqual, 6000)
			})

		})

		Convey("When 批量删除购物车中sku1 接口", func() {

			jsonStr := `{
							"userId":1002,
							"skuIdList":[1,2]
						}`
			requestAPI := "/api/v1/:appId/cartBatDelete"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_ORDER)

			Convey("Then 批量删除购物车中sku 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM shopping_cart WHERE user_id = ? AND status = 1"
				err := common.DB.QueryRow(sql, 1002).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 2)

			})

		})

		Convey("When 批量删除购物车中sku2 接口", func() {

			jsonStr := `{
							"userId":1002,
							"skuIdList":[1]
						}`
			requestAPI := "/api/v1/:appId/cartBatDelete"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_ORDER)

			Convey("Then 批量删除购物车中sku 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM shopping_cart WHERE user_id = ? AND sku_id = ? AND status = 1"
				err := common.DB.QueryRow(sql, 1002, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)

			})

		})

		Convey("When 拉取用户购物车sku列表 接口", func() {

			requestAPI := "/api/v1/1/user/1002/cartlist"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_ORDER)

			Convey("Then 拉取用户购物车sku列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result)
				queryUserCartResp := &model.QueryUserCartResp{}
				json.Unmarshal(josnbyte, &queryUserCartResp)
				So(len(queryUserCartResp.Data), ShouldEqual, 2)
			})
		})

		Convey("When 拉取购物车商品数 接口", func() {

			requestAPI := "/api/v1/1/shopCart/querySkuCount?userId=1002"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_ORDER)

			Convey("Then 拉取购物车商品数 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["skuCount"], ShouldEqual, 4)
			})
		})

		Convey("When 拉取用户订单列表1 接口", func() {

			requestAPI := "/api/v1/1/order/user/1002"
			jsonStr := `{
							"pageSize": 20,
							"pageNumber":1,
							"sortKey": 1,
							"sortType": 0,
							"type":0
						}`
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_ORDER)

			Convey("Then 拉取用户订单列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryUserOrderResp := &model.QueryUserOrderResp{}
				json.Unmarshal(josnbyte, &queryUserOrderResp)
				So(queryUserOrderResp.Total, ShouldEqual, 1)
				So(queryUserOrderResp.UserOrder[0].OrderId, ShouldEqual, "20171211")
			})
		})

		Convey("When 拉取用户订单列表2 接口", func() {

			requestAPI := "/api/v1/1/order/user/1002"
			jsonStr := `{
							"pageSize": 20,
							"pageNumber":1,
							"sortKey": 1,
							"sortType": 0,
							"type":1
						}`
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_ORDER)

			Convey("Then 拉取用户订单列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryUserOrderResp := &model.QueryUserOrderResp{}
				json.Unmarshal(josnbyte, &queryUserOrderResp)
				So(queryUserOrderResp.Total, ShouldEqual, 1)
				So(queryUserOrderResp.UserOrder[0].OrderId, ShouldEqual, "20171213")
			})
		})
		//common.ExecuteSQLScriptFile("./../../config/product_clean.sql")
	})

}

