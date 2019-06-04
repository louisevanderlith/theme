package routers

import (
	"fmt"
	"strings"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/theme/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	secure "github.com/louisevanderlith/secure/core"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilter(s, host)

	beego.Router("/v1/asset/:group", controllers.NewAssetCtrl(ctrlmap), "get:GetAll")
	beego.Router("/v1/asset/:group/:file", controllers.NewAssetCtrl(ctrlmap), "get:Get")
}

func EnableFilter(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)

	ctrlmap.Add("/v1/asset", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	return ctrlmap
}
