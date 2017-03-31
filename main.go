package main

import (
	_ "codesystem/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"codesystem/models"
)

func init()  {
	models.RegisterDB()
}
func main() {
	orm.Debug=true
	beego.Run()
}

