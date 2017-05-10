package routers

import (
	"hello/controllers"

	"hello/models"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/123", &controllers.MainController{})
	beego.Handler("/callback", &models.LineHandler{})

	beego.Handler("/currency/?:name/?:time", &models.LineHandler{}) //monthly
	beego.Handler("/currency/?:name/ltm", &models.LineHandler{})    //last three month
	beego.Handler("/currency/?:name/l6m", &models.LineHandler{})    //last half year
}
