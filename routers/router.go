package routers

import (
	"codesystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/buildconf",&controllers.BuildConf{})
	beego.Router("/proinfo",&controllers.ProInfo{})


}
