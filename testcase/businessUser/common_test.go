package businessUser

import (
	"testing"
	"maizuo.com/aura/aura-sit/service"
	"maizuo.com/aura/aura-sit/common"
	"fmt"
	"os"
)

func TestMain(m *testing.M) {
	fmt.Println("所有测试执行前")
	service.SetupTestConfig()
	//common.DestroySchema()
	//common.SetupNewSchema()
	//common.CreateTables()
	common.Init()
	common.SetUPDB()
	ret := m.Run()
	fmt.Println("所有测试执行后")
	os.Exit(ret)
}
