package routers

import (
	"run/appweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/appurl", &controllers.AppUrl{})
}
