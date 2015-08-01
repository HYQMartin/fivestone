package controllers

import (
	"github.com/astaxie/beego"
)

type ChooseRoomController struct {
	beego.Controller
}

func (t *ChooseRoomController) Get() {
	t.Data["IsWuziqiPVP"] = true
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	t.TplNames = "ChooseRoom.html"
}
