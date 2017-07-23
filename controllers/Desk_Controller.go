package controllers

import (
	"fivestone/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type WuziqiPVP1Controller struct {
	beego.Controller
}

func (t *WuziqiPVP1Controller) Get() {
	over := t.GetString("over")
	a, _ := t.GetInt("new")
	b, _ := t.GetInt("regreat")
	num := t.GetString("num")
	if a != 0 {
		if GetAGobangLen(num) != 0 {
			c := GetCount(num)
			SetCount(num, c+1)
			SetWinner(num, "reset")
			ChangeWhite_Done(num, "reset")
			//IsGameOver = false
			fmt.Println("--------------这是第", GetCount(num), "盘五子棋", "----------------")
			SaveGobang(num)
			NewGobang_Desk(num)
		}
		return
	}
	if over != "" {
		na, err := t.Ctx.Request.Cookie("name")
		aname := na.Value
		if err != nil {
			return
		}
		if over == "winner" {

			if aname != "" {
				models.AddTotalAndWin(aname, true)
			}
		}
		if over == "loser" {
			if aname != "" {
				models.AddTotalAndWin(aname, false)
			}
		}
	}

	if b != 0 {
		DelAStone_Desk(num)
		return
	}
	desknum := t.GetString("desknum")

	t.Data["SelectColor"] = "/Wuziqi/PVP/black" + desknum + "?desknum=" + desknum
	fmt.Println(t.Data["SelectColor"])
	t.Data["DeskNum"] = desknum
	t.Data["Color_Chiness"] = "请选择颜色"
	t.Data["OtherStoneColor"] = "black"
	t.Data["Count"] = GetCount(desknum)
	t.Data["ChessArr"] = GetChessArr_Desk(desknum)
	t.Data["IsWuziqiPVP"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplName = "WuziqiDesk.html"
}

//var TotalGame map[string]OneGame
func GetChessArr_Desk(num string) (chessarr [15][15]int) {
	switch num {
	case "1":
		for _, astone := range Desk_1.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "2":
		for _, astone := range Desk_2.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "3":
		for _, astone := range Desk_3.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "4":
		for _, astone := range Desk_4.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "5":
		for _, astone := range Desk_5.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "6":
		for _, astone := range Desk_6.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "7":
		for _, astone := range Desk_7.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "8":
		for _, astone := range Desk_8.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "9":
		for _, astone := range Desk_9.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	case "10":
		for _, astone := range Desk_10.AGobang {
			chessarr[astone.X][astone.Y] = astone.Color
		}
	}

	return
}
func AddAStone_Desk(num string, x int, y int, color int) {
	one := AStone{x, y, color}
	switch num {
	case "1":
		l := len(Desk_1.AGobang)
		Desk_1.AGobang[l+1] = one
	case "2":
		l := len(Desk_2.AGobang)
		Desk_2.AGobang[l+1] = one
	case "3":
		l := len(Desk_3.AGobang)
		Desk_3.AGobang[l+1] = one
	case "4":
		l := len(Desk_4.AGobang)
		Desk_4.AGobang[l+1] = one
	case "5":
		l := len(Desk_5.AGobang)
		Desk_5.AGobang[l+1] = one
	case "6":
		l := len(Desk_6.AGobang)
		Desk_6.AGobang[l+1] = one
	case "7":
		l := len(Desk_7.AGobang)
		Desk_7.AGobang[l+1] = one
	case "8":
		l := len(Desk_8.AGobang)
		Desk_8.AGobang[l+1] = one
	case "9":
		l := len(Desk_9.AGobang)
		Desk_9.AGobang[l+1] = one
	case "10":
		l := len(Desk_10.AGobang)
		Desk_10.AGobang[l+1] = one
	}
	return
}
func NewGobang_Desk(num string) {
	switch num {
	case "1":
		l := len(Desk_1.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_1.AGobang, i)
		}
	case "2":
		l := len(Desk_2.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_2.AGobang, i)
		}
	case "3":
		l := len(Desk_3.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_3.AGobang, i)
		}
	case "4":
		l := len(Desk_4.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_4.AGobang, i)
		}
	case "5":
		l := len(Desk_5.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_5.AGobang, i)
		}
	case "6":
		l := len(Desk_6.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_6.AGobang, i)
		}
	case "7":
		l := len(Desk_7.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_7.AGobang, i)
		}
	case "8":
		l := len(Desk_8.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_8.AGobang, i)
		}
	case "9":
		l := len(Desk_9.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_9.AGobang, i)
		}
	case "10":
		l := len(Desk_10.AGobang)
		for i := 1; i <= l; i++ {
			delete(Desk_10.AGobang, i)
		}
	}
	return
}
func DelAStone_Desk(num string) {
	switch num {
	case "1":
		if len(Desk_1.AGobang) == 0 {
			Desk_1.White_Done = "reset"
			return
		}
		if len(Desk_1.AGobang)%2 == 1 {
			Desk_1.White_Done = "yes"
		} else {
			Desk_1.White_Done = "no"
		}
		delete(Desk_1.AGobang, len(Desk_1.AGobang))
	case "2":
		if len(Desk_2.AGobang) == 0 {
			Desk_2.White_Done = "reset"
			return
		}
		if len(Desk_2.AGobang)%2 == 1 {
			Desk_2.White_Done = "yes"
		} else {
			Desk_2.White_Done = "no"
		}
		delete(Desk_2.AGobang, len(Desk_2.AGobang))
	case "3":
		if len(Desk_3.AGobang) == 0 {
			Desk_3.White_Done = "reset"
			return
		}
		if len(Desk_3.AGobang)%2 == 1 {
			Desk_3.White_Done = "yes"
		} else {
			Desk_3.White_Done = "no"
		}
		delete(Desk_3.AGobang, len(Desk_3.AGobang))
	case "4":
		if len(Desk_4.AGobang) == 0 {
			Desk_4.White_Done = "reset"
			return
		}
		if len(Desk_4.AGobang)%2 == 1 {
			Desk_4.White_Done = "yes"
		} else {
			Desk_4.White_Done = "no"
		}
		delete(Desk_4.AGobang, len(Desk_4.AGobang))
	case "5":
		if len(Desk_5.AGobang) == 0 {
			Desk_5.White_Done = "reset"
			return
		}
		if len(Desk_5.AGobang)%2 == 1 {
			Desk_5.White_Done = "yes"
		} else {
			Desk_5.White_Done = "no"
		}
		delete(Desk_5.AGobang, len(Desk_5.AGobang))
	case "6":
		if len(Desk_6.AGobang) == 0 {
			Desk_6.White_Done = "reset"
			return
		}
		if len(Desk_6.AGobang)%2 == 1 {
			Desk_6.White_Done = "yes"
		} else {
			Desk_6.White_Done = "no"
		}
		delete(Desk_6.AGobang, len(Desk_6.AGobang))
	case "7":
		if len(Desk_7.AGobang) == 0 {
			Desk_7.White_Done = "reset"
			return
		}
		if len(Desk_7.AGobang)%2 == 1 {
			Desk_7.White_Done = "yes"
		} else {
			Desk_7.White_Done = "no"
		}
		delete(Desk_7.AGobang, len(Desk_7.AGobang))
	case "8":
		if len(Desk_8.AGobang) == 0 {
			Desk_8.White_Done = "reset"
			return
		}
		if len(Desk_8.AGobang)%2 == 1 {
			Desk_8.White_Done = "yes"
		} else {
			Desk_8.White_Done = "no"
		}
		delete(Desk_8.AGobang, len(Desk_8.AGobang))
	case "9":
		if len(Desk_9.AGobang) == 0 {
			Desk_9.White_Done = "reset"
			return
		}
		if len(Desk_9.AGobang)%2 == 1 {
			Desk_9.White_Done = "yes"
		} else {
			Desk_9.White_Done = "no"
		}
		delete(Desk_9.AGobang, len(Desk_9.AGobang))
	case "10":
		if len(Desk_10.AGobang) == 0 {
			Desk_10.White_Done = "reset"
			return
		}
		if len(Desk_10.AGobang)%2 == 1 {
			Desk_10.White_Done = "yes"
		} else {
			Desk_10.White_Done = "no"
		}
		delete(Desk_10.AGobang, len(Desk_10.AGobang))
	}
	return
}

type WuziqiPVPG1Controller struct {
	beego.Controller
}

func (t *WuziqiPVPG1Controller) Get() {
	num := t.GetString("num")

	var askflush map[string]string
	askflush = make(map[string]string)
	askflush["flush"] = "yes"
	aj, _ := json.Marshal(askflush)

	var askflushno map[string]interface{}
	askflushno = make(map[string]interface{})
	askflushno["flush"] = "no"
	askflushno["last"] = GetLastChess(num)
	nj, _ := json.Marshal(askflushno)
	if GetAGobangLen(num) == 0 {
		t.Ctx.WriteString(string(aj))
		return
	}
	if GetWinner(num) == "white" || GetWinner(num) == "black" {
		t.Ctx.WriteString(string(nj))
		return
	}
	js := GetAllChess(num)
	t.Ctx.WriteString(js)
	return
}
func GetLastChess(num string) (last []int) {
	switch num {
	case "1":
		last = []int{Desk_1.AGobang[len(Desk_1.AGobang)].X, Desk_1.AGobang[len(Desk_1.AGobang)].Y, Desk_1.AGobang[len(Desk_1.AGobang)].Color}
	case "2":
		last = []int{Desk_2.AGobang[len(Desk_2.AGobang)].X, Desk_2.AGobang[len(Desk_2.AGobang)].Y, Desk_2.AGobang[len(Desk_2.AGobang)].Color}
	case "3":
		last = []int{Desk_3.AGobang[len(Desk_3.AGobang)].X, Desk_3.AGobang[len(Desk_3.AGobang)].Y, Desk_3.AGobang[len(Desk_3.AGobang)].Color}
	case "4":
		last = []int{Desk_4.AGobang[len(Desk_4.AGobang)].X, Desk_4.AGobang[len(Desk_4.AGobang)].Y, Desk_4.AGobang[len(Desk_4.AGobang)].Color}
	case "5":
		last = []int{Desk_5.AGobang[len(Desk_5.AGobang)].X, Desk_5.AGobang[len(Desk_5.AGobang)].Y, Desk_5.AGobang[len(Desk_5.AGobang)].Color}
	case "6":
		last = []int{Desk_6.AGobang[len(Desk_6.AGobang)].X, Desk_6.AGobang[len(Desk_6.AGobang)].Y, Desk_6.AGobang[len(Desk_6.AGobang)].Color}
	case "7":
		last = []int{Desk_7.AGobang[len(Desk_7.AGobang)].X, Desk_7.AGobang[len(Desk_7.AGobang)].Y, Desk_7.AGobang[len(Desk_7.AGobang)].Color}
	case "8":
		last = []int{Desk_8.AGobang[len(Desk_8.AGobang)].X, Desk_8.AGobang[len(Desk_8.AGobang)].Y, Desk_8.AGobang[len(Desk_8.AGobang)].Color}
	case "9":
		last = []int{Desk_9.AGobang[len(Desk_9.AGobang)].X, Desk_9.AGobang[len(Desk_9.AGobang)].Y, Desk_9.AGobang[len(Desk_9.AGobang)].Color}
	case "10":
		last = []int{Desk_10.AGobang[len(Desk_10.AGobang)].X, Desk_10.AGobang[len(Desk_10.AGobang)].Y, Desk_10.AGobang[len(Desk_10.AGobang)].Color}
	}
	return
}
func GetCount(num string) (count int) {
	switch num {
	case "1":
		count = Desk_1.Count
	case "2":
		count = Desk_2.Count
	case "3":
		count = Desk_3.Count
	case "4":
		count = Desk_4.Count
	case "5":
		count = Desk_5.Count
	case "6":
		count = Desk_6.Count
	case "7":
		count = Desk_7.Count
	case "8":
		count = Desk_8.Count
	case "9":
		count = Desk_9.Count
	case "10":
		count = Desk_10.Count
	}
	return
}
func SetCount(num string, count int) {
	switch num {
	case "1":
		Desk_1.Count = count
	case "2":
		Desk_2.Count = count
	case "3":
		Desk_3.Count = count
	case "4":
		Desk_4.Count = count
	case "5":
		Desk_5.Count = count
	case "6":
		Desk_6.Count = count
	case "7":
		Desk_7.Count = count
	case "8":
		Desk_8.Count = count
	case "9":
		Desk_9.Count = count
	case "10":
		Desk_10.Count = count
	}
	return
}
func SetWinner(num string, winner string) {
	switch num {
	case "1":
		Desk_1.Winner = winner
	case "2":
		Desk_2.Winner = winner
	case "3":
		Desk_3.Winner = winner
	case "4":
		Desk_4.Winner = winner
	case "5":
		Desk_5.Winner = winner
	case "6":
		Desk_6.Winner = winner
	case "7":
		Desk_7.Winner = winner
	case "8":
		Desk_8.Winner = winner
	case "9":
		Desk_9.Winner = winner
	case "10":
		Desk_10.Winner = winner
	}
	return
}
func SaveGobang(num string) {
	switch num {
	case "1":
		Desk_1.GobangSave[Desk_1.Count] = Desk_1.AGobang
	case "2":
		Desk_2.GobangSave[Desk_2.Count] = Desk_2.AGobang
	case "3":
		Desk_3.GobangSave[Desk_3.Count] = Desk_3.AGobang
	case "4":
		Desk_4.GobangSave[Desk_4.Count] = Desk_4.AGobang
	case "5":
		Desk_5.GobangSave[Desk_5.Count] = Desk_5.AGobang
	case "6":
		Desk_6.GobangSave[Desk_6.Count] = Desk_6.AGobang
	case "7":
		Desk_7.GobangSave[Desk_7.Count] = Desk_7.AGobang
	case "8":
		Desk_8.GobangSave[Desk_8.Count] = Desk_8.AGobang
	case "9":
		Desk_9.GobangSave[Desk_9.Count] = Desk_9.AGobang
	case "10":
		Desk_10.GobangSave[Desk_10.Count] = Desk_10.AGobang
	}
	return
}
func GetAllChess(num string) string {
	var gobang map[string]interface{}
	gobang = make(map[string]interface{})
	var desk OneGame
	switch num {
	case "1":
		desk = Desk_1
	case "2":
		desk = Desk_2
	case "3":
		desk = Desk_3
	case "4":
		desk = Desk_4
	case "5":
		desk = Desk_5
	case "6":
		desk = Desk_6
	case "7":
		desk = Desk_7
	case "8":
		desk = Desk_8
	case "9":
		desk = Desk_9
	case "10":
		desk = Desk_10
	}
	for k, v := range desk.AGobang {
		var astone []int
		astone = append(astone, v.X)
		astone = append(astone, v.Y)
		astone = append(astone, v.Color)
		gobang[strconv.Itoa(k)] = astone
	}
	gobang["len"] = len(desk.AGobang)
	js, err := json.Marshal(gobang)
	if err != nil {
		fmt.Println(err)
	}
	return string(js)
}
func GetAGobangLen(num string) (length int) {
	switch num {
	case "1":
		length = len(Desk_1.AGobang)
	case "2":
		length = len(Desk_2.AGobang)
	case "3":
		length = len(Desk_3.AGobang)
	case "4":
		length = len(Desk_4.AGobang)
	case "5":
		length = len(Desk_5.AGobang)
	case "6":
		length = len(Desk_6.AGobang)
	case "7":
		length = len(Desk_7.AGobang)
	case "8":
		length = len(Desk_8.AGobang)
	case "9":
		length = len(Desk_9.AGobang)
	case "10":
		length = len(Desk_10.AGobang)
	}
	return
}
func GetWinner(num string) (winner string) {
	switch num {
	case "1":
		winner = Desk_1.Winner
	case "2":
		winner = Desk_2.Winner
	case "3":
		winner = Desk_3.Winner
	case "4":
		winner = Desk_4.Winner
	case "5":
		winner = Desk_5.Winner
	case "6":
		winner = Desk_6.Winner
	case "7":
		winner = Desk_7.Winner
	case "8":
		winner = Desk_8.Winner
	case "9":
		winner = Desk_9.Winner
	case "10":
		winner = Desk_10.Winner
	}
	return
}

type WuziqiPVPS1Controller struct {
	beego.Controller
}

func (t *WuziqiPVPS1Controller) Get() {
	winner := t.GetString("win")
	if winner != "" {
		num := t.GetString("num")
		if winner == "black" {
			SetWinner(num, "black")
		}
		if winner == "white" {
			SetWinner(num, "white")
		}
	}

	iam := t.GetString("iam")
	if len(iam) < 6 {
		t.Ctx.WriteString("end")
		return
	}
	s1 := []byte(iam)
	s2 := s1[0:5]
	s3 := s1[5:]
	color := string(s2)
	num := string(s3)
	var totalDesk map[string]OneGame
	totalDesk = make(map[string]OneGame)
	totalDesk["1"] = Desk_1
	totalDesk["2"] = Desk_2
	totalDesk["3"] = Desk_3
	totalDesk["4"] = Desk_4
	totalDesk["5"] = Desk_5
	totalDesk["6"] = Desk_6
	totalDesk["7"] = Desk_7
	totalDesk["8"] = Desk_8
	totalDesk["9"] = Desk_9
	totalDesk["10"] = Desk_10

	if totalDesk[num].Winner == "black" || totalDesk[num].Winner == "white" {
		t.Ctx.WriteString("no")
		return
	}
	if color == "white" {
		if len(totalDesk[num].AGobang) == 0 {
			t.Ctx.WriteString("no")
			return
		}
		if totalDesk[num].White_Done == "yes" {
			t.Ctx.WriteString("no")
			return
		} else {
			t.Ctx.WriteString("yes")
			ChangeWhite_Done(num, "yes")
			x, _ := t.GetInt("x")
			y, _ := t.GetInt("y")
			AddAStone_Desk(num, x, y, 1)
			return
		}
	}
	if color == "black" {
		if totalDesk[num].White_Done == "no" {
			t.Ctx.WriteString("no")
			return
		} else {
			t.Ctx.WriteString("yes")
			ChangeWhite_Done(num, "no")
			x, _ := t.GetInt("x")
			y, _ := t.GetInt("y")
			AddAStone_Desk(num, x, y, -1)
			return
		}
	}

	t.Ctx.WriteString("end")
	return
}
func ChangeWhite_Done(num string, value string) {
	switch num {
	case "1":
		Desk_1.White_Done = value
	case "2":
		Desk_2.White_Done = value
	case "3":
		Desk_3.White_Done = value
	case "4":
		Desk_4.White_Done = value
	case "5":
		Desk_5.White_Done = value
	case "6":
		Desk_6.White_Done = value
	case "7":
		Desk_7.White_Done = value
	case "8":
		Desk_8.White_Done = value
	case "9":
		Desk_9.White_Done = value
	case "10":
		Desk_10.White_Done = value
	}
}
