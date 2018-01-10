package businessUser

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"maizuo.com/aura/aura-sit/common"
	"maizuo.com/aura/aura-sit/model"
	"encoding/json"
)

func Test_User(t *testing.T) {

	Convey("Given 数据库数据已初始化", t, func(){
		//common.ExecuteSQLScriptFile("./../../config/user_clean.sql")
		//common.ExecuteSQLScriptFile("./../../config/user.sql")

		/*************收货地址相关接口 **************/
		Convey("When 1.增加用户收货地址 接口", func() {
			req := &model.AddUserAddress{
				UserId:1001,
				Name:"vicen",
				CountryId:86,
				ProvinceId:440000,
				CityId:440300,
				DistrictId:440305,
				Address:"科苑路科兴科学园B3栋16楼1",
				PostCode:"518000",
				Phone:"15018521673",
				IsDefault:0,
				IdCard:"4414231990000000",
			}
			jsonByte, _ := json.Marshal(req)
			jsonStr := string(jsonByte)
			requestAPI := "/api/v1/1/user/1001/address"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 增加用户收货地址 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 2.修改用户收货地址 接口", func() {
			req := &model.UpdateUserAddress{
				Id:1,
				Name:"vicen",
				CountryId:86,
				ProvinceId:440000,
				CityId:440300,
				DistrictId:440305,
				Address:"科苑路科兴科学园B3栋16楼2",
				PostCode:"518000",
				Phone:"15018521673",
				IsDefault:0,
				IdCard:"4414231990000001",
			}
			jsonByte, _ := json.Marshal(req)
			jsonStr := string(jsonByte)

			requestAPI := "/api/v1/1/user/1001/address"
			result, _ := common.AccessAPIWithPostBody("PUT", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 修改用户收货地址 判定", func() {
				So(result.Status, ShouldEqual, 0)
				//resultData := result.Data.([]interface {})
				//So(len(resultData), ShouldEqual, 2)
			})

		})

		Convey("When 3.拉取用户收货地址列表 接口", func() {
			requestAPI := "/api/v1/1/user/1001/address"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_USER)

			Convey("Then 拉取用户收货地址列表 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 4.删除用户收货地址 接口", func() {
			requestAPI := "/api/v1/1/user/1001/address/1"
			result, _ := common.AccessAPIWithPostBody("DELETE", requestAPI, "", common.BUSINESS_USER)

			Convey("Then 删除用户收货地址 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 5.拉取用户默认收货地址 接口", func() {
			requestAPI := "/api/v1/1/user/1001/defaultaddress"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_USER)

			Convey("Then 拉取用户默认收货地址 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 6.修改用户默认收货地址 接口", func() {
			requestAPI := "/api/v1/1/modifyDefaultaddress"
			jsonStr := `{
    					"userId":1001,
    					"id":3
					}`
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 修改用户默认收货地址 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		///*************账号相关接口 **************/
		//Convey("When 7.账号注册 接口", func() {
		//	req := &model.AddUser{
		//		RegistType:2,//普通注册
		//		OpenId:"12312",
		//		UUID:"dsds01",
		//		OpenType:0,
		//		AccountName:"vicen-test",
		//		Mobile:"15018521673",
		//		NickName:"hello",
		//		HeadIcon:"afsdfs",
		//		Gender:1,
		//		Mail:"244699129@qq.com",
		//		Birthday:"2017-08-28",
		//		SmsId:100,
		//		SmsCode:"988655",
		//		Password:"12312312931823123fh",
		//	}
		//	jsonByte, _ := json.Marshal(req)
		//	jsonStr := string(jsonByte)
		//
		//	requestAPI := "/api/v1/1/userRegist"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 账户注册 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//	})
		//
		//})

		//Convey("When 8.第三方账户登录 接口", func() {
		//	jsonStr := `{
		//			"openId":"oGAXg0QCmA1za881dYuoeFI-eLSA",
    		//			"openType":0
		//		    }`
		//
		//	requestAPI := "/api/v1/1/userOpenIdLogin"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 第三方账户登录 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//	})
		//
		//})

		//Convey("When 9.手机验证码登录 接口", func() {
		//	jsonStr := `{
   		//		        "smsId":0,
    		//		        "mobile":"12345678",
    		//		        "smsCode":"xxxx",
    		//			"autoRegiest":0
		//		    }`
		//
		//	requestAPI := "/api/v1/1/mobileSmsCodeLogin"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 手机验证码登录 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})
		//
		//Convey("When 10.自有账号登录 接口", func() {
		//	jsonStr := `{
   		//			 "loginType":1,
   		//			 "account":"15018521673",
   		//			 "password":"xxxxx",
   		//			 "key":"tw6XqdrKnsh7QubxdPEM%2FeRDfsTUrtZHt7wnADmj5ss%3D",
   		//			 "value":"3548"
		//			}`
		//
		//	requestAPI := "/api/v1/1/userLogin"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 自有账号登录 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})
		//
		//Convey("When 11.第三方账号绑定已有账号 接口", func() {
		//	jsonStr := `{
    		//			"openId":"fadsfsaxxxxx",
   		//			"uuId":"xxxxxx",
    		//			"openType":0,
    		//			"userId":1000000
		//			}`
		//
		//	requestAPI := "/api/v1/1/userOpenIdBind"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 第三方账号绑定已有账号 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})

		Convey("When 12.根据uuid查询userId 接口", func() {
			requestAPI := "/api/v1/1/uuidQueryUserId?uuId=ohz_DjlV4kIoFTgFi3r5iy-4CJyE"
			//form表单 uuId
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_USER)

			Convey("Then 根据uuid查询userId 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		//Convey("When 13.绑定手机 接口", func() {
		//	requestAPI := "/api/v1/1/bindMobile"
		//	jsonStr := `{
		//			"smsId": 530,
		//			"mobile": "123456789",
		//			"smsCode":"123455",
		//			"userId":13
		//			}`
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 绑定手机 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})
		//
		//Convey("When 14.解绑手机 接口", func() {
		//	requestAPI := "/api/v1/1/unbindMobile"
		//	jsonStr := `{
		//			"smsId": 530,
		//			"mobile": "123456789",
		//			"smsCode":"123455",
		//			"userId":13
		//			}`
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 解绑手机 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})

		/*************资料相关接口 **************/
		Convey("When 15.查询基本资料 接口", func() {
			jsonStr := `{
    					"userId":23
					}`

			requestAPI := "/api/v1/1/retrieveUserBaseInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 查询基本资料 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 16.查询扩展资料 接口", func() {
			jsonStr := `{
    					"userId":23
					}`

			requestAPI := "/api/v1/1/retrieveUserExtInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 查询扩展资料 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 17.修改基本资料 接口", func() {
			jsonStr := `{
    				 	"userId":23,
    					"nickName":"vicen-test",
    					"headIcon":"http://xxxx",
   					"gender":0,
    					"mail":"xxx@maizuo.com",
    					"birthday":"1990-10-01"
					}`

			requestAPI := "/api/v1/1/updateUserBaseInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 修改基本资料 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 18.修改扩展资料 接口", func() {
			jsonStr := `{
   					 "userId":23,
   					 "realName":"姓名",
  					 "marriage":0,
    					 "income":"10000-20000",
    					 "identity":"441423199011136000",
   					 "education":2,
   					 "career":1
					}`

			requestAPI := "/api/v1/1/updateUserExtInfo"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 修改扩展资料 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})


		///*************密码相关接口 **************/
		//Convey("When 19.修改密码 接口", func() {
		//	jsonStr := `{
    		//			"userId":1000000,
    		//			"oldPassword":"xxxxxxx",
    		//			"newPassword":"xxxxx",
    		//			"key":"tw6XqdrKnsh7QubxdPEM%2FeRDfsTUrtZHt7wnADmj5ss%3D",
   		//			 "value":"3548"
		//			}`
		//
		//	requestAPI := "/api/v1/1/userChangePassword"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 修改密码 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})
		//
		//Convey("When 20.重置密码 接口", func() {
		//	jsonStr := `{
    		//			"userId":1000000,
    		//			"oldPassword":"xxxxxxx",
    		//			"newPassword":"xxxxx",
    		//			"key":"tw6XqdrKnsh7QubxdPEM%2FeRDfsTUrtZHt7wnADmj5ss%3D",
   		//			 "value":"3548"
		//			}`
		//
		//	requestAPI := "/api/v1/1/userResetPassword"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 重置密码 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})
		//
		//Convey("When 21.设置支付密码 接口", func() {
		//	jsonStr := `{
    		//			"type":0,
  		//		        "userId":1000000,
    		//			"mobile":"12345678",
    		//			"smsId":0,
   		//			"smsCode":"1234"，
    		//			"password":"xxxxx"
		//			}`
		//
		//	requestAPI := "/api/v1/1/userResetPayPwd"
		//	result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)
		//
		//	Convey("Then 设置支付密码 判定", func() {
		//		So(result.Status, ShouldEqual, 0)
		//		resultData := result.Data.([]interface {})
		//		So(len(resultData), ShouldEqual, 2)
		//	})
		//
		//})

		Convey("When 22.校验支付密码 接口", func() {
			jsonStr := `{
    					"userId":24,
    					"password":"4297f44b13955235245b2497399d7a93",
    					"key":"",
    					"value":""
					}`

			requestAPI := "/api/v1/1/checkPayPwd"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 校验支付密码 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 23.查询用户是否有支付密码 接口", func() {
			jsonStr := `{
					"userId":24
					}`

			requestAPI := "/api/v1/1/userhasPayPwd"
			result, _ := common.AccessAPIWithPostBody("POST", requestAPI, jsonStr, common.BUSINESS_USER)

			Convey("Then 查询用户是否有支付密码 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		Convey("When 24.查询用户是否有登录密码 接口", func() {
			requestAPI := "/api/v1/1/userhasPwd?userId=23"
			result, _ := common.AccessAPIWithPostBody("GET", requestAPI, "", common.BUSINESS_USER)

			Convey("Then 查询用户是否有登录密码 判定", func() {
				So(result.Status, ShouldEqual, 0)
			})

		})

		common.ExecuteSQLScriptFile("./../../config/user_clean.sql")
	})

}



