package controllers

import (
"github.com/astaxie/beego"
"codesystem/models"
)
type ProInfo struct {
	beego.Controller
}
var (
	Projectgroup,Clustername string
)

func (proinfo *ProInfo)Get(){

	Proinfosearch:=proinfo.Input().Get("Proinfosearch")

	switch Proinfosearch {
	case "on":

		proinfo.Redirect("/proinfo",301)
		return
	default:

	}
	proinfo.TplName="proinfo.html"
	Projectgroup=proinfo.Input().Get("projectgroup")
	Clustername=proinfo.Input().Get("clustername")
	proinfo.Data["prostruct"]=models.SearchProInfo(Projectgroup,Clustername)

}
func (proinfo *ProInfo)Post() {

}

