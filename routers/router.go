package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/theme/controllers"
)

func Setup(poxy *droxolite.Epoxy) {
	//Asset
	assCtrl := &controllers.AssetController{}
	assGroup := droxolite.NewRouteGroup("asset", assCtrl)
	assGroup.AddRoute("/{key:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, assCtrl.Get)
	assGroup.AddRoute("/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Unknown, assCtrl.GetAll)
	poxy.AddGroup(assGroup)
	/*ctrlmap := EnableFilter(s, host)

	beego.Router("/v1/asset/:group", controllers.NewAssetCtrl(ctrlmap), "get:GetAll")
	beego.Router("/v1/asset/:group/:file", controllers.NewAssetCtrl(ctrlmap), "get:Get")*/
}

/*
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
}*/
