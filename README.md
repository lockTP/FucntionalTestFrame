# aura-sit

# aura平台系统集成测试

## 准备工作
发布项目置aurapt上
data层需要连接SIT测试数据库


## 代码解析
<code>
func Test_Delivery(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func(){
		common.ExecuteSQLScriptFile("./../../config/delivery_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/delivery.sql")

		Convey("When 发货请求 接口", func() {
			jsonStr := `{
					    "orderId":["201610191952205099"]
					}`
			requestAPI := "/api/v1/1/deliverycenter/delivery"
			result, _ := common.AccessAPI("PUT", requestAPI, jsonStr, common.BUSINESS_DELIVERY)

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
		common.ExecuteSQLScriptFile("./../../config/delivery_clean.sql")
		}

</code>
### 执行顺序
代码执行顺序为：Given -->  When --> Then<br>
Given为执行接口的先决条件，通常为初始化数据库，配置环境等等<br>
When为测试预期结果之前做的调用操作<br>
Then为预期结果测试<br>

### 增删改查
增：通常在调用接口后，期望数据库有新的数据，因此可能会使用query去查询数据库确保数据库有新增的数据<br>
删：通常在调用接口后，期望原有的数据库少了某些数据，我们项目是软删除，也是需要query查询确保数据被删除<br>
改：通常在调用接口后，期望原有的数据被更新，需要query查询确保数据被更改<br>
查：通常在调用接口后，直接可从返回值中判断该接口是否执行正常<br>

## 第三方接口调用
在使用第三方调用时，使用mockee来进行虚拟，可修改配置文件讲发送地址改为下方地址：<br>
http://mockee.miz.so/<br>
内部接口调用方式如：http://mockee.miz.so/aura/delivery<br>

## 预期效果
通过测试类，能够保证在代码进行修改更新后，原有的接口依然能够正常执行与使用。大家在日常工作中遇到了接口出现bug，如某种情况没有考虑到导致返回不是预期返回，请讲此bug添加入测试类，以保证下一次不会出现。
平时开发时可以利用这些测试类来保证修改的代码不会对早期代码或其他人的代码造成影响。