package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Handler("/callback", &controllers.LineHandler{})
	beego.Handler("/currency/?:time/?:name", &controllers.CurrencyHandler{})

	// beego.Router("/123", &controllers.MainController{})
}
