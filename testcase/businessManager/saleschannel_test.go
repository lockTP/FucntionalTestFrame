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



func Test_Saleschannel(t *testing.T)  {
	Convey("Given 数据库为空", t, func() {

		Convey("When 查询销售渠道信息 接口", func() {
			pageSize := 10
			pageNumber := 1
			sortKey := 1
			sortType := 1
			requestAPI := "/api/v1/saleschannel/queryList?pageSize="+ strconv.Itoa(pageSize) + "&pageNumber=" + strconv.Itoa(pageNumber) + "&sortKey=" + strconv.Itoa(sortKey) + "&sortType=" + strconv.Itoa(sortType)
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)

			Convey("Then 查询销售渠道信息结果 判定", func() {
				resultData := result.Data.(map[string]interface{})
				So(resultData["salesChannelInfoList"], ShouldNotBeNil)
			})

		})

		Convey("When 创建销售渠道信息 接口", func() {
			jsonStr := "{\"channelName\":\"convey渠道测试\",\"channelStatus\":1,\"channelDesc\":\"convey渠道测试描述\"}"
			requestAPI := "/api/v1/saleschannel/create"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 创建销售渠道信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})


	})


	Convey("Given 数据库数据已初始化", t, func() {
		common.ExecuteSQLScriptFile("./../../config/salesChannel_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/salesChannel.sql")
		Convey("When 修改销售渠道信息 接口", func() {
			requestAPI := "/api/v1/saleschannel/update"
			jsonStr := `{
							"salesChannelId": 101,
							"channelName": "新渠道名",
							"channelStatus": 1,
							"channelDesc": "新渠道描述"
						}`

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 修改销售渠道信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var channel_names string
				var channel_desc string
				sql := "SELECT channel_names, channel_desc FROM sales_channel_info WHERE id = ? AND status = 0"
				common.DB.QueryRow(sql, 101).Scan(&channel_names, &channel_desc)
				So(channel_names, ShouldEqual, "新渠道名")
				So(channel_desc, ShouldEqual, "新渠道描述")

			})

		})

		Convey("When 删除销售渠道信息 接口", func() {
			requestAPI := "/api/v1/saleschannel/delete"
			jsonStr := "{\"salesChannelId\":102}"

			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 删除销售渠道信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM sales_channel_info WHERE id = ? AND status = 0"
				common.DB.QueryRow(sql, 102).Scan(&count)
				So(count, ShouldEqual, 0)

			})

		})

		Convey("When 批量添加销售渠道商品信息1 接口", func() {
			jsonStr := `{
					"list":[{
						"salesChannelId":100,
					    "skuList":[
					    {
						"skuId":1,
						"price":1000,
						"availableInventory":200
					    },{
						"skuId":3,
						"price":1000,
						"availableInventory":-1
						}]
					}]

				}`
			requestAPI := "/api/v1/saleschannelproduct/listAdd"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 批量添加销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var product_id, sku_id, sales_channel_id, price, inventory, available_inventory int32
				sql := "SELECT product_id, sku_id, sales_channel_id, price, inventory, available_inventory FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ?"
				err := common.DB.QueryRow(sql, 100, 1).Scan(&product_id, &sku_id, &sales_channel_id, &price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(product_id, ShouldEqual, 11)
				So(sku_id, ShouldEqual, 1)
				So(sales_channel_id, ShouldEqual, 100)
				So(price, ShouldEqual, 1000)
				So(inventory, ShouldEqual, 0) //sku可用库存不足时，采用共用库存
				So(available_inventory, ShouldEqual, -1)
				//渠道属性测试
				var count uint32
				sql = "SELECT count(1) FROM product_attr WHERE product_id = ? AND attr_id = ? AND attr_value = ? AND status = 0"
				err = common.DB.QueryRow(sql, 11, 27, 100).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
				var count2 uint32
				sql = "SELECT count(1) FROM product_attr WHERE product_id = ? AND attr_id = ? AND attr_value = ? AND status = 0"
				err = common.DB.QueryRow(sql, 13, 27, 100).Scan(&count2)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count2, ShouldEqual, 1)
			})

		})

		Convey("When 批量添加销售渠道商品信息2 接口", func() {
			jsonStr := `{
					"list":[{
						"salesChannelId":100,
					    "skuList":[
					    {
						"skuId":1,
						"price":1000,
						"availableInventory":50
					    }]
					}]

				}`
			requestAPI := "/api/v1/saleschannelproduct/listAdd"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 批量添加销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var product_id, sku_id, sales_channel_id, price, inventory, available_inventory int32
				sql := "SELECT product_id, sku_id, sales_channel_id, price, inventory, available_inventory FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ?"
				err := common.DB.QueryRow(sql, 100, 1).Scan(&product_id, &sku_id, &sales_channel_id, &price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(product_id, ShouldEqual, 11)
				So(sku_id, ShouldEqual, 1)
				So(sales_channel_id, ShouldEqual, 100)
				So(price, ShouldEqual, 1000)
				So(inventory, ShouldEqual, 50)
				So(available_inventory, ShouldEqual, 50)
			})

		})

		Convey("When 批量添加销售渠道商品信息3 接口", func() {
			jsonStr := `{
					"list":[{
						"salesChannelId":100,
					    "skuList":[
					    {
						"skuId":1,
						"price":1000,
						"availableInventory":-1
					    }]
					}]

				}`
			requestAPI := "/api/v1/saleschannelproduct/listAdd"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 批量添加销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var product_id, sku_id, sales_channel_id, price, inventory, available_inventory int32
				sql := "SELECT product_id, sku_id, sales_channel_id, price, inventory, available_inventory FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ?"
				err := common.DB.QueryRow(sql, 100, 1).Scan(&product_id, &sku_id, &sales_channel_id, &price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(product_id, ShouldEqual, 11)
				So(sku_id, ShouldEqual, 1)
				So(sales_channel_id, ShouldEqual, 100)
				So(price, ShouldEqual, 1000)
				So(inventory, ShouldEqual, 0)
				So(available_inventory, ShouldEqual, -1)
			})

		})





		Convey("When  批量删除sku销售渠道信息 接口", func() {
			jsonStr := `{
				    "list":[
				    {
					"salesChannelId":100,
					"skuIdList":[2]
				    }
				    ]
				}`
			requestAPI := "/api/v1/saleschannelproduct/DeleteSalesChannelSkuInfoList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量删除sku销售渠道信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 100, 2).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
				//渠道属性测试
				var count2 uint32
				sql = "SELECT count(1) FROM product_attr WHERE product_id = ? AND attr_id = ? AND attr_value = ? AND status = 0"
				err = common.DB.QueryRow(sql, 12, 27, 100).Scan(&count2)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count2, ShouldEqual, 0)
			})

		})

		Convey("When  批量修改sku销售渠道价格1 接口", func() {
			jsonStr := `{
				    "salesChannelId":100,
				    "type":1,
				    "changePrice":-1,
				    "rate":150,
				    "precision":0,
				    "skuIdList":[2]
				}`
			requestAPI := "/api/v1/saleschannelproduct/UpdateSalesChannelSkuListPrice"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量修改sku销售渠道价格 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var price int
				sql := "SELECT price FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 100, 2).Scan(&price)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(price, ShouldEqual, 99)
			})

		})

		Convey("When  批量修改sku销售渠道价格2 接口", func() {
			jsonStr := `{
				    "salesChannelId":100,
				    "type":2,
				    "changePrice":-1,
				    "rate":150,
				    "precision":0,
				    "skuIdList":[2]
				}`
			requestAPI := "/api/v1/saleschannelproduct/UpdateSalesChannelSkuListPrice"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量修改sku销售渠道价格 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var price int
				sql := "SELECT price FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 100, 2).Scan(&price)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(price, ShouldEqual, 150)
			})

		})

		Convey("When  批量取消sku销售渠道库存，使用共用库存 接口", func() {
			jsonStr := `{
					    "salesChannelId":100,
					    "skuIdList":[2]
					}`
			requestAPI := "/api/v1/saleschannelproduct/UpdateSalesChannelSkuListInventory"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量取消sku销售渠道库存，使用共用库存 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var inventory, available_inventory int
				sql := "SELECT inventory, available_inventory FROM sku_sales_channel_info WHERE sales_channel_id = ? AND sku_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 100, 2).Scan(&inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(inventory, ShouldEqual, 20)
				So(available_inventory, ShouldEqual, -1)
				var sku_inventory, sku_available_inventory int
				sql = "SELECT inventory, available_inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&sku_inventory, &sku_available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(sku_inventory, ShouldEqual, 200)
				So(sku_available_inventory, ShouldEqual, 130)
			})

		})

		Convey("When  多条件查询sku销售渠道信息 接口", func() {
			jsonStr := `{
					    "salesChannelId":100,
					    "productName":"",
					    "categoryId":0,
					    "productId":0,
					    "supplierId":0,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":1,
					    "sortType":1
					}`
			requestAPI := "/api/v1/saleschannelproduct/QuerySalesChannelProductInfoList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  多条件查询sku销售渠道信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["ProductInfoList"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				querySalesChannelProductInfoListResp := &model.QuerySalesChannelProductInfoListResp{}
				json.Unmarshal(josnbyte, &querySalesChannelProductInfoListResp)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].ProductId, ShouldEqual, 12)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].ProductName, ShouldEqual, "sit测试商品2")
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].AvailableInventory, ShouldEqual, 30)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].Inventory, ShouldEqual, 50)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].Price, ShouldEqual, 120)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].OriginAvailableInventory, ShouldEqual, 100)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].OriginInventory, ShouldEqual, 200)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].OriginPrice, ShouldEqual, 100)
				So(querySalesChannelProductInfoListResp.ProductInfoList[0].SkuList[0].TotalAvailableInventory, ShouldEqual, 130)
			})

		})

		Convey("When  多条件查询不在渠道内的sku信息 接口", func() {
			jsonStr := `{
					    "salesChannelId":100,
					    "productName":"",
					    "categoryId":0,
					    "productId":0,
					    "supplierId":0,
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":1,
					    "sortType":1
					}`
			requestAPI := "/api/v1/saleschannelproduct/QueryNotSalesChannelProductInfoList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  多条件查询不在渠道内的sku信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["ProductInfoList"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				querySalesChannelProductInfoListResp := &model.QuerySalesChannelProductInfoListResp{}
				json.Unmarshal(josnbyte, &querySalesChannelProductInfoListResp)
				So(len(querySalesChannelProductInfoListResp.ProductInfoList), ShouldEqual, 2)
			})

		})

		Convey("When  复制渠道商品信息到空渠道 接口", func() {
			jsonStr := `{
					    "fromSalesChannelId":100,
					    "toSalesChannelId":103
					}`
			requestAPI := "/api/v1/saleschannelproduct/CopySalesChannelProductInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  复制渠道商品信息到空渠道 判定", func() {
				So(result.Status, ShouldEqual, 0)
				sql := "SELECT product_id, sku_id, sales_channel_id, price, inventory, available_inventory FROM sku_sales_channel_info WHERE sales_channel_id = ? AND status = 0"
				rows, err := common.DB.Query(sql, 103)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer rows.Close()
				for rows.Next(){
					var product_id, sku_id, sales_channel_id, price, inventory, available_inventory int
					rows.Scan(&product_id, &sku_id, &sales_channel_id, &price, &inventory, &available_inventory)
					So(product_id, ShouldEqual, 12)
					So(sku_id, ShouldEqual, 2)
					So(sales_channel_id, ShouldEqual, 103)
					So(price, ShouldEqual, 120)
					So(inventory, ShouldEqual, 0)
					So(available_inventory, ShouldEqual, -1)
				}
			})

		})


		Convey("When  批量更新销售渠道商品信息1 接口", func() {//分配库存调整
			jsonStr := `{
				    "list":[
				    {
					"salesChannelId":100,
					"skuList":[
					{
					    "skuId":2,
					    "price":1000,
					    "availableInventory":10
					}]
				    }
				    ]
				}`
			requestAPI := "/api/v1/saleschannelproduct/listUpdate"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量更新销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var price, inventory, available_inventory int
				sql := "SELECT price, inventory, available_inventory FROM sku_sales_channel_info WHERE sku_id = ? AND sales_channel_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 100).Scan(&price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(price, ShouldEqual, 1000)
				So(available_inventory, ShouldEqual, 10)
				So(inventory, ShouldEqual, 30)
				var sku_inventory, sku_available_inventory int
				sql = "SELECT inventory, available_inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&sku_inventory, &sku_available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(sku_inventory, ShouldEqual, 200)
				So(sku_available_inventory, ShouldEqual, 120)
			})

		})

		Convey("When  批量更新销售渠道商品信息2 接口", func() {
			jsonStr := `{
				    "list":[
				    {
					"salesChannelId":100,
					"skuList":[
					{
					    "skuId":2,
					    "price":1000,
					    "availableInventory":1000
					}]
				    }
				    ]
				}`
			requestAPI := "/api/v1/saleschannelproduct/listUpdate"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量更新销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 6606005)
				So(result.Msg, ShouldEqual, "设置的skuId（2）在渠道Id（100)中过高，sku库存不足")
			})

		})

		Convey("When  批量更新销售渠道商品信息3 接口", func() {//分配转共用
			jsonStr := `{
				    "list":[
				    {
					"salesChannelId":100,
					"skuList":[
					{
					    "skuId":2,
					    "price":1000,
					    "availableInventory":-1
					}]
				    }
				    ]
				}`
			requestAPI := "/api/v1/saleschannelproduct/listUpdate"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量更新销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var price, inventory, available_inventory int
				sql := "SELECT price, inventory, available_inventory FROM sku_sales_channel_info WHERE sku_id = ? AND sales_channel_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 100).Scan(&price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(price, ShouldEqual, 1000)
				So(available_inventory, ShouldEqual, -1)
				So(inventory, ShouldEqual, 20)
				var sku_inventory, sku_available_inventory int
				sql = "SELECT inventory, available_inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&sku_inventory, &sku_available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(sku_inventory, ShouldEqual, 200)
				So(sku_available_inventory, ShouldEqual, 130)
			})

		})

		Convey("When  批量更新销售渠道商品信息4 接口", func() {//共用转分配
			jsonStr := `{
				    "list":[
				    {
					"salesChannelId":104,
					"skuList":[
					{
					    "skuId":2,
					    "price":1000,
					    "availableInventory":30
					}]
				    }
				    ]
				}`
			requestAPI := "/api/v1/saleschannelproduct/listUpdate"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量更新销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var price, inventory, available_inventory int
				sql := "SELECT price, inventory, available_inventory FROM sku_sales_channel_info WHERE sku_id = ? AND sales_channel_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 104).Scan(&price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(price, ShouldEqual, 1000)
				So(available_inventory, ShouldEqual, 30)
				So(inventory, ShouldEqual, 30)
				var sku_inventory, sku_available_inventory int
				sql = "SELECT inventory, available_inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&sku_inventory, &sku_available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(sku_inventory, ShouldEqual, 200)
				So(sku_available_inventory, ShouldEqual, 70)
			})

		})


		Convey("When  批量更新销售渠道商品信息5 接口", func() {//共用转分配(0边界值测试)
			jsonStr := `{
				    "list":[
				    {
					"salesChannelId":104,
					"skuList":[
					{
					    "skuId":2,
					    "price":1000,
					    "availableInventory":0
					}]
				    }
				    ]
				}`
			requestAPI := "/api/v1/saleschannelproduct/listUpdate"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then  批量更新销售渠道商品信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var price, inventory, available_inventory int
				sql := "SELECT price, inventory, available_inventory FROM sku_sales_channel_info WHERE sku_id = ? AND sales_channel_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 104).Scan(&price, &inventory, &available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(price, ShouldEqual, 1000)
				So(available_inventory, ShouldEqual, 0)
				So(inventory, ShouldEqual, 0)
				var sku_inventory, sku_available_inventory int
				sql = "SELECT inventory, available_inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&sku_inventory, &sku_available_inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(sku_inventory, ShouldEqual, 200)
				So(sku_available_inventory, ShouldEqual, 100)
			})

		})


		Convey("When 修改sku可用库存1 接口", func() {
			jsonStr := `{
					    "salesChannelId":100,
					    "skuId":2,
					    "availableInventory":80
					}`
			requestAPI := "/api/v1/skuExt/updateInventory"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改sku可用库存 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var availableInventory, inventory int
				sql := "SELECT available_inventory, inventory FROM sku_sales_channel_info WHERE sku_id = ? AND sales_channel_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 100).Scan(&availableInventory, &inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(availableInventory, ShouldEqual, 80)
				So(inventory, ShouldEqual, 100)
				var skuAvailableInventory, skuInventory int
				sql = "SELECT available_inventory, inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&skuAvailableInventory, &skuInventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(skuAvailableInventory, ShouldEqual, 100)
				So(skuInventory, ShouldEqual, 250)


			})

		})

		Convey("When 修改sku可用库存2 接口", func() {
			jsonStr := `{
					    "salesChannelId":104,
					    "skuId":2,
					    "availableInventory":80
					}`
			requestAPI := "/api/v1/skuExt/updateInventory"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改sku可用库存 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var availableInventory, inventory int
				sql := "SELECT available_inventory, inventory FROM sku_sales_channel_info WHERE sku_id = ? AND sales_channel_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 104).Scan(&availableInventory, &inventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(availableInventory, ShouldEqual, -1)
				So(inventory, ShouldEqual, 0)
				var skuAvailableInventory, skuInventory int
				sql = "SELECT available_inventory, inventory FROM sku_info WHERE id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&skuAvailableInventory, &skuInventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(skuAvailableInventory, ShouldEqual, 80)
				So(skuInventory, ShouldEqual, 180)


			})

		})

		Convey("When 修改sku可用库存3 接口", func() {
			jsonStr := `{
					    "salesChannelId":0,
					    "skuId":3,
					    "availableInventory":80
					}`
			requestAPI := "/api/v1/skuExt/updateInventory"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改sku可用库存 判定", func() {
				So(result.Status, ShouldEqual, 6606019)
			})

		})

		Convey("When 修改sku可用库存4 接口", func() {
			jsonStr := `{
					    "salesChannelId":0,
					    "skuId":2,
					    "availableInventory":80
					}`
			requestAPI := "/api/v1/skuExt/updateInventory"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改sku可用库存 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var skuAvailableInventory, skuInventory int
				sql := "SELECT available_inventory, inventory FROM sku_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&skuAvailableInventory, &skuInventory)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(skuAvailableInventory, ShouldEqual, 80)
				So(skuInventory, ShouldEqual, 180)
			})

		})

		Convey("When  查询单个商品所有销售渠道信息 接口", func() {
			requestAPI := "/api/v1/saleschannelproduct/QueryProductAllSalesChannelInfo?productId=12"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)
			Convey("Then  查询单个商品所有销售渠道信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryProductAllSalesChannelInfoResp := &model.QueryProductAllSalesChannelInfoResp{}
				json.Unmarshal(josnbyte, &queryProductAllSalesChannelInfoResp)
				for _, productAllSalesChannelInfoList := range queryProductAllSalesChannelInfoResp.List {
					if productAllSalesChannelInfoList.SalesChannelId == 100 {
						So(productAllSalesChannelInfoList.SkuList[0].AvailableInventory, ShouldEqual, 30)
						So(productAllSalesChannelInfoList.SkuList[0].Inventory, ShouldEqual, 50)
						So(productAllSalesChannelInfoList.SkuList[0].TotalAvailableInventory, ShouldEqual, 130)
					} else if productAllSalesChannelInfoList.SalesChannelId == 104 {
						So(productAllSalesChannelInfoList.SkuList[0].AvailableInventory, ShouldEqual, -1)
						So(productAllSalesChannelInfoList.SkuList[0].Inventory, ShouldEqual, 0)
						So(productAllSalesChannelInfoList.SkuList[0].TotalAvailableInventory, ShouldEqual, 130)
					}
				}

			})

		})

		//common.ExecuteSQLScriptFile("./../../config/salesChannel_clean.sql")
	})

}

