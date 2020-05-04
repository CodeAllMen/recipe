package routers

import (
	"github.com/astaxie/beego"
	"github.com/yy/recipe/controllers"
	"github.com/yy/recipe/controllers/lp"
	"github.com/yy/recipe/controllers/lp/mls"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.PlayController{})
	beego.AutoRouter(&controllers.SearchController{})
	beego.Router("/lp", &lp.SubPage{}, "*:Lp")
	beego.Router("/pin", &lp.SubPage{}, "*:Pin")

	beego.Router("/mls/lp", &mls.SubPage{}, "*:Lp")
	beego.Router("/mls/pin", &mls.SubPage{}, "*:Pin")
	beego.Router("/mls/thank", &mls.SubPage{}, "*:Thank")

}
