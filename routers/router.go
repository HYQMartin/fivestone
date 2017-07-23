package routers

import (
	"fivestone/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页
	beego.Router("/", &controllers.MainController{})
	beego.Router("/first", &controllers.FirstController{})
	beego.Router("/second", &controllers.SecondController{})

	beego.Router("/THWL", &controllers.THWLController{})
	beego.Router("/Wuziqi/PVC", &controllers.WuziqiController{})

	beego.Router("/Wuziqi/PVP", &controllers.WuziqiPVPController{})
	beego.Router("/Wuziqi/PVP/black", &controllers.WuziqiPVPBController{})
	beego.Router("/Wuziqi/PVP/white", &controllers.WuziqiPVPWController{})
	beego.Router("/Wuziqi/PVP/save", &controllers.WuziqiPVPSController{})
	beego.Router("/Wuziqi/PVP/getall", &controllers.WuziqiPVPGController{})
	beego.Router("/Result", &controllers.ResultController{})

	beego.Router("/Wuziqi/PVP/ChooseDesk", &controllers.ChooseDeskController{})
	beego.Router("/Wuziqi/PVP/ChooseRoom", &controllers.ChooseRoomController{})
	beego.Router("/Wuziqi/PVP/MatchingRoom", &controllers.MatchingRoomController{})

	//beego.Router("/Wuziqi/PVP/Desk_1", &controllers.Desk_1_Controller{})
	beego.Router("/Wuziqi/PVP1", &controllers.WuziqiPVP1Controller{})
	beego.Router("/Wuziqi/PVP/black1", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white1", &controllers.WuziqiPVPW1Controller{})
	beego.Router("/Wuziqi/PVP/save1", &controllers.WuziqiPVPS1Controller{})
	beego.Router("/Wuziqi/PVP/getall1", &controllers.WuziqiPVPG1Controller{})

	beego.Router("/Wuziqi/PVP/black2", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white2", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black3", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white3", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black4", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white4", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black5", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white5", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black6", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white6", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black7", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white7", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black8", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white8", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black9", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white9", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/black10", &controllers.WuziqiPVPB1Controller{})
	beego.Router("/Wuziqi/PVP/white10", &controllers.WuziqiPVPW1Controller{})

	beego.Router("/Wuziqi/PVP/login", &controllers.LoginController{})
	beego.Router("/Wuziqi/PVP/regist", &controllers.RegistController{})

}
