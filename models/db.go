package models
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func RegisterDB()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:jackandtom@tcp(192.168.0.59:3307)/dafy?charset=utf8")
}

func SearchProInfo(Projectgroup,Clustername string)(info  []orm.ParamsList){
	orm.Debug=true
	o:=orm.NewOrm()
	o.Raw("select ProjectName,DomainAlias,DomainName,TomcatName,TomcatPort,ProjectDescription,ClusterName,UserAlias from tt where ProjectGroup=? and ClusterName=?;",Projectgroup,Clustername).ValuesList(&info)
	return
}

