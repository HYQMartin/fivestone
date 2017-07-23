package main

import (
	"fivestone/models"
	_ "fivestone/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	//models.AddUser()
	//models.AddTotalAndWin("martin", true)
	beego.Run()
}
