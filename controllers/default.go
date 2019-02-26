package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//hp := strconv.Itoa(beego.BConfig.Listen.HTTPPort)
	//c.Ctx.WriteString("App name: " + beego.BConfig.AppName +
	//	"\nHttp port: " +hp +
	//	"\nRun mode: " + beego.BConfig.RunMode)
	//
	//logs.SetLogger("console")
	//logs.SetLevel(logs.LevelInfo)
	//logs.Trace("Trace test")
	//logs.Info("Info test")

	c.TplName = "home.html"

	//c.Data["Website"] = "earchow.com"
	//c.Data["Email"] = "ziranjuanchow@gmail.com"
	//
	//c.Data["TrueCond"] = true
	//c.Data["FlaseCond"] = false
	//
	//type u struct {
	//	Name string
	//	Age  int
	//	Sex  string
	//}
	//
	//user := &u{
	//	Name: "Chow",
	//	Age:  20,
	//	Sex:  "Male",
	//}
	//
	//c.Data["User"] = user
	//
	//nums:=[]int{0,1,2,3,4,5,6,7,8,9}
	//c.Data["Nums"] = nums
	//
	////mo ban bian liang
	//c.Data["TplVar"] = "Hey guys."
	//
	//c.Data["Html"] = "<div>Hello beego</div>"
	//
	//c.Data["Pipe"] = "<div>Hello beego</div>"
}
