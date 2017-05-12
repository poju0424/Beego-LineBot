package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/123", &controllers.MainController{})
	beego.Handler("/callback", &controllers.LineHandler{})

	beego.Router("/currency/?:time/?:name", &controllers.CurrencyController{}) //monthly
	// beego.Handler("/currency/?:name/ltm", &controllers.LineHandler{})          //last three month
	// beego.Handler("/currency/?:name/l6m", &controllers.LineHandler{})          //last half year
}
