package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"encoding/json"
	"maizuo.com/aura/aura-sit/model"
	"fmt"
)


func Test_Bill(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func() {
		common.ExecuteSQLScriptFile("./../../config/manager-bill_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/manager-bill.sql")
		Convey("When 拉取结算单周期内的订单差异1 接口", func() {
			jsonStr := `{
							"pageSize": 10,
							"pageNumber": 1,
							"sortKey": 1,
							"sortType": 1,
							"beginTime":1480659597,
							"endTime":1545891579,
							"supplierId": 1,
							"orderId":"",
							"serviceId":""
						}`
			requestAPI := "/api/v1/order/queryBillOrderBalanceList"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 拉取结算单周期内的订单差异1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				josnbyte, _ := json.Marshal(resultData)
				queryBillOrderBalanceListResp := &model.QueryBillOrderBalanceListResp{}
				json.Unmarshal(josnbyte, &queryBillOrderBalanceListResp)
				So(queryBillOrderBalanceListResp.Total, ShouldEqual, 2)

			})

		})


		Convey("When 新增结算单周期内的订单差异项1 接口", func() {
			jsonStr := `{
							"balanceName": "差异项名称添加",
							"serviceId": "123456123",
							"supplierId": 1,
							"orderId": "20171130",
							"balanceAmount": 1000,
							"balanceAttachment": "url",
							"operator": "SIT Tester",
							"remark": "差异补入",
							"source":1
						}`
			requestAPI := "/api/v1/order/addBillOrderBalance"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 新增结算单周期内的订单差异项1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var typeName,serviceId,orderId string
				var balanceAmount int
				sql := "SELECT type_name, service_id, order_id, balance_amount FROM order_balance_info WHERE type_name = ? AND status = 0"
				err := common.DB.QueryRow(sql, "差异项名称添加").Scan(&typeName, &serviceId, &orderId, &balanceAmount)
				if err != nil {
					fmt.Println(err)
				}
				So(typeName, ShouldEqual, "差异项名称添加")
				So(serviceId, ShouldEqual, "123456123")
				So(orderId, ShouldEqual, "20171130")
				So(balanceAmount, ShouldEqual, 1000)
			})

		})

		Convey("When 新增结算单周期内的订单差异项2 接口", func() {//订单、供应商异常
			jsonStr := `{
							"balanceName": "差异项名称添加",
							"serviceId": "123456123",
							"supplierId": 2,
							"orderId": "20171130",
							"balanceAmount": 1000,
							"balanceAttachment": "url",
							"operator": "SIT Tester",
							"remark": "差异补入",
							"source":1
						}`
			requestAPI := "/api/v1/order/addBillOrderBalance"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 新增结算单周期内的订单差异项2 判定", func() {
				So(result.Status, ShouldEqual, 6606201)
			})

		})
	})
}
