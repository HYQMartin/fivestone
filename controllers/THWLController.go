package controllers

import (
	"github.com/astaxie/beego"
)

type THWLController struct {
	beego.Controller
}

func (t *THWLController) Get() {
	t.Data["IsTHWL"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "THWL.html"
}
