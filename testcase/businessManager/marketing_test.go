package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"strconv"
	"encoding/json"
	"maizuo.com/aura/aura-sit/model"
	"fmt"
)



func Test_Marketing(t *testing.T)  {
	Convey("Given 数据库为空", t, func() {

		common.ExecuteSQLScriptFile("./../../config/coupon_clean.sql")
		Convey("When 优惠券列表查询 接口", func() {
			requestAPI := "/api/v1/coupons/list"
			jsonStr := `{
							"pageSize": 20,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsStatus": 9999,
							"title": ""
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券列表查询 判定", func() {
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				So(resultData["total"], ShouldEqual, 0)
			})

		})


	})


	Convey("Given 数据库数据已初始化", t, func() {
		common.ExecuteSQLScriptFile("./../../config/coupon_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/coupon.sql")
		Convey("When 优惠券列表查询1 接口", func() {
			requestAPI := "/api/v1/coupons/list"
			jsonStr := `{
							"pageSize": 20,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsStatus": 9999,
							"title": ""
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券列表查询 判定", func() {
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				So(resultData["total"], ShouldEqual, 4)
			})

		})

		Convey("When 优惠券列表查询2 接口", func() {
			requestAPI := "/api/v1/coupons/list"
			jsonStr := `{
							"pageSize": 20,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsStatus": 9999,
							"title": "title"
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券列表查询 判定", func() {
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				So(resultData["total"], ShouldEqual, 1)
			})

		})

		Convey("When 优惠券列表查询3 接口", func() {
			requestAPI := "/api/v1/coupons/list"
			jsonStr := `{
							"pageSize": 20,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsStatus": 2,
							"title": ""
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券列表查询 判定", func() {
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				So(resultData["total"], ShouldEqual, 1)
			})

		})

		Convey("When 优惠券详情查询1 接口", func() {
			couponsId := 1
			requestAPI := "/api/v1/coupons/query?couponsId="+ strconv.Itoa(couponsId)
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情查询 判定", func() {
				resultData := result.Data.(map[string]interface{})
				So(resultData["couponsId"], ShouldEqual, 1)
				So(resultData["title"], ShouldEqual, "sit优惠券描述")
				So(resultData["couponsStatus"], ShouldEqual, 0)
			})

		})

		Convey("When 优惠券详情查询2 接口", func() {
			couponsId := 99
			requestAPI := "/api/v1/coupons/query?couponsId="+ strconv.Itoa(couponsId)
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情查询 判定", func() {
				So(result.Status, ShouldEqual, 6003)
			})

		})


		Convey("When 优惠券详情新增1 接口", func() {
			requestAPI := "/api/v1/coupons/create"
			jsonStr := `{
							"title": "优惠券名称（促销描述）",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 0,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 10,
							"salesChannelId": 1,
							"returnFlag": 0,
							"scopeType": 0,
    						"couponsCode" : "asdasdw"
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情新增 判定", func() {
				So(result.Status, ShouldEqual, 0)
				couponsId := uint64(result.Data.(float64))
				var count int
				sql := "SELECT count(id) FROM coupons WHERE title like ? AND status = 0"
				err := common.DB.QueryRow(sql, "%优惠券名称（促销描述）%").Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
				sql = "SELECT count(id) FROM coupons_code WHERE coupons_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, &couponsId).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)

			})

		})

		Convey("When 优惠券详情新增2 接口", func() {
			requestAPI := "/api/v1/coupons/create"
			jsonStr := `{
							"title": "优惠券名称（促销描述）",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 1,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 10,
							"salesChannelId": 1,
							"returnFlag": 0,
							"scopeType": 0,
    						"couponsCode" : "asdasdw"
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情新增 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM coupons WHERE title like ? AND status = 0"
				err := common.DB.QueryRow(sql, "%优惠券名称（促销描述）%").Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
			})

		})

		Convey("When 优惠券详情新增3 接口", func() {//渠道内统一码唯一性校验
			requestAPI := "/api/v1/coupons/create"
			jsonStr := `{
							"title": "优惠券名称（促销描述）",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 2,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 10,
							"salesChannelId": 0,
							"returnFlag": 0,
							"scopeType": 0,
    						"couponsCode" : "这是一个重复优惠券码！"
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情新增 判定", func() {
				So(result.Status, ShouldEqual, 6606030)
			})

		})


		Convey("When 优惠券详情新增4 接口", func() {//渠道内统一码唯一性校验
			requestAPI := "/api/v1/coupons/create"
			jsonStr := `{
							"title": "优惠券名称（促销描述）",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 2,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 10,
							"salesChannelId": 1,
							"returnFlag": 0,
							"scopeType": 0,
    						"couponsCode" : "这是一个重复优惠券码！"
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情新增 判定", func() {
				So(result.Status, ShouldEqual, 6606030)
			})

		})

		Convey("When 优惠券详情更新1 接口", func() {//未激活可修改
			requestAPI := "/api/v1/coupons/update"
			jsonStr := `{
							"couponsId":1,
							"title": "优惠券名称（促销描述）",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 2,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 15,
							"salesChannelId": 1,
							"returnFlag": 0,
							"scopeType": 0,
    						"couponsCode" : "asdasdw"
						}`

			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情更新 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var title string
				sql := "SELECT title FROM coupons WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&title)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(title, ShouldEqual, "优惠券名称（促销描述）")
				var count int
				sql = "SELECT count(id) FROM coupons_code WHERE coupons_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})


		Convey("When 优惠券详情更新2 接口", func() {//已开始
			requestAPI := "/api/v1/coupons/update"
			jsonStr := `{
							"couponsId":2,
							"title": "sit优惠券更新描述",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 2,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 5,
							"salesChannelId": 1,
							"returnFlag": 0,
							"scopeType": 0,
    						"couponsCode" : "asdasdw"
						}`

			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情更新 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var title string
				var totalCount uint32
				sql := "SELECT title, total_count FROM coupons WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&title, &totalCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(title, ShouldEqual, "title优惠券描述")
				So(totalCount, ShouldEqual, 5)
				//激活允许修改码数量
				var count int
				sql = "SELECT count(id) FROM coupons_code WHERE coupons_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 5)
				sql = "SELECT coupons_code FROM coupons_code WHERE coupons_id = ? AND status = 0 "
				rows, err := common.DB.Query(sql, 2)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer rows.Close()
				for rows.Next() {
					var code string
					rows.Scan(&code)
					So(code, ShouldEqual, "这是一个重复优惠券码！")
				}
				rows.Close()

			})

		})

		Convey("When 优惠券详情更新3 接口", func() {//更改适用范围删除商品属性
			requestAPI := "/api/v1/coupons/update"
			jsonStr := `{
							"couponsId":1,
							"title": "sit优惠券更新描述",
							"discountType": 0,
							"minDiscountAmount": 100,
							"maxDiscountAmount": 100,
							"minDiscountRate": 0,
							"maxDiscountRate": 0,
							"limitAmount": 1000,
							"maximumPerUser": 1,
							"codeType": 2,
							"beginTime": 1490759799,
							"endTime": 1490759899,
							"totalCount": 5,
							"salesChannelId": 1,
							"returnFlag": 0,
							"scopeType": 1,
    						"couponsCode" : "asdasdw"
						}`

			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券详情更新 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var title string
				var totalCount uint32
				sql := "SELECT title, total_count FROM coupons WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&title, &totalCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				var count uint32
				So(title, ShouldEqual, "sit优惠券更新描述")
				So(totalCount, ShouldEqual, 5)
				//商品优惠券属性
				sql = "SELECT count(1) FROM product_attr WHERE attr_id IN(?,?) AND attr_value = ? AND status = 0"
				err = common.DB.QueryRow(sql, 22,23, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})


		Convey("When 优惠券失效1 接口", func() {
			requestAPI := "/api/v1/coupons/invalid"
			jsonStr := `{
							"couponsId":2
						}`

			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券失效 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var couponsStatus, count int
				sql := "SELECT coupons_status FROM coupons WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&couponsStatus)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(couponsStatus, ShouldEqual, 1)
				sql = "SELECT count(1) FROM coupons_code WHERE coupons_id = ? AND code_status = 4 AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 2)

			})

		})

		Convey("When 优惠券范围列表查询1 接口", func() {
			requestAPI := "/api/v1/coupons/queryCouponsScopeList"
			jsonStr := `{
							"pageSize": 10,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsId": 1,
							"scopeType": 2,
							"categoryIds":[2],
							"productStatus":9999
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券范围列表查询 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				queryCouponsScopeListResp := &model.QueryCouponsScopeListResp{}
				json.Unmarshal(josnbyte, &queryCouponsScopeListResp)
				So(queryCouponsScopeListResp.Products.Total, ShouldEqual, 2)
				flag := false
				for _, v := range queryCouponsScopeListResp.Products.CouponsScopeProducts {
					if v.ProductId == 1001 {
						flag = true
						So(v.CategoryId, ShouldEqual, 2)
						So(v.CategoryName, ShouldEqual, "category")
						So(v.Price, ShouldEqual, 100)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})


		Convey("When 优惠券范围列表查询2 接口", func() {
			requestAPI := "/api/v1/coupons/queryCouponsScopeList"
			jsonStr := `{
							"pageSize": 10,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsId": 1,
							"scopeType": 1
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券范围列表查询 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				queryCouponsScopeListResp := &model.QueryCouponsScopeListResp{}
				json.Unmarshal(josnbyte, &queryCouponsScopeListResp)
				So(len(queryCouponsScopeListResp.CategoryIds), ShouldEqual, 1)
			})

		})


		Convey("When 优惠券不在范围内的商品查列表询1 接口", func() {
			requestAPI := "/api/v1/coupons/queryNotCouponsScopeList"
			jsonStr := `{
							"pageSize": 10,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"couponsId": 1,
							"productStatus":9999
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券不在范围内的商品查列表询 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				couponsScopeProductList := &model.CouponsScopeProductList{}
				json.Unmarshal(josnbyte, &couponsScopeProductList)
				So(couponsScopeProductList.Total, ShouldEqual, 2)
				flag := false
				for _, v := range couponsScopeProductList.CouponsScopeProducts {
					if v.ProductId == 1003 {
						flag = true
						So(v.CategoryId, ShouldEqual, 2)
						So(v.CategoryName, ShouldEqual, "category")
						So(v.Price, ShouldEqual, 100)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})


		Convey("When 优惠券范围批量新增1 接口", func() {
			requestAPI := "/api/v1/coupons/AddCouponsScopeList"
			jsonStr := `{
							"couponsId": 1,
							"scopeType": 2,
							"objectIds": [1003]
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券范围批量新增 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM coupons_scope WHERE object_id = ? AND type = 1 AND status = 0"
				err := common.DB.QueryRow(sql, 1003).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
				var attrValue string
				sql = "SELECT attr_value FROM product_attr WHERE product_id = ? AND attr_id = 23 AND status = 0"
				err = common.DB.QueryRow(sql, 1003).Scan(&attrValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrValue, ShouldEqual, "1")
			})

		})

		Convey("When 优惠券范围批量新增2 接口", func() {
			requestAPI := "/api/v1/coupons/AddCouponsScopeList"
			jsonStr := `{
							"couponsId": 1,
							"scopeType": 1,
							"objectIds": [3]
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券范围批量新增 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM coupons_scope WHERE object_id = ? AND type = 0 AND status = 0"
				err := common.DB.QueryRow(sql, 3).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
				var attrValue string
				sql = "SELECT attr_value FROM product_attr WHERE product_id = ? AND attr_id = 22 AND status = 0"
				err = common.DB.QueryRow(sql, 1004).Scan(&attrValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrValue, ShouldEqual, "1")
			})

		})


		Convey("When  优惠券范围批量删除1 接口", func() {
			requestAPI := "/api/v1/coupons/DeleteCouponsScopeList"
			jsonStr := `{
							"couponsId": 1,
							"scopeType": 2,
							"objectIds": [1001]
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  优惠券范围批量删除 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM coupons_scope WHERE object_id = ? AND type = 1 AND status = 1"
				err := common.DB.QueryRow(sql, 1001).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
				var attrValue int
				sql = "SELECT count(1) FROM product_attr WHERE product_id = ? AND attr_id = 23 AND attr_value = 1 AND status = 1"
				err = common.DB.QueryRow(sql, 1001).Scan(&attrValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrValue, ShouldEqual, 1)
			})

		})

		Convey("When 优惠券激活1 接口", func() {
			requestAPI := "/api/v1/coupons/activate"
			jsonStr := `{
							"couponsId":1
						}`

			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券激活 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM coupons_code WHERE coupons_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 10)
				var couponStatus int
				sql = "SELECT coupons_status FROM coupons WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1).Scan(&couponStatus)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(couponStatus, ShouldEqual, 2)
			})

		})


		Convey("When 优惠券激活2 接口", func() {//类目范围的优惠券，但是没选择类目激活时失败
			requestAPI := "/api/v1/coupons/activate"
			jsonStr := `{
							"couponsId":3
						}`

			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 优惠券激活 判定", func() {
				So(result.Status, ShouldEqual, 6606032)
			})

		})

		//common.ExecuteSQLScriptFile("./../../config/coupon_clean.sql")
	})

}

