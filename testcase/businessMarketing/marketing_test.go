package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"encoding/json"
	"maizuo.com/aura/aura-sit/model"
	"fmt"
)



func Test_Marketing(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func() {
		common.ExecuteSQLScriptFile("./../../config/business-coupon_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/business-coupon.sql")

		Convey("When 1.领取优惠券 接口", func() {
			requestAPI := "/api/v1/0/marketing/ApplyCoupon"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			formData := "userId=80&couponId=1"
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, "", common.BUSINESS_Marketing, headerMap, formData)
			Convey("Then 领取优惠券 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData, ShouldNotBeNil)
			})

		})

		Convey("When 2.兑换优惠券 接口", func() {
			requestAPI := "/api/v1/0/marketing/ExchangeCouponCode"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			formData := "userId=80&couponCode=NFYHS123B"
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, "", common.BUSINESS_Marketing, headerMap, formData)
			Convey("Then 兑换优惠券 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData, ShouldNotBeNil)
			})

		})

		Convey("When 3.查询优惠券信息 接口", func() {
			requestAPI := "/api/v1/0/marketing/QueryCouponCodeListInfo"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`

			jsonStr := `{
							"userId":1001,
							"couponCodeIdList":[3,8]
						}`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_Marketing, headerMap, "")
			Convey("Then 查询优惠券信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				fmt.Println(result.Data, "-", result.Msg)
				resultData := result.Data.(map[string]interface{})
				So(resultData["couponCodeList"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(result.Data)
				queryCouponCodeListInfoResp := &model.QueryCouponCodeListInfoResp{}
				json.Unmarshal(josnbyte, &queryCouponCodeListInfoResp)
				So(len(queryCouponCodeListInfoResp.CouponCodeList), ShouldEqual, 2)
				flag := false
				for _, couponCode := range queryCouponCodeListInfoResp.CouponCodeList {
					if couponCode.CouponsCodeId == 3 {
						flag = true
						So(couponCode.ScopeType, ShouldEqual, 3)
						So(couponCode.LimitAmount, ShouldEqual, 800)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})

		Convey("When 4.查询用户优惠券列表 接口", func() {
			requestAPI := "/api/v1/0/marketing/QueryUserCoupon"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			requestAPI += "?pageNumber=1&pageSize=10&sortKey=1&sortType=desc&type=3&userId=1002"
			result, _ := common.AccessAPIWithDataAndBody("GET", requestAPI, "", common.BUSINESS_Marketing, headerMap, "")
			Convey("Then 查询用户优惠券列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				fmt.Println(result.Data, "-", result.Msg)
				resultData := result.Data.(map[string]interface{})
				So(resultData["couponList"], ShouldNotBeNil)
				So(resultData["couponCount"], ShouldEqual, 2)
			})

		})

		Convey("When 5.查询用户优惠券数量 接口", func() {
			requestAPI := "/api/v1/0/marketing/QueryUserCouponCount"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			requestAPI += "?type=3&userId=1002"
			result, _ := common.AccessAPIWithDataAndBody("GET", requestAPI, "", common.BUSINESS_Marketing, headerMap, "")
			Convey("Then 查询用户优惠券数量 判定", func() {
				So(result.Status, ShouldEqual, 0)
				fmt.Println(result.Data, "-", result.Msg)
				resultData := result.Data.(map[string]interface{})
				So(resultData["count"], ShouldEqual, 2)
			})

		})

		Convey("When 6.根据skuId列表查询用户可用的优惠券列表 接口", func() {
			requestAPI := "/api/v1/0/marketing/QueryUserCouponBySkuList"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`

			jsonStr := `{
   					 "userId":1001,
					 "skuList":[

        					{
         					   "skuId":1,
         					   "skuCount":1
       						 },
       						 {
         					   "skuId":2,
         					   "skuCount":1
       						 }]
					}`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_Marketing, headerMap, "")
			Convey("Then 根据skuId列表查询用户可用的优惠券列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				fmt.Println(result.Data, "-", result.Msg)
				resultData := result.Data.(map[string]interface{})
				So(resultData["couponList"], ShouldNotBeNil)
				jsonByte, _ := json.Marshal(result.Data)
				queryUserCouponBySkuListResp := &model.QueryUserCouponBySkuListResp{}
				json.Unmarshal(jsonByte, &queryUserCouponBySkuListResp)
				So(len(queryUserCouponBySkuListResp.CouponList), ShouldEqual, 2)
				flag := false
				for _, couponCode := range queryUserCouponBySkuListResp.CouponList {
					fmt.Println(couponCode)
					if couponCode.CouponsId == 1 {
						flag = true
						So(couponCode.ScopeType, ShouldEqual, 3)
						So(couponCode.LimitAmount, ShouldEqual, 800)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})

		Convey("When 7.根据skuid列表查询用户指定优惠券的使用明细 接口", func() {
			requestAPI := "/api/v1/0/marketing/QueryUserCouponUsedDetailBySkuList"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`

			jsonStr := `{
					"userId": 1001,
					"couponSkuList": [
					{
					    "couponCodeId": 3,
 					   "skuList": [
						{
						    "skuId": 1,
						    "skuCount": 1
						},
						{
						    "skuId": 2,
						    "skuCount": 1
						}
					    ]
					}]
				}`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_Marketing, headerMap, "")
			Convey("Then 根据skuid列表查询用户指定优惠券的使用明细 判定", func() {
				So(result.Status, ShouldEqual, 0)
				fmt.Println(result.Data, "-", result.Msg)
				resultData := result.Data.(map[string]interface{})
				So(resultData["couponDetailInfoList"], ShouldNotBeNil)
				So(resultData["sign"], ShouldNotEqual, "")
			})

		})

		Convey("When 8.根据couponId列表查询优惠券信息 接口", func() {
			requestAPI := "/api/v1/0/marketing/QueryCouponBaseInfoByCouponIdList"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`

			jsonStr := `{
							"couponIdList":[1,2]
						}`
			headerMap := make(map[string]string, 0)
			headerMap["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_Marketing, headerMap, "")
			Convey("Then 根据couponId列表查询优惠券信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				fmt.Println(result.Data, "-", result.Msg)
				resultData := result.Data.(map[string]interface{})
				So(resultData["couponList"], ShouldNotBeNil)
				jsonByte, _ := json.Marshal(result.Data)
				queryCouponBaseInfoByCouponIdListResp := &model.QueryCouponBaseInfoByCouponIdListResp{}
				json.Unmarshal(jsonByte, &queryCouponBaseInfoByCouponIdListResp)
				So(len(queryCouponBaseInfoByCouponIdListResp.CouponList), ShouldEqual, 2)
				flag := false
				for _, couponCode := range queryCouponBaseInfoByCouponIdListResp.CouponList {
					fmt.Println(couponCode)
					if couponCode.CouponsId == 1 {
						flag = true
						So(couponCode.ScopeType, ShouldEqual, 3)
						So(couponCode.LimitAmount, ShouldEqual, 800)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})

		//common.ExecuteSQLScriptFile("./../../config/coupon_clean.sql")
	})

}

