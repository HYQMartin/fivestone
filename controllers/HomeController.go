package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["FirstStep"] = `
			<div align="middle">
				<nav>
  				<ul class="pagination pagination-lg">
   				<li>
      				<a href="#" aria-label="Previous">
        			<span aria-hidden="true">&laquo;</span>
      				</a>
    			</li>
    			<li><a href="/first">第一阶段(1.27-2.4)：</a></li>
    			<li><a href="/second">第二阶段(2.12-~)：</a></li>
    			<li>
      				<a href="#" aria-label="Next">
        			<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>
  				</ul>
				</nav>
			</div>
	`
	c.Data["Author"] = "叶青，子龙"
	c.Data["Time"] = time.Now().Format(time.Stamp)
	c.Data["IsHome"] = true

	namestruct, err := c.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		c.Data["UserName"] = name

	}

	c.Data["IsLogin"] = CheckAccount(c.Ctx)
	c.TplName = "Home.html"

}

type FirstController struct {
	beego.Controller
}

func (c *FirstController) Get() {
	c.TplName = "Home.html"
	c.Data["FirstStep"] = `
			<div align="middle">
				<nav>
  				<ul class="pagination pagination-lg">
   				<li>
      				<a href="/" aria-label="Previous">
        			<span aria-hidden="true">&laquo;</span>
      				</a>
    			</li>
    			<li class="active"><a href="/first">第一阶段(1.27-2.4)：</a></li>
    			<li><a href="/second">第二阶段(2.12-~)：</a></li>
    			<li>
      				<a href="/second" aria-label="Next">
        			<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>
  				</ul>
				</nav>
			</div>
				<h4 class="text-muted">起始于 2015-1-27</h4>
				<p>1-27：一堆不懂，不确定搞得出来，试着开始做做！ 只弄了30min - -</p>
				<p>1-28：准备弄却只剩下半个小时弄，然后睡觉,弄了点THWL页面</p>
				<p>1-29：无耻地内嵌一个人机版本，PVP版本暂时不知道怎么下手</p>
				<p>1-30：PVP基本搞定，还有很多没做。还需要做轮流下棋，另外统计胜率，做Result页面。下棋右边可以做个页面聊天的。胜利一方显示胜利，输的一方显示输了。</p>
				<p>1-31：做了轮流下棋。胜率统计和Result页面待做。聊天再考虑。另外出现下棋速度慢的不好体验。</p>
				<p>2-1：出去玩了</p>
				<P>2-2：以后下班后开始做。今天啥都没做 - -！</P>
				<p>2-3：没带最新代码，玩游戏去了</p>
				<p>2-4：子龙加入，顿时前段好看了不少。修改了部分小问题</p>
				<p>2-5：2-6回深圳  暂停</p>
	`
	c.Data["Author"] = "叶青，子龙"
	c.Data["Time"] = time.Now().Format(time.Stamp)
	c.Data["IsHome"] = true

}

type SecondController struct {
	beego.Controller
}

func (c *SecondController) Get() {
	c.Data["FirstStep"] = `
			<div align="middle">
				<nav>
  				<ul class="pagination pagination-lg">
   				<li>
      				<a href="/first" aria-label="Previous">
        			<span aria-hidden="true">&laquo;</span>
      				</a>
    			</li>
    			<li><a href="/first">第一阶段(1.27-2.4)：</a></li>
    			<li class="active"><a href="/second">第二阶段(2.12-~)：</a></li>
    			<li class="disabled">
      				<a href="#" aria-label="Next">
        			<span aria-hidden="true">&raquo;</span>
      				</a>
    			</li>
  				</ul>
				</nav>
			</div>
				<h4 class="text-muted">起始于 2015-1-27</h4>
				<p>2-12：继续做五子棋模块。</p>
				<p>2-13: 准备制作悔棋  </p>
				<p>2-14: 准备修改显示上个棋子，另外去除刷新页面这一操作，把部分按钮请求去除跳转页面操作，仅仅上传数据</p>
				<p>2-15: 开始做多张桌子，支持多对玩家娱乐 </p>
				<p>2-16~2-22:春节休息</p>
				<p>2-23:做了用户登录注册，修改了Result页面</p>
				<br></br> <br></br> <br></br><br></br><br></br><br></br><br></br><br></br>
	`
	c.Data["Author"] = "叶青，子龙"
	c.Data["Time"] = time.Now().Format(time.Stamp)
	c.Data["IsHome"] = true
	c.TplName = "Home.html"

}
