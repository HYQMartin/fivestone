package controllers

import (
	"github.com/astaxie/beego"
)

type MatchingRoomController struct {
	beego.Controller
}

func (t *MatchingRoomController) Get() {
	t.Data["IsWuziqiPVP"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)

	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "MatchingRoom.html"
}
