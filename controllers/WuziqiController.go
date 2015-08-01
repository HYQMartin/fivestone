package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type WuziqiController struct {
	beego.Controller
}

func (t *WuziqiController) Get() {
	t.Data["IsWuziqi"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "Wuziqi.html"
}

var TotalGame map[string]OneGame
var Desk_0 OneGame

type OneGame struct {
	//Blacker    string         //black name
	//White      string         //white name
	Desk       int            //第几桌
	Count      int            //当前桌第几盘
	White_Done string         //白色棋子是否下完
	Winner     string         //black  or white
	AGobang    map[int]AStone //统一成Desk_0.Winner
	GobangSave map[int]interface{}
}

func GetChessArr() (chessarr [15][15]int) {
	for _, astone := range Desk_0.AGobang {
		chessarr[astone.X][astone.Y] = astone.Color
	}
	return
}
func AddAStone(x int, y int, color int) {
	l := len(Desk_0.AGobang)
	one := AStone{x, y, color}
	Desk_0.AGobang[l+1] = one
	return
}
func NewGobang() {
	l := len(Desk_0.AGobang)
	for i := 1; i <= l; i++ {
		delete(Desk_0.AGobang, i)
	}
	return
}
func DelAStone() {
	if len(Desk_0.AGobang) == 0 {
		Desk_0.White_Done = "reset"
		return
	}
	if len(Desk_0.AGobang)%2 == 1 {
		Desk_0.White_Done = "yes"
	} else {
		Desk_0.White_Done = "no"
	}
	delete(Desk_0.AGobang, len(Desk_0.AGobang))
	return
}

type AStone struct {
	X     int `json:"X"`
	Y     int `json:"Y"`
	Color int `json:"Color"` //1=white,-1=white,0=no
}

func init() {
	Desk_0 = OneGame{
		AGobang:    make(map[int]AStone),
		GobangSave: make(map[int]interface{}),
	}
	//---------
	TotalGame = make(map[string]OneGame)
}

type WuziqiPVPController struct {
	beego.Controller
}

func (t *WuziqiPVPController) Get() {
	a, _ := t.GetInt("new")
	b, _ := t.GetInt("regreat")
	if a != 0 {
		if len(Desk_0.AGobang) != 0 {
			Desk_0.Count = Desk_0.Count + 1
			Desk_0.Winner = "reset"
			Desk_0.White_Done = "reset"
			//IsGameOver = false
			fmt.Println("--------------这是第", Desk_0.Count, "盘五子棋", "----------------")
			Desk_0.GobangSave[Desk_0.Count] = Desk_0.AGobang
			NewGobang()
		}
	}
	if b != 0 {
		DelAStone()
	}
	t.Data["Color_Chiness"] = "请选择颜色"
	t.Data["OtherStoneColor"] = "white"
	t.Data["Desk_0.Count"] = Desk_0.Count
	t.Data["ChessArr"] = GetChessArr()
	t.Data["IsWuziqiPVP"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "WuziqiPVP.html"
}

type WuziqiPVPBController struct {
	beego.Controller
}

func (t *WuziqiPVPBController) Get() {
	winner := t.GetString("win")
	if winner == "black" {
		Desk_0.Winner = "black"
		//IsGameOver = true
	}
	t.Data["Color_Chiness"] = "黑色"
	t.Data["StoneColor"] = "black"
	t.Data["Desk_0.Count"] = Desk_0.Count
	t.Data["OtherStoneColor"] = "white"
	t.Data["ChessArr"] = GetChessArr()
	t.Data["IsWuziqiPVP"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "WuziqiPVP.html"
}

type WuziqiPVPWController struct {
	beego.Controller
}

func (t *WuziqiPVPWController) Get() {
	winner := t.GetString("win")
	if winner == "white" {
		Desk_0.Winner = "white"
		//IsGameOver = true
	}
	t.Data["Color_Chiness"] = "白色"
	t.Data["StoneColor"] = "white"
	t.Data["OtherStoneColor"] = "black"
	t.Data["Desk_0.Count"] = Desk_0.Count
	t.Data["IsWuziqiPVP"] = true
	t.Data["ChessArr"] = GetChessArr()
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "WuziqiPVP.html"
}

type WuziqiPVPGController struct {
	beego.Controller
}

func (t *WuziqiPVPGController) Get() {
	var askflush map[string]string
	askflush = make(map[string]string)
	askflush["flush"] = "yes"
	aj, _ := json.Marshal(askflush)

	var askflushno map[string]interface{}
	askflushno = make(map[string]interface{})
	askflushno["flush"] = "no"
	askflushno["last"] = []int{Desk_0.AGobang[len(Desk_0.AGobang)].X, Desk_0.AGobang[len(Desk_0.AGobang)].Y, Desk_0.AGobang[len(Desk_0.AGobang)].Color}
	nj, _ := json.Marshal(askflushno)
	if len(Desk_0.AGobang) == 0 {
		t.Ctx.WriteString(string(aj))
		return
	}
	if Desk_0.Winner == "white" || Desk_0.Winner == "black" {
		t.Ctx.WriteString(string(nj))
		return
	}
	var gobang map[string]interface{}
	gobang = make(map[string]interface{})
	for k, v := range Desk_0.AGobang {
		var astone []int
		astone = append(astone, v.X)
		astone = append(astone, v.Y)
		astone = append(astone, v.Color)
		gobang[strconv.Itoa(k)] = astone
	}
	gobang["len"] = len(Desk_0.AGobang)
	js, err := json.Marshal(gobang)
	if err != nil {
		fmt.Println(err)
	}
	t.Ctx.WriteString(string(js))
	return
}

type WuziqiPVPSController struct {
	beego.Controller
}

func (t *WuziqiPVPSController) Get() {
	iam := t.GetString("iam")
	if iam != "" {
		if Desk_0.Winner == "black" || Desk_0.Winner == "white" {
			t.Ctx.WriteString("no")
			return
		}
		if iam == "white" {
			if len(Desk_0.AGobang) == 0 {
				t.Ctx.WriteString("no")
				return
			}
			if Desk_0.White_Done == "yes" {
				t.Ctx.WriteString("no")
				return
			} else {
				t.Ctx.WriteString("yes")
				Desk_0.White_Done = "yes"
				x, _ := t.GetInt("x")
				y, _ := t.GetInt("y")
				AddAStone(x, y, 1)
				return
			}
		}
		if iam == "black" {
			if Desk_0.White_Done == "no" {
				t.Ctx.WriteString("no")
				return
			} else {
				t.Ctx.WriteString("yes")
				Desk_0.White_Done = "no"
				x, _ := t.GetInt("x")
				y, _ := t.GetInt("y")
				AddAStone(x, y, -1)
				return
			}
		}
	}
	t.Ctx.WriteString("end")
}
