package testcase

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
)



func Test_Public(t *testing.T)  {

	Convey("Given 数据库数据已初始化", t, func(){
		common.ExecuteSQLScriptFile("./../../config/public_clean.sql")
		common.ExecuteSQLScriptFile("./../../config/public.sql")

		Convey("When 拉取province列表 接口", func() {
			requestAPI := "/api/v1/1/province"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PUBLIC)

			Convey("Then 拉取province列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.([]interface {})
				So(len(resultData), ShouldEqual, 2)
			})

		})

		Convey("When 拉取province’s city列表 接口", func() {
			requestAPI := "/api/v1/1/province/110000"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PUBLIC)

			Convey("Then 拉取province’s city列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.([]interface {})
				So(len(resultData), ShouldEqual, 1)
				datamap := resultData[0].(map[string]interface{})
				So(datamap["name"], ShouldEqual, "北京市")
			})

		})

		Convey("When 拉取city’s district列表 接口", func() {
			requestAPI := "/api/v1/1/city/130100"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PUBLIC)

			Convey("Then 拉取city’s district列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.([]interface {})
				So(len(resultData), ShouldEqual, 2)
			})

		})

		Convey("When 拉取图片验证码 接口", func() {
			requestAPI := "/api/v1/1/getImgAndKey"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_PUBLIC)

			Convey("Then 拉取图片验证码 判定", func() {
				So(result.Status, ShouldEqual, 0)
				resultData := result.Data.(map[string]interface{})
				So(resultData["imgUrl"], ShouldEqual, "http://captcha.maizuo.com/captcha/code/getImg?key=tw6XqdrKnsh7QubxdPE2FeRDfsTUrtZ")
				So(resultData["imgKey"], ShouldEqual, "tw6XqdrKnsh7QubxdPE2FeRDfsTUrtZ")
			})

		})


		common.ExecuteSQLScriptFile("./../../config/public_clean.sql")
	})

}

