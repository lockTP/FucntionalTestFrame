package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"fmt"
)



func Test_Delivery(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func(){
		common.ExecuteSQLScriptFile("./../../config/delivery_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/delivery.sql")

		Convey("When 发货请求 接口", func() {
			jsonStr := `{
					    "orderId":["201610191952205099"]
					}`
			requestAPI := "/api/v1/1/deliverycenter/delivery"
			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_DELIVERY)

			Convey("Then 发货请求 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var orderStatus int
				sql := "SELECT order_status FROM order_info WHERE order_id = ?"
				err := common.DB.QueryRow(sql, "201610191952205099").Scan(&orderStatus)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(orderStatus, ShouldEqual, 6)
			})

		})

		Convey("When 接收业务系统发货结果 接口", func() {
			jsonStr := `{
					    "orderId":"201610191952205098",
					    "status": 0,
					    "deliveryParams":"abcdefg"
					}`
			requestAPI := "/api/v1/1/deliverycenter/notify"
			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_DELIVERY)

			Convey("Then 接收业务系统发货结果 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var orderStatus, deliveryResultCount int
				sql := "SELECT order_status FROM order_info WHERE order_id = ?"
				err := common.DB.QueryRow(sql, "201610191952205098").Scan(&orderStatus)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(orderStatus, ShouldEqual, 4)
				sql = "SELECT count(id) FROM order_delivery_result_info WHERE order_id = ?"
				err = common.DB.QueryRow(sql, "201610191952205098").Scan(&deliveryResultCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(deliveryResultCount, ShouldEqual, 1)
			})

		})

		Convey("When 查询业务系统发货结果 接口", func() {
			jsonStr := `{
					    "orderId":"201610191952205098"
					}`
			requestAPI := "/api/v1/1/deliverycenter/query"
			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_DELIVERY)

			Convey("Then 查询业务系统发货结果 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var orderStatus, deliveryResultCount int
				sql := "SELECT order_status FROM order_info WHERE order_id = ?"
				err := common.DB.QueryRow(sql, "201610191952205098").Scan(&orderStatus)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(orderStatus, ShouldEqual, 4)
				sql = "SELECT count(id) FROM order_delivery_result_info WHERE order_id = ?"
				err = common.DB.QueryRow(sql, "201610191952205098").Scan(&deliveryResultCount)
				if err != nil {
					fmt.Println(err.Error())
				}
				So(deliveryResultCount, ShouldEqual, 1)
			})

		})



		common.ExecuteSQLScriptFile("./../../config/delivery_clean.sql")
	})

}

