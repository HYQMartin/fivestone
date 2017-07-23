package controllers

import (
	"github.com/astaxie/beego"
)

type ChooseDeskController struct {
	beego.Controller
}

func (t *ChooseDeskController) Get() {
	t.Data["IsWuziqiPVP"] = true
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	t.TplName = "ChooseDesk.html"
}
