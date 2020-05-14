package routers

import (
	"github.com/astaxie/beego"
	"github.com/yy/recipe/controllers"
	"github.com/yy/recipe/controllers/lp"
	"github.com/yy/recipe/controllers/lp/at"
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

	// at
	beego.Router("/at/lp", &at.SubPage{}, "*:Lp")
	beego.Router("/at/tan", &at.SubPage{}, "*:Tan")
	beego.Router("/at/confirm", &at.SubPage{}, "*:Confirm")
	beego.Router("/at/tnc", &at.SubPage{}, "*:Condition")
	beego.Router("/at/help", &at.SubPage{}, "*:Help")
	beego.Router("/at/privacy", &at.SubPage{}, "*:Privacy")
	beego.Router("/at/AGB", &at.SubPage{}, "*:AGB")
	beego.Router("/at/Impressum", &at.SubPage{}, "*:Impressum")
	beego.Router("/at/Datenschutz", &at.SubPage{}, "*:Datenschutz")
	beego.Router("/at/KONTAKT", &at.SubPage{}, "*:KONTAKT")
	beego.Router("/at/KUNDIGUNG", &at.SubPage{}, "*:KUNDIGUNG")
	beego.Router("/at/Rucktrittsrechts", &at.SubPage{}, "*:Rucktrittsrechts")
	beego.Router("/at/FAQ", &at.SubPage{}, "*:FAQ")
	beego.Router("/at/welcome", &at.SubPage{}, "*:Welcome")

}
