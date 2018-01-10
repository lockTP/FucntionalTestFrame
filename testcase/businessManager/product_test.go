package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"

	"maizuo.com/aura/aura-sit/model"
	"encoding/json"
	"fmt"
)



func Test_Product(t *testing.T)  {
	Convey("Given 数据库为空", t, func() {
		common.ExecuteSQLScriptFile("./../../config/product_clean.sql")

		Convey("When 根据后端分类id获取商品列表接口 接口", func() {
			req := &model.ProductListFormngReq{
				ProductName: "",
				CategoryID: []uint32{},
				SupplierIDs: []uint32{},
				ProductSource: 9999,
				ProductStatus: 9999,
				ProductID: 0,
				SkuCode: "",
				SkuId: 0,
				Brand: "",
				PageSize:   10,
				PageNumber: 1,
				SortKey:    1,
				SortType:   1,
			}
			jsonbyte, _ := json.Marshal(req)
			jsonStr := string(jsonbyte)
			requestAPI := "/api/v1/product/list"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 根据后端分类id获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual, 6003)
				So(result.Data, ShouldBeNil)
			})

		})

		Convey("When 增加选项接口 接口", func() {
			jsonStr := `{
					    "optionName":"颜色",
					    "optionValueList":["红","蓝"]
					}`
			requestAPI := "/api/v1/option/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 增加选项接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var optionCount, optionValueCount int
				sql := "SELECT count(id) FROM option_info WHERE option_name = ?"
				err := common.DB.QueryRow(sql, "颜色").Scan(&optionCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				sql2 := "SELECT count(id) FROM option_value_info"
				err = common.DB.QueryRow(sql2).Scan(&optionValueCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionCount, ShouldEqual, 1)
				So(optionValueCount, ShouldEqual, 2)
			})

		})


	})

	Convey("Given 数据库数据已初始化", t, func(){
		common.ExecuteSQLScriptFile("./../../config/product_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/product.sql")

		Convey("When 根据后端分类id获取商品列表接口 接口", func() {
			req := &model.ProductListFormngReq{
				ProductName: "",
				CategoryID: []uint32{},
				SupplierIDs: []uint32{},
				ProductSource: 9999,
				ProductStatus: 9999,
				ProductID: 0,
				SkuCode: "",
				SkuId: 0,
				Brand: "",
				PageSize:   10,
				PageNumber: 1,
				SortKey:    1,
				SortType:   1,
			}
			jsonbyte, _ := json.Marshal(req)
			jsonStr := string(jsonbyte)
			requestAPI := "/api/v1/product/list"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 根据后端分类id获取商品列表接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["total"], ShouldEqual, 3)
				josnbyte, _ := json.Marshal(resultData)
				queryProductInfoList := &model.QueryProductInfoList{}
				json.Unmarshal(josnbyte, &queryProductInfoList)
				flagContain := false  //用于确保有期望的商品1002
				for _, v :=  range queryProductInfoList.List{
					if v.ProductId == 1002 {
						flagContain = true
						So(v.MasterName, ShouldEqual, "sit测试商品上架")
						So(v.CategoryName, ShouldEqual, "category")
						So(v.ParentCategoryName, ShouldEqual, "parent category")
					}
				}
				So(flagContain, ShouldBeTrue)
			})

		})

		Convey("When 拉取选项列表 接口", func() {

			requestAPI := "/api/v1/option/list?pageSize=10&pageNumber=1&searchKey=&sortKey=1&sortType=1"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)

			Convey("Then 拉取选项列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["total"], ShouldEqual, 3)
			})

		})

		Convey("When 拉取选项的值列表 接口", func() {

			requestAPI := "/api/v1/option/value/list?optionId=1001&pageSize=10&pageNumber=1&sortKey=1&sortType=1"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)

			Convey("Then 拉取选项的值列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["optionValueList"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryOptionValueList := &model.QueryOptionValueList{}
				json.Unmarshal(josnbyte, &queryOptionValueList)
				So(resultData["total"], ShouldEqual, 4)
				So(queryOptionValueList.OptionId, ShouldEqual, 1001)
				So(queryOptionValueList.OptionName, ShouldEqual, "选项1")
				So(len(queryOptionValueList.OptionValueList), ShouldEqual, 4)
			})

		})


		Convey("When 修改选项 接口", func() {

			jsonStr := `{
					    "optionId":1001,
					    "optionName":"修改选项值"
					}`
			requestAPI := "/api/v1/option/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改选项 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var optionName string
				sql := "SELECT option_name FROM option_info WHERE id = ?"
				err := common.DB.QueryRow(sql, 1001).Scan(&optionName)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionName, ShouldEqual, "修改选项值")
			})

		})

		Convey("When 删除选项 接口1", func() {

			jsonStr := `{
					    "optionId":1001
					}`
			requestAPI := "/api/v1/option/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除选项 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var optionCount, optionValueCount int
				sql := "SELECT count(id) FROM option_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1001).Scan(&optionCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionCount, ShouldEqual, 0)
				sql2 := "SELECT count(id) FROM option_value_info WHERE option_id = ? AND status = 0"
				err = common.DB.QueryRow(sql2, 1001).Scan(&optionValueCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionValueCount, ShouldEqual, 0)
			})

		})

		Convey("When 删除选项 接口2", func() {//关联sku的option不允许删除

			jsonStr := `{
					    "optionId":1002
					}`
			requestAPI := "/api/v1/option/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除选项 判定", func() {
				So(result.Status, ShouldEqual, 6606000)
				var optionCount, optionValueCount int
				sql := "SELECT count(id) FROM option_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1002).Scan(&optionCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionCount, ShouldEqual, 1)
				sql2 := "SELECT count(id) FROM option_value_info WHERE option_id = ? AND status = 0"
				err = common.DB.QueryRow(sql2, 1002).Scan(&optionValueCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionValueCount, ShouldEqual, 2)
			})

		})

		Convey("When 新增选项值 接口", func() {

			jsonStr := `{
					    "optionId":1001,
					    "optionValue":"选项值15"
					}`
			requestAPI := "/api/v1/option/value/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 新增选项值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM option_value_info WHERE option_id = ?"
				err := common.DB.QueryRow(sql, 1001).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 5)
			})

		})


		Convey("When 修改选项值 接口", func() {

			jsonStr := `{
					    "optionValueId":1,
					    "optionId":1001,
					    "optionValue":"修改选项值"
					}`
			requestAPI := "/api/v1/option/value/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改选项值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var optionValue string
				sql := "SELECT option_value FROM option_value_info WHERE option_id = ?"
				err := common.DB.QueryRow(sql, 1001).Scan(&optionValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionValue, ShouldEqual, "修改选项值")
			})

		})


		Convey("When 删除选项值 接口", func() {

			jsonStr := `{
					    "optionValueId":1
					}`
			requestAPI := "/api/v1/option/value/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除选项值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM option_value_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})

		Convey("When 拉取品类选项列表 接口", func() {

			requestAPI := "/api/v1/category/option/list?categoryId=2&pageSize=10&pageNumber=1&sortKey=1&sortType=1"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)

			Convey("Then 拉取品类选项列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryCategoryOptionList := &model.QueryCategoryOptionList{}
				json.Unmarshal(josnbyte, &queryCategoryOptionList)
				So(queryCategoryOptionList.Total, ShouldEqual, 1)
				So(len(queryCategoryOptionList.List), ShouldEqual, 1)
				So(queryCategoryOptionList.List[0].OptionName, ShouldEqual, "选项1")
				So(len(queryCategoryOptionList.List[0].OptionValueList), ShouldEqual, 2)
			})

		})


		Convey("When 增加品类选项 接口", func() {
			jsonStr := `{
					    "categoryId":2,
					    "optionId":1002,
					    "optionName":"选项2",
					    "optionValueList":["选项值21","选项值22"]
					}`
			requestAPI := "/api/v1/category/option/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 增加品类选项 判定", func() {
				So(result.Status, ShouldEqual, 0)
				sql := "SELECT option_value FROM category_option_value_info WHERE option_id = ? AND category_id = ?"
				rows, err := common.DB.Query(sql,1002, 2)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer rows.Close()
				optionValueList := []string{}
				for rows.Next() {
					var optionValue string
					rows.Scan(&optionValue)
					optionValueList = append(optionValueList, optionValue)
				}
				So(len(optionValueList), ShouldEqual, 2)
				var optionName string
				sql2 := "SELECT option_name FROM category_option_info WHERE option_id = ? AND category_id = ?"
				err = common.DB.QueryRow(sql2,1002, 2).Scan(&optionName)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionName, ShouldEqual, "选项2")

			})

		})


		Convey("When 修改品类选项 接口", func() {
			jsonStr := `{
					    "categoryId":2,
					    "optionId":1002,
					    "optionName":"选项3",
					    "optionValueList":["选项值31","选项值32"]
					}`
			requestAPI := "/api/v1/category/option/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改品类选项 判定", func() {
				So(result.Status, ShouldEqual, 0)
				sql := "SELECT option_value FROM category_option_value_info WHERE option_id = ? AND category_id = ?"
				rows, err := common.DB.Query(sql,1002, 2)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer rows.Close()
				optionValueList := []string{}
				for rows.Next() {
					var optionValue string
					rows.Scan(&optionValue)
					optionValueList = append(optionValueList, optionValue)
				}
				So(len(optionValueList), ShouldEqual, 2)
				var optionName string
				sql2 := "SELECT option_name FROM category_option_info WHERE option_id = ? AND category_id = ?"
				err = common.DB.QueryRow(sql2,1002, 2).Scan(&optionName)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(optionName, ShouldEqual, "选项3")

			})

		})


		Convey("When 删除品类选项 接口", func() {
			jsonStr := `{
					    "categoryId":2,
					    "optionId":1001
					}`
			requestAPI := "/api/v1/category/option/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除品类选项 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count1, count2 int
				sql := "SELECT count(id) FROM category_option_value_info WHERE option_id = ? AND category_id = ?"
				err := common.DB.QueryRow(sql,1002, 2).Scan(&count1)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count1, ShouldEqual, 0)
				sql2 := "SELECT count(id) FROM category_option_info WHERE option_id = ? AND category_id = ?"
				err = common.DB.QueryRow(sql2,1002, 2).Scan(&count2)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count2, ShouldEqual, 0)

			})

		})


		Convey("When 拉取属性列表1 接口", func() {
			jsonStr := `{
					    "searchKey": "",
					    "attrRange":[],
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"updated_at",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/attr/list"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 拉取属性列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryAttrList := &model.QueryAttrList{}
				json.Unmarshal(josnbyte, &queryAttrList)
				So(queryAttrList.Total, ShouldEqual, 5)
				So(len(queryAttrList.List), ShouldEqual, 5)
				for _, v := range queryAttrList.List {
					if v.AttrId == 1001 {
						So(v.AttrValueList[0].AttrValueId, ShouldEqual, 1)
						So(v.AttrValueList[0].AttrValue, ShouldEqual, "这是一个属性值")
						So(v.AttrValueList[0].AttrValueDisplay, ShouldEqual, "这是一个属性值(显示)")
					}
				}

			})

		})


		Convey("When 拉取属性列表2 接口", func() {
			jsonStr := `{
					    "searchKey": "",
					    "attrRange":[1],
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"updated_at",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/attr/list"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 拉取属性列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["list"], ShouldNotBeNil)
				josnbyte, _ := json.Marshal(resultData)
				queryAttrList := &model.QueryAttrList{}
				json.Unmarshal(josnbyte, &queryAttrList)
				fmt.Println(queryAttrList.List)
				So(queryAttrList.Total, ShouldEqual, 2)
				So(len(queryAttrList.List), ShouldEqual, 2)
			})

		})


		Convey("When 拉取属性的值列表 接口", func() {
			requestAPI := "/api/v1/attr/value/list?attrId=1001&pageSize=10&pageNumber=1&sortKey=1&sortType=1"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)

			Convey("Then 拉取属性的值列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				attrList := &model.AttrList{}
				json.Unmarshal(josnbyte, &attrList)
				So(attrList.Total, ShouldEqual, 1)
				So(attrList.AttrId, ShouldEqual, 1001)
				So(attrList.AttrName, ShouldEqual, "属性1")
				So(len(attrList.AttrType), ShouldEqual, 1)
				So(attrList.AttrValueList[0].AttrValue, ShouldEqual,"这是一个属性值")
				So(attrList.AttrValueList[0].AttrValueDisplay, ShouldEqual,"这是一个属性值(显示)")
			})

		})

		Convey("When 增加属性 接口", func() {
			jsonStr := `{
					    "attrName":"颜色",
					    "attrRange":0,
					    "attrStyle": 1,
					    "remark":"说明",
					    "attrType": [0,1],
					    "attrValueList":[{
						"attrValue":"red",
						"attrValueDisplay":"红"},{
						"attrValue":"blue",
						"attrValueDisplay":"蓝"
						}]
					}`
			requestAPI := "/api/v1/attr/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 增加属性 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count1, count2 int
				sql := "SELECT count(id) FROM attr_info WHERE attr_range = 0 AND status = 0"
				err := common.DB.QueryRow(sql).Scan(&count1)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count1, ShouldEqual, 3)
				var attrRange, attrStyle int
				var remark string
				sql = "SELECT remark, attr_range, attr_style FROM attr_info WHERE attr_name = '颜色'"
				err = common.DB.QueryRow(sql).Scan(&remark, &attrRange, &attrStyle)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(remark, ShouldEqual, "说明")
				So(attrRange, ShouldEqual, 0)
				So(attrStyle, ShouldEqual, 1)

				sql = "SELECT count(id) FROM attr_value_info WHERE attr_id in (SELECT id FROM attr_info WHERE attr_name = '颜色')"
				err = common.DB.QueryRow(sql).Scan(&count2)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count2, ShouldEqual, 2)
			})

		})

		Convey("When 修改属性1 接口", func() {
			jsonStr := `{
					    "attrId":1001,
					    "attrStyle": 1,
					    "attrName":"尺码",
					    "remark":"说明"
					}`
			requestAPI := "/api/v1/attr/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改属性 判定", func() {
				So(result.Status, ShouldEqual, 6606009)
			})

		})

		Convey("When 修改属性2 接口", func() {
			jsonStr := `{
					    "attrId":1001,
					    "attrStyle": 0,
					    "attrName":"尺码",
					    "remark":"说明"
					}`
			requestAPI := "/api/v1/attr/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改属性 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var attrName string
				sql := "SELECT attr_name FROM attr_info WHERE id = ?"
				err := common.DB.QueryRow(sql, 1001).Scan(&attrName)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrName, ShouldEqual, "尺码")
			})

		})

		Convey("When 删除属性1 接口", func() {
			jsonStr := `{
					    "attrId":1001
					}`
			requestAPI := "/api/v1/attr/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除属性 判定", func() {
				So(result.Status, ShouldEqual, 6606009)
			})

		})

		Convey("When 删除属性2 接口", func() {
			jsonStr := `{
					    "attrId":1002
					}`
			requestAPI := "/api/v1/attr/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除属性 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM attr_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1002).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})


		Convey("When 新增属性值 接口", func() {
			jsonStr := `{
					    "attrId":1001,
					    "attrValue":"这是第二个属性值",
					    "attrDisplay":"这是第二个属性值(显示)"
					}`
			requestAPI := "/api/v1/attr/value/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 新增属性值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM attr_value_info WHERE attr_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1001).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 2)
			})

		})

		Convey("When 修改属性值 接口1", func() {
			jsonStr := `{
					    "attrValueId":1,
					    "attrId":1001,
					    "attrValue":"这是修改的属性值"
					}`
			requestAPI := "/api/v1/attr/value/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改属性值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var attrValue string
				sql := "SELECT attr_value FROM attr_value_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&attrValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrValue, ShouldEqual, "这是修改的属性值")
			})

		})

		Convey("When 修改属性值 接口2", func() {
			jsonStr := `{
					    "attrValueId":1,
					    "attrId":1001,
					    "attrValue":"这是修改的属性值",
						"attrValueDisplay":"这是修改的显示值"
					}`
			requestAPI := "/api/v1/attr/value/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改属性值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var attrValue, attrValueDisplay string
				sql := "SELECT attr_value,attr_value_display FROM attr_value_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&attrValue,&attrValueDisplay)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrValue, ShouldEqual, "这是修改的属性值")
				So(attrValueDisplay, ShouldEqual, "这是修改的显示值")
				var categoryAttrValue, categoryAttrValueDisplay string
				sql = "SELECT attr_value,attr_value_display FROM category_attr_value_info WHERE attr_id = ? AND attr_value = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1001, "这是修改的属性值").Scan(&categoryAttrValue,&categoryAttrValueDisplay)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(categoryAttrValue, ShouldEqual, "这是修改的属性值")
				So(categoryAttrValueDisplay, ShouldEqual, "这是修改的显示值")
			})

		})

		Convey("When 删除属性值1 接口", func() {
			jsonStr := `{
					    "attrValueId":1
					}`
			requestAPI := "/api/v1/attr/value/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除属性值 判定", func() {
				So(result.Status, ShouldEqual, 6606009)
			})

		})

		Convey("When 删除属性值2 接口", func() {
			jsonStr := `{
					    "attrValueId":2
					}`
			requestAPI := "/api/v1/attr/value/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 删除属性值 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM attr_value_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})

		Convey("When 拉取属性类型列表 接口", func() {
			requestAPI := "/api/v1/attrType/list?pageSize=10&pageNumber=1&sortKey=1&sortType=1"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)
			Convey("Then 拉取属性类型列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				queryAttrTypeList := &model.QueryAttrTypeList{}
				json.Unmarshal(josnbyte, &queryAttrTypeList)
				So(queryAttrTypeList.Total, ShouldEqual, 3)
			})

		})

		Convey("When 增加属性类型 接口", func() {
			jsonStr := `{
					    "attrTypeName":"搜索",
					    "remark": "说明"
					}`
			requestAPI := "/api/v1/attrType/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 增加属性类型 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM attr_type WHERE status = 0"
				err := common.DB.QueryRow(sql).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 4)
			})

		})

		Convey("When 修改属性类型1 接口", func() {
			jsonStr := `{
					    "attrTypeId":1,
					    "attrTypeName":"搜索功能",
					    "remark":"abc"
					}`
			requestAPI := "/api/v1/attrType/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 修改属性类型 判定", func() {
				So(result.Status, ShouldEqual, 6606008)
			})

		})

		Convey("When 修改属性类型2 接口", func() {
			jsonStr := `{
					    "attrTypeId":1001,
					    "attrTypeName":"搜索功能",
					    "remark":"abc"
					}`
			requestAPI := "/api/v1/attrType/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 修改属性类型 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var name, remark string
				sql := "SELECT name, remark FROM attr_type WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1001).Scan(&name, &remark)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(name, ShouldEqual, "搜索功能")
				So(remark, ShouldEqual, "abc")
			})

		})

		Convey("When 删除属性类型1 接口", func() {
			jsonStr := `{
					    "attrTypeId":1001
					}`
			requestAPI := "/api/v1/attrType/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 删除属性类型 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM attr_type WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1001).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})

		Convey("When 删除属性类型2 接口", func() {
			jsonStr := `{
					    "attrTypeId":1
					}`
			requestAPI := "/api/v1/attrType/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 删除属性类型 判定", func() {
				So(result.Status, ShouldEqual, 6606008)
				var count int
				sql := "SELECT count(id) FROM attr_type WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
			})

		})

		//添加共有属性至类目测试
		Convey("When  增加品类属性接口 接口", func() {
			jsonStr := `{
						"categoryId":2,
						"attrId":1099,
						"attrName":"属性99",
						"attrMandatory": 0,
						"attrPublic":1,
						"attrValueList":[{
									"attrValue":"black",
									"attrValueDisplay":"黑"},
									{
									"attrValue":"blue",
									"attrValueDisplay":"蓝"
						}]
				}`
			requestAPI := "/api/v1/category/attr/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 增加品类属性接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM category_attr_info_new WHERE category_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 4)
				//商品与属性类型关联关系绑定验证
				var count1 int
				sql = "SELECT count(id) FROM product_attr WHERE product_id = ? AND attr_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1002, 1099).Scan(&count1)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count1, ShouldEqual, 2)
			})

		})

		//更新共有属性至类目测试
		Convey("When  更新品类属性接口 接口", func() {
			jsonStr := `{
						"categoryId":2,
						"attrId":1001,
						"attrName":"属性1",
						"attrMandatory": 0,
						"attrPublic":1,
						"attrValueList":[{
									"attrValue":"black",
									"attrValueDisplay":"黑"},
									{
									"attrValue":"blue",
									"attrValueDisplay":"蓝"
						}]
}`
			requestAPI := "/api/v1/category/attr/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 更新品类属性接口 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM category_attr_info_new WHERE category_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 3)
				//商品与属性类型关联关系绑定验证
				var count1 int
				sql = "SELECT count(id) FROM product_attr WHERE product_id = ? AND attr_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1002, 1001).Scan(&count1)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count1, ShouldEqual, 2)
			})

		})


		//删除品类属性，属性未绑定商品，可删除
		Convey("When  删除品类属性1 接口", func() {
			jsonStr := `{
							"categoryId":2,
							"attrId":1002
						}`
			requestAPI := "/api/v1/category/attr/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 删除品类属性 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM category_attr_info_new WHERE category_id = ? AND attr_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 1002).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})

		//删除品类属性，属性已绑定商品，不能删除
		Convey("When  删除品类属性2 接口", func() {
			jsonStr := `{
							"categoryId":2,
							"attrId":1001
						}`
			requestAPI := "/api/v1/category/attr/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 删除品类属性 判定", func() {
				So(result.Status, ShouldEqual, 1)
				var count int
				sql := "SELECT count(id) FROM category_attr_info_new WHERE category_id = ? AND attr_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2, 1001).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
			})

		})

		Convey("When 增加商品属性 接口", func() {
			jsonStr := `{
					    "productId":1002,
					    "attrRange":1,
					    "attrList": [
						{
						    "attrId":1002,
						    "attrValue":"这是一个要删除的属性值"
						}]
					}`
			requestAPI := "/api/v1/product/attr/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 增加商品属性 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM product_attr WHERE product_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1002).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 7)
				//商品与属性类型关联关系绑定验证
				var count1 int
				sql = "SELECT count(id) FROM product_attr_type WHERE product_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1002).Scan(&count1)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count1, ShouldEqual, 3)
			})

		})

		Convey("When 修改商品属性 接口", func() {
			jsonStr := `{
					    "productId":1002,
					    "attrRange":1,
					    "attrList": [
						{
						    "attrId":1001,
						    "attrValue":"这是一个要删除的属性值"
						}]
					}`
			requestAPI := "/api/v1/product/attr/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 修改商品属性 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var attrValue string
				sql := "SELECT attr_value FROM product_attr WHERE product_id = ? AND attr_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1002, 1001).Scan(&attrValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				//只删除更新attrRange为1的属性，其他的不动（不删除）
				sql = "SELECT attr_value FROM product_attr WHERE product_id = ? AND attr_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 1002, 1).Scan(&attrValue)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(attrValue, ShouldEqual, "1002商品描述")
			})

		})



		Convey("When 拉取供应商列表 接口", func() {
			jsonStr := `{
					    "pageSize":10,
					    "pageNumber":1,
					    "sortKey":"updated_at",
					    "sortType":"desc"
					}`
			requestAPI := "/api/v1/supplier/list"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 拉取供应商列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})


		Convey("When 增加product 接口", func() {
			jsonStr := `{
					    "masterName":"充电宝",
					    "slaveName":"机器猫充电宝, 满100减20",
					    "supplierId":1,
					    "categoryId": 2,
					    "productSource":1,
					    "isVirtualProduct": 1,
					    "displaySalesCount":1000,
					    "deliveryType":0,
						"productType":1,
						"sortNumber":120
					}`
			requestAPI := "/api/v1/product/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 增加product 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var id, supplierId, productSource, categoryId, sortNumber, productType int
				sql := "SELECT id, supplier_id, product_source, sort_number, product_type FROM product_info WHERE master_name = ? AND status = 0"
				err := common.DB.QueryRow(sql, "充电宝").Scan(&id, &supplierId, &productSource, &sortNumber, &productType)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(supplierId, ShouldEqual, 1)
				So(productSource, ShouldEqual, 1)
				So(sortNumber, ShouldEqual, 120)
				So(productType, ShouldEqual, 1)
				sql = "SELECT category_id FROM product_category WHERE product_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, id).Scan(&categoryId)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(categoryId, ShouldEqual, 2)
				//商品自动继承该类目下的隐性属性值
				var count int
				sql = "SELECT count(1) FROM product_attr WHERE product_id = ? AND attr_id = ? AND attr_value = ? AND status = 0"
				err = common.DB.QueryRow(sql, id, 22, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 1)
			})

		})


		Convey("When 修改product 接口", func() {
			jsonStr := `{
					    "productId":1003,
					    "masterName":"充电宝",
					    "slaveName":"机器猫充电宝, 满100减20",
					    "supplierId":1,
					    "categoryId": 3,
					    "productSource":1,
					    "isVirtualProduct":1,
					    "displaySalesCount":999,
					    "deliveryType":0
					}`
			requestAPI := "/api/v1/product/set"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 修改product 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var id, supplierId, productSource, displaySalesCount int
				var masterName string
				sql := "SELECT id, master_name, supplier_id, product_source, display_sales_count FROM product_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1003).Scan(&id, &masterName, &supplierId, &productSource, &displaySalesCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(supplierId, ShouldEqual, 1)
				So(masterName, ShouldEqual, "充电宝")
				So(productSource, ShouldEqual, 1)
				So(displaySalesCount, ShouldEqual, 999)
			})

		})


		Convey("When 删除product 接口", func() {
			jsonStr := `{
					    "productId":1003
					}`
			requestAPI := "/api/v1/product/del"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 删除product 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(id) FROM product_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1003).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 0)
			})

		})



		Convey("When 根据商品ID和属性范围ID拉取属性1 接口", func() {
			jsonStr := `{
					    "productId": 1002,
					    "attrRange": [1]
					}`
			requestAPI := "/api/v1/product/attrRange"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 根据商品ID和属性范围ID拉取属性1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				queryAttrsByPIDAdnAttrRange := &model.QueryAttrsByPIDAdnAttrRange{}
				json.Unmarshal(josnbyte, &queryAttrsByPIDAdnAttrRange)
				So(len(queryAttrsByPIDAdnAttrRange.AttrValues), ShouldEqual, 1)
			})

		})

		Convey("When 根据商品ID和属性范围ID拉取属性2 接口", func() { //剔除属性为1和2的（hardcode去除商品须知和商品描述）
			jsonStr := `{
					    "productId": 1002,
					    "attrRange": [0]
					}`
			requestAPI := "/api/v1/product/attrRange"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 根据商品ID和属性范围ID拉取属性2 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				queryAttrsByPIDAdnAttrRange := &model.QueryAttrsByPIDAdnAttrRange{}
				json.Unmarshal(josnbyte, &queryAttrsByPIDAdnAttrRange)
				So(len(queryAttrsByPIDAdnAttrRange.AttrValues), ShouldEqual, 0)
			})

		})

		Convey("When 根据主商品sku ID拉取套餐商品sku1 接口", func() {
			jsonStr := `{
							"skuIdList":[1]
						}`
			requestAPI := "/api/v1/product/QueryProductSetInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 根据主商品sku ID拉取套餐商品sku1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				masterSkuList := &model.MasterSkuList{}
				json.Unmarshal(josnbyte, &masterSkuList)
				So(len(masterSkuList.MasterSkuList), ShouldEqual, 1)
				So(len(masterSkuList.MasterSkuList[0].SkuCounts), ShouldEqual, 1)
				So(masterSkuList.MasterSkuList[0].SkuCounts[0].Count, ShouldEqual, 10)
				//子sku信息
				So(masterSkuList.MasterSkuList[0].SkuCounts[0].ProductId, ShouldEqual, 1002)
				So(masterSkuList.MasterSkuList[0].SkuCounts[0].MasterName, ShouldEqual, "sit测试商品上架")
				So(masterSkuList.MasterSkuList[0].SkuCounts[0].SkuMasterName, ShouldEqual, "master_name2")
			})

		})

		Convey("When 根据主商品sku ID拉取套餐商品sku2 接口", func() {
			jsonStr := `{
							"skuIdList":[1,2]
						}`
			requestAPI := "/api/v1/product/QueryProductSetInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 根据主商品sku ID拉取套餐商品sku2 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				masterSkuList := &model.MasterSkuList{}
				json.Unmarshal(josnbyte, &masterSkuList)
				So(len(masterSkuList.MasterSkuList), ShouldEqual, 2)
				So(len(masterSkuList.MasterSkuList[0].SkuCounts) + len(masterSkuList.MasterSkuList[1].SkuCounts), ShouldEqual, 1)

			})

		})

		Convey("When 新增更新套餐礼包商品关联sku 接口", func() {
			jsonStr := `{
						"masterSkuId": 1,
						"skuList": [
							{
								"skuId": 2,
								"skuCount": 15
							},
							{
								"skuId": 3,
								"skuCount": 8
							}
						]
					}`
			requestAPI := "/api/v1/product/updateProductSetInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 新增更新套餐礼包商品关联sku 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var count int
				sql := "SELECT count(1) FROM product_set_info WHERE master_sku_id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&count)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(count, ShouldEqual, 2)
				var skuCount uint32
				sql = "SELECT sku_count FROM product_set_info WHERE sku_id = ? AND status = 0"
				err = common.DB.QueryRow(sql, 2).Scan(&skuCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(skuCount, ShouldEqual, 15)
			})

		})


		Convey("When 新增更新套餐礼包商品关联sku1 接口", func() {//供应商不同业务异常测试
			jsonStr := `{
						"masterSkuId": 1,
						"skuList": [
							{
								"skuId": 2,
								"skuCount": 15
							},
							{
								"skuId": 4,
								"skuCount": 3
							}
						]
					}`
			requestAPI := "/api/v1/product/updateProductSetInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 新增更新套餐礼包商品关联sku1 判定", func() {
				So(result.Status, ShouldEqual, 6606033)
				So(result.Msg, ShouldEqual, "4（skuId）的供应商和发货方式与主sku不相同。")
			})

		})

		Convey("When 新增更新套餐礼包商品关联sku2 接口", func() {//skuId不存在测试
			jsonStr := `{
						"masterSkuId": 1,
						"skuList": [
							{
								"skuId": 222222,
								"skuCount": 15
							},
							{
								"skuId": 3,
								"skuCount": 3
							}
						]
					}`
			requestAPI := "/api/v1/product/updateProductSetInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 新增更新套餐礼包商品关联sku2 判定", func() {
				So(result.Status, ShouldEqual, 6606033)
				So(result.Msg, ShouldEqual, "222222（skuId）不存在。")
			})

		})


		Convey("When 根据skuid 列表拉取商品基本信息 接口", func() {//供应商不同业务异常测试
			jsonStr := `{
							"skuIdList":[1]
						}`
			requestAPI := "/api/v1/product/QueryProductInfobySkuidlist"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 根据skuid 列表拉取商品基本信息 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				data := &model.QueryProductInfobySkuidlistResp{}
				json.Unmarshal(josnbyte, &data)
				So(data.SkuList[0].MasterName, ShouldEqual, "sit测试商品")
				So(data.SkuList[0].SlaveName, ShouldEqual, "sit测试商品附属")
				So(data.SkuList[0].SkuMasterName, ShouldEqual, "master_name")
				So(data.SkuList[0].SkuSlaveName, ShouldEqual, "slave_name")
				So(data.SkuList[0].Image, ShouldEqual, "https://pic.maizuo.com/manager/6cc3be520715bc134b21c9e808677569.jpg")
				So(data.SkuList[0].SkuCode, ShouldEqual, "123456")
			})

		})

		Convey("When 设置商品上下架1 接口", func() {//礼包商品上下架测试(子商品下架，则礼包商品一并下架)
			jsonStr := `{
							"productId":1002,
							"productStatus": 0
						}`
			requestAPI := "/api/v1/product/status"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)
			Convey("Then 设置商品上下架1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				sql := "SELECT product_status FROM product_info WHERE id IN (1001,1002) AND status = 0"
				rows, err := common.DB.Query(sql)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer rows.Close()
				for rows.Next() {
					var productStatus uint32
					rows.Scan(&productStatus)
					So(productStatus, ShouldEqual, 0)
				}
			})
		})

		Convey("When 获取商品详细信息1 接口", func() {//管理系统礼包商品库存不使用礼包计算库存

			requestAPI := "/api/v1/product?productId=1001"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_MANAGER)
			Convey("Then 获取商品详细信息1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				josnbyte, _ := json.Marshal(result.Data)
				data := &model.ProductDetailInfo{}
				json.Unmarshal(josnbyte, &data)
				So(data.MasterName, ShouldEqual, "sit测试商品")
				So(data.ProductSkuList[0].AvailableInventory, ShouldEqual, 100)
			})
		})
		//common.ExecuteSQLScriptFile("./../../config/product_clean.sql")
	})

}

