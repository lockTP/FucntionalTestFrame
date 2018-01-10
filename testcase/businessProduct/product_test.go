package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"encoding/json"
	"maizuo.com/aura/aura-sit/model"
)



func Test_Product(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func(){
		common.ExecuteSQLScriptFile("./../../config/product_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/product.sql")

		Convey("When 拉取单个品类 接口", func() {

			requestAPI := "/api/v1/1/category/2"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 拉取单个品类 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["id"], ShouldEqual, 2)
				So(resultData["parentId"], ShouldEqual, 1)
				So(resultData["name"], ShouldEqual, "category")
			})

		})

		Convey("When 拉取品类列表 接口", func() {

			requestAPI := "/api/v1/1/category"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 拉取品类列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.([]interface{})
				So(len(resultData), ShouldEqual, 3)

			})

		})


		Convey("When 获取商品列表接口1 接口", func() {
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/category/0/product"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				So(len(queryProductInfoList.List), ShouldEqual, 2) //只包含上架商品
				flag := false
				flag2 := false
				for _, v := range queryProductInfoList.List{
					if v.Id == 1002{
						flag = true
						So(v.MasterName, ShouldEqual, "sit测试商品上架")
						So(len(v.SkuList), ShouldEqual, 1)
						So(v.SkuList[0].AvailableInventory, ShouldEqual, 120)
						So(v.SkuList[0].Inventory, ShouldEqual, 120)
					}
					if v.Id == 1001 {//礼包商品测试
						flag2 = true
						So(v.SkuList[0].AvailableInventory, ShouldEqual, 12)
						So(v.SkuList[0].Inventory, ShouldEqual, 12)
					}
				}
				So(flag, ShouldEqual, true)
				So(flag2, ShouldEqual, true)
			})

		})

		Convey("When 获取商品列表接口11 接口", func() {//渠道拉取商品列表测试，渠道ID为3
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/category/0/product"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 获取商品列表接口11 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				So(len(queryProductInfoList.List), ShouldEqual, 1) //只包含上架商品
				flag := false
				for _, v := range queryProductInfoList.List{
					if v.Id == 1002{
						flag = true
						So(v.MasterName, ShouldEqual, "sit测试商品上架")
						So(len(v.SkuList), ShouldEqual, 1)
						So(v.SkuList[0].AvailableInventory, ShouldEqual, 120)
						So(v.SkuList[0].Inventory, ShouldEqual, 120)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})




		Convey("When 获取商品列表接口2 接口", func() {
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/category/2/product"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				So(len(queryProductInfoList.List), ShouldEqual, 2)
				flag := false
				for _, v := range queryProductInfoList.List{
					if v.Id == 1002{
						flag = true
						So(v.MasterName, ShouldEqual, "sit测试商品上架")
						So(len(v.SkuList), ShouldEqual, 1)
						So(v.SkuList[0].AvailableInventory, ShouldEqual, 120)
						So(v.SkuList[0].Inventory, ShouldEqual, 120)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})

		Convey("When 获取商品列表接口21 接口", func() {//渠道拉取商品列表测试，渠道ID为3
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/category/2/product"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 获取商品列表接口21 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				So(len(queryProductInfoList.List), ShouldEqual, 1)
				flag := false
				for _, v := range queryProductInfoList.List{
					if v.Id == 1002{
						flag = true
						So(v.MasterName, ShouldEqual, "sit测试商品上架")
						So(len(v.SkuList), ShouldEqual, 1)
						So(v.SkuList[0].AvailableInventory, ShouldEqual, 120)
						So(v.SkuList[0].Inventory, ShouldEqual, 120)
					}
				}
				So(flag, ShouldEqual, true)
			})

		})

		Convey("When 获取商品列表接口3 接口", func() {
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/category/4/product"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual,6003)

			})

		})




		Convey("When 获取商品列表接口4(vip product) 接口", func() {
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/vipProduct"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				So(len(queryProductInfoList.List), ShouldEqual, 2)

			})

		})

		//带渠道，后端类目底下无商品测试
		Convey("When 获取商品列表接口4 接口", func() {
			jsonStr := `{
					    "isMultiSku": 1,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/category/3/product"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual,6003)

			})

		})

		Convey("When 获取商品详细信息 接口", func() {
			requestAPI := "/api/v1/1/product/1002"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 获取商品详细信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				productDetailInfo := &model.ProductDetailInfo{}
				json.Unmarshal(josnbyte, &productDetailInfo)
				So(productDetailInfo.Id, ShouldEqual, 1002)
				So(productDetailInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(productDetailInfo.SlaveName, ShouldEqual, "sit测试商品附属上架")
				So(productDetailInfo.SalesCount, ShouldEqual, 10)
				So(len(productDetailInfo.ProductSkuList), ShouldEqual, 2)
				So(productDetailInfo.ProductSkuList[0].Inventory, ShouldEqual, 120)
				So(productDetailInfo.ProductSkuList[0].AvailableInventory, ShouldEqual, 120)
				So(len(productDetailInfo.Options), ShouldEqual, 1)
				So(len(productDetailInfo.Options[0].Item), ShouldEqual, 2)
			})
		})

		Convey("When 获取商品详细信息1 接口", func() {//礼包商品测试
			requestAPI := "/api/v1/1/product/1001"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 获取商品详细信息1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				productDetailInfo := &model.ProductDetailInfo{}
				json.Unmarshal(josnbyte, &productDetailInfo)
				So(productDetailInfo.ProductSkuList[0].Inventory, ShouldEqual, 12)
				So(productDetailInfo.ProductSkuList[0].AvailableInventory, ShouldEqual, 12)
			})
		})

		Convey("When 获取商品描述信息 接口", func() {
			requestAPI := "/api/v1/:appId/product/1002/desc"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 获取商品描述信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				So(resultData["desc"], ShouldEqual, "1002商品描述")
			})
		})

		Convey("When 获取商品购买须知 接口", func() {
			requestAPI := "/api/v1/:appId/product/1002/buyNote"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 获取商品购买须知 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				So(resultData["buyNote"], ShouldEqual, "1002商品购买须知")
			})
		})

		Convey("When 根据sku Ids 获取sku主要信息1 接口", func() {
			jsonStr := `{
					    "skuIds":["2"]
					}`
			requestAPI := "/api/v1/1/product/skuList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 根据sku Ids 获取sku主要信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].Sku.MasterName, ShouldEqual, "master_name2")
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].Sku.Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].Sku.AvailableInventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].ProductAttrs,ShouldNotBeNil)
			})

		})


		Convey("When 根据sku Ids 获取sku主要信息2 接口", func() {
			jsonStr := `{
					    "skuIds":["2"]
					}`
			requestAPI := "/api/v2/1/product/skuList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 根据sku Ids 获取sku主要信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].Sku.MasterName, ShouldEqual, "master_name2")
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].Sku.Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].Sku.AvailableInventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].ProductAttrs,ShouldBeNil)
			})

		})



		Convey("When 根据sku Ids 获取sku主要信息3 接口", func() {//带渠道测试
			jsonStr := `{
					    "skuIds":["1","2"]
					}`
			requestAPI := "/api/v2/1/product/skuList"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 根据sku Ids 获取sku主要信息3 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].Sku.MasterName, ShouldEqual, "master_name2")
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].Sku.Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].Sku.AvailableInventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].ProductAttrs,ShouldBeNil)
			})

		})

		Convey("When 根据sku Ids 获取sku主要信息4 接口", func() {//不带渠道测试
			jsonStr := `{
					    "skuIds":["1","2"]
					}`
			requestAPI := "/api/v2/1/product/skuList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 根据sku Ids 获取sku主要信息4 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 2)
			})

		})



		Convey("When 根据procduct Ids 获取product主要信息1 接口", func() {
			jsonStr := `{
					    "productIds":["1002"],
					    "isMultiSku": 1
					}`
			requestAPI := "/api/v1/1/product/productListbyIds"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 根据procduct Ids 获取product主要信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].AvailableInventory, ShouldEqual, 120)
			})

		})

		Convey("When 根据procduct Ids 获取product主要信息2 接口", func() {
			jsonStr := `{
					    "productIds":["1002"],
					    "isMultiSku": 0
					}`
			requestAPI := "/api/v1/1/product/productListbyIds"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 根据procduct Ids 获取product主要信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].AvailableInventory, ShouldEqual, 120)
			})

		})

		Convey("When 根据procduct Ids 获取product主要信息3 接口", func() {//带渠道测试
			jsonStr := `{
					    "productIds":["1002"],
					    "isMultiSku": 0
					}`
			requestAPI := "/api/v1/1/product/productListbyIds"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 根据procduct Ids 获取product主要信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].AvailableInventory, ShouldEqual, 120)
			})

		})

		Convey("When 根据procduct Ids 获取product主要信息4 接口", func() {//带渠道测试2
			jsonStr := `{
					    "productIds":["1001"],
					    "isMultiSku": 0
					}`
			requestAPI := "/api/v1/1/product/productListbyIds"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 根据procduct Ids 获取product主要信息 判定", func() {
				So(result.Status, ShouldEqual, 6003)
			})

		})

		Convey("When 根据procduct Ids 获取product主要信息5 接口", func() {//带渠道测试3
			jsonStr := `{
					    "productIds":["1001","1002"],
					    "isMultiSku": 0
					}`
			requestAPI := "/api/v1/1/product/productListbyIds"
			malldata := `ewogICAgICAgICAgICAiY2hhbm5lbElkIjoxMDAxLAogICAgICAgICAgICAic2FsZXNDaGFubmVsSWQiOjMsCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICAgICAgICAgImNsaWVudElwIjoiMTkyLjE2OC4xLjEiLAogICAgICAgICAgICAiYXBwU3RvcmVJZCI6IjAtbWFpenVvIiwKICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgICAgICAgICAidXNlcklkIjoyMDEwNjk4MTgsCiAgICAgICAgICAgICJjaXR5SWQiOjEwLAogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCAKICAgICAgICAgICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAgICAgICAgICJsb2NhdGlvblR5cGUiOjYxLAogICAgICAgICAgICAiY2FycmllciI6MQogICAgICAgfQ==`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 根据procduct Ids 获取product主要信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductSkuList := &model.QueryProductSkuList{}
				json.Unmarshal(josnbyte, &queryProductSkuList)
				So(len(queryProductSkuList.QueryProductSku), ShouldEqual, 1)
				So(queryProductSkuList.QueryProductSku[0].ProductBaseInfo.MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].Inventory, ShouldEqual, 120)
				So(queryProductSkuList.QueryProductSku[0].SkuList[0].AvailableInventory, ShouldEqual, 120)
			})

		})


		Convey("When 根据属性值获取商品列表 接口", func() {
			jsonStr := `{
					    "attrId":1001,
					    "attrValue":"这是一个属性值",
					    "isMultiSku": 0,
					    "pageSize":30,
					    "pageNumber":1,
					    "sortKey":"price",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/1/product/productListbyAttr"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 根据属性值获取商品列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				So(len(queryProductInfoList.List), ShouldEqual, 1)
				So(queryProductInfoList.List[0].MasterName, ShouldEqual, "sit测试商品上架")
				So(queryProductInfoList.List[0].SkuList[0].AvailableInventory, ShouldEqual, 120)
				So(queryProductInfoList.List[0].SkuList[0].Inventory, ShouldEqual, 120)
			})

		})


		Convey("When 拉取子品类列表 接口", func() {
			requestAPI := "/api/v1/1/subcategory?parentId=1"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PRODUCT)

			Convey("Then 拉取子品类列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.([]interface {})
				flag1 := false
				for _, data := range resultData {
					datamap := data.(map[string]interface{})
					if datamap["id"].(float64) == 2 {
						flag1 = true
						So(datamap["image"], ShouldEqual, "xxxxxy.jpg")
					}
				}
				So(flag1, ShouldEqual, true)
			})

		})


		Convey("When 查询商品指定类型的attr字段value1 接口", func() {
			jsonStr := `{
					    "productList":[
					    {
						"productId":1002,
						"attrTypeId":101
					    }]
					}`
			requestAPI := "/api/v1/1/product/queryMultiProductAttrListByAttrTypeId"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT)

			Convey("Then 查询商品指定类型的attr字段value 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryMultiProductAttrListByAttrTypeId := &model.QueryMultiProductAttrListByAttrTypeId{}
				json.Unmarshal(josnbyte, &queryMultiProductAttrListByAttrTypeId)
				So(len(queryMultiProductAttrListByAttrTypeId.ProductList), ShouldEqual, 1)
				So(len(queryMultiProductAttrListByAttrTypeId.ProductList[0].AttrList), ShouldEqual, 1)
			})

		})

		Convey("When 查询商品指定类型的attr字段value2 接口", func() {//渠道过滤属性拉取（渠道为1）
			jsonStr := `{
					    "productList":[
					    {
						"productId":1002,
						"attrTypeId":3
					    }]
					}`
			requestAPI := "/api/v1/1/product/queryMultiProductAttrListByAttrTypeId"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MSwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 查询商品指定类型的attr字段value 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryMultiProductAttrListByAttrTypeId := &model.QueryMultiProductAttrListByAttrTypeId{}
				json.Unmarshal(josnbyte, &queryMultiProductAttrListByAttrTypeId)
				So(len(queryMultiProductAttrListByAttrTypeId.ProductList), ShouldEqual, 1)
				So(len(queryMultiProductAttrListByAttrTypeId.ProductList[0].AttrList), ShouldEqual, 2)
			})

		})

		Convey("When 查询商品指定类型的attr字段value3 接口", func() {//渠道过滤属性拉取（渠道为0）
			jsonStr := `{
					    "productList":[
					    {
						"productId":1002,
						"attrTypeId":3
					    }]
					}`
			requestAPI := "/api/v1/1/product/queryMultiProductAttrListByAttrTypeId"
			malldata := `ew0KICAgICAgICAgICAgImNoYW5uZWxJZCI6MTAwMSwNCiAgICAgICAgICAgICJzYWxlc0NoYW5uZWxJZCI6MCwNCiAgICAgICAgICAgICJ2ZXJzaW9uIjoiNC44IiwNCiAgICAgICAgICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwNCiAgICAgICAgICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLA0KICAgICAgICAgICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLA0KICAgICAgICAgICAgInVzZXJJZCI6MjAxMDY5ODE4LA0KICAgICAgICAgICAgImNpdHlJZCI6MTAsDQogICAgICAgICAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LCANCiAgICAgICAgICAgICJsYXRpdHVkZSI6MjIuNTQ3MTAxLA0KICAgICAgICAgICAgImxvY2F0aW9uVHlwZSI6NjEsDQogICAgICAgICAgICAiY2FycmllciI6MQ0KICAgICAgIH0=`
			m := make(map[string]string, 0)
			m["malldata"] = malldata
			result, _ := common.AccessAPIWithDataAndBody("POST", requestAPI, jsonStr, common.BUSINESS_PRODUCT, m, "")

			Convey("Then 查询商品指定类型的attr字段value 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface {})
				josnbyte, _ := json.Marshal(resultData)
				queryMultiProductAttrListByAttrTypeId := &model.QueryMultiProductAttrListByAttrTypeId{}
				json.Unmarshal(josnbyte, &queryMultiProductAttrListByAttrTypeId)
				So(len(queryMultiProductAttrListByAttrTypeId.ProductList), ShouldEqual, 1)
				So(len(queryMultiProductAttrListByAttrTypeId.ProductList[0].AttrList), ShouldEqual, 1)
			})

		})

		//common.ExecuteSQLScriptFile("./../../config/product_clean.sql")
	})

}

