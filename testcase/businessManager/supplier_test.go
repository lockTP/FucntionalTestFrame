package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"

	"fmt"
)



func Test_Supplier(t *testing.T)  {
	Convey("Given 数据库不为空", t, func() {
		common.ExecuteSQLScriptFile("./../../config/supplier_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/supplier.sql")
		Convey("When 增加供应商 接口", func() {
			jsonStr := `{
							"supplierName":"添加供应商",
							"settlementPeriod": 1,
							"supplierStatus":1
						}`
			requestAPI := "/api/v1/supplier/add"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 查询销售渠道信息结果 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var settlementPeriod,nextSettlementPeriod,supplierStatus uint32
				sql := "SELECT settlement_period,next_settlement_period, supplier_status FROM supplier_info WHERE supplier_name = ? AND status = 0"
				err := common.DB.QueryRow(sql, "添加供应商").Scan(&settlementPeriod, &nextSettlementPeriod, &supplierStatus)
				if err != nil {
					fmt.Println(err)
				}
				So(settlementPeriod, ShouldEqual, 1)
				So(nextSettlementPeriod, ShouldEqual, 1)
				So(supplierStatus, ShouldEqual, 1)
			})

		})

		Convey("When 修改供应商（aura管理系统使用）1 接口", func() {
			jsonStr := `{
							"id":1,
							"supplierName":"修改供应商",
							"nextSettlementPeriod": 2,
							"supplierStatus":1
						}`
			requestAPI := "/api/v1/supplier/set"
			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改供应商（aura管理系统使用）1 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var supplierName string
				var settlementPeriod,supplierStatus,nextSettlementPeriod uint32
				sql := "SELECT supplier_name, settlement_period,next_settlement_period, supplier_status FROM supplier_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 1).Scan(&supplierName, &settlementPeriod,&nextSettlementPeriod, &supplierStatus)
				if err != nil {
					fmt.Println(err)
				}
				So(settlementPeriod, ShouldEqual, 1)
				So(nextSettlementPeriod, ShouldEqual, 2)
				So(supplierStatus, ShouldEqual, 1)
				So(supplierName, ShouldEqual, "修改供应商")

			})

		})

		Convey("When 修改供应商（aura管理系统使用）2 接口", func() {//供应商当前结算周期为0时，同步更新
			jsonStr := `{
							"id":2,
							"supplierName":"修改供应商",
							"nextSettlementPeriod": 2,
							"supplierStatus":1
						}`
			requestAPI := "/api/v1/supplier/set"
			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_MANAGER)

			Convey("Then 修改供应商（aura管理系统使用）2 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var supplierName string
				var settlementPeriod,supplierStatus,nextSettlementPeriod uint32
				sql := "SELECT supplier_name, settlement_period,next_settlement_period, supplier_status FROM supplier_info WHERE id = ? AND status = 0"
				err := common.DB.QueryRow(sql, 2).Scan(&supplierName, &settlementPeriod,&nextSettlementPeriod, &supplierStatus)
				if err != nil {
					fmt.Println(err)
				}
				So(settlementPeriod, ShouldEqual, 2)
				So(nextSettlementPeriod, ShouldEqual, 2)
				So(supplierStatus, ShouldEqual, 1)
				So(supplierName, ShouldEqual, "修改供应商")

			})

		})

		Convey("When 删除供应商 接口", func() {
			requestAPI := "/api/v1/supplier/del/1"
			result, _ := common.AccessAPIWithPostBody("DELETE", requestAPI, "", common.BUSINESS_MANAGER)

			Convey("Then 删除供应商 判定", func() {
				So(result.Status, ShouldEqual, 0)
				var status uint32
				sql := "SELECT status FROM supplier_info WHERE id = ?"
				err := common.DB.QueryRow(sql, 1).Scan(&status)
				if err != nil {
					fmt.Println(err)
				}
				So(status, ShouldEqual, 1)

			})

		})
	})
}
