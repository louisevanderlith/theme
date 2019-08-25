package routers

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/theme/controllers/assets"
)

func Setup(poxy resins.Epoxi) {
	//Asset
	metaGroup := routing.NewRouteGroup("asset", mix.JSON)
	metaGroup.AddRoute("Assets by Group", "/{group:[a-z]+}", "GET", roletype.Unknown, assets.GetAll)
	poxy.AddGroup(metaGroup)

	fileGroup := routing.NewRouteGroup("asset", mix.Octet)
	fileGroup.AddRoute("Get Asset", "/{group:[a-z]+}/{file}", "GET", roletype.Unknown, assets.Get)
	poxy.AddGroup(fileGroup)
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
