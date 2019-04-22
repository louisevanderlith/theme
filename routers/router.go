// @APIVersion 1.0.0
// @Title Router API
// @Description API for the Router
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/theme/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	secure "github.com/louisevanderlith/secure/core"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	beego.Router("/v1/asset/:group", controllers.NewAssetCtrl(ctrlmap), "get:GetAll")
	beego.Router("/v1/asset/:group/:file", controllers.NewAssetCtrl(ctrlmap), "get:Get")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)

	ctrlmap.Add("/asset", emptyMap)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	return ctrlmap
}
