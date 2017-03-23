package routers

import (
	"hello/controllers"

	"hello/models"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Handler("/callback", &models.LineHandler{})
}
