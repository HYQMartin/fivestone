package controllers

import (
	"github.com/astaxie/beego"
)

var Desk_1 OneGame
var Desk_2 OneGame
var Desk_3 OneGame
var Desk_4 OneGame
var Desk_5 OneGame
var Desk_6 OneGame
var Desk_7 OneGame
var Desk_8 OneGame
var Desk_9 OneGame
var Desk_10 OneGame

func init() {
	Desk_1 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_2 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_3 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_4 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_5 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_6 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_7 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_8 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_9 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	Desk_10 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
}

type WuziqiPVPB1Controller struct {
	beego.Controller
}

func (t *WuziqiPVPB1Controller) Get() {
	desknum := t.GetString("desknum")
	t.Data["SelectColor"] = "/Wuziqi/PVP/white" + desknum + "?desknum=" + desknum
	t.Data["DeskNum"] = desknum
	t.Data["Color_Chiness"] = "黑色"
	t.Data["StoneColor"] = "black"
	t.Data["Count"] = GetCount(desknum)
	t.Data["OtherStoneColor"] = "white"
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	t.Data["ChessArr"] = GetChessArr_Desk(desknum)
	t.Data["IsWuziqiPVP"] = true
	t.TplNames = "WuziqiDesk.html"
	return
}

type WuziqiPVPW1Controller struct {
	beego.Controller
}

func (t *WuziqiPVPW1Controller) Get() {
	desknum := t.GetString("desknum")
	t.Data["SelectColor"] = "/Wuziqi/PVP/black" + desknum + "?desknum=" + desknum
	t.Data["DeskNum"] = desknum
	t.Data["Color_Chiness"] = "白色"
	t.Data["StoneColor"] = "white"
	t.Data["OtherStoneColor"] = "black"
	t.Data["Count"] = GetCount(desknum)
	t.Data["IsWuziqiPVP"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	t.Data["ChessArr"] = GetChessArr_Desk(desknum)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "WuziqiDesk.html"
	return
}
