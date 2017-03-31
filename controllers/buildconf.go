package controllers

import (
	"github.com/astaxie/beego"
)
type BuildConf struct {
	beego.Controller
}

func (buildconf *BuildConf)Get(){
	buildconf.TplName="buildconf.html"
}
func (buildconf *BuildConf)Post(){

}