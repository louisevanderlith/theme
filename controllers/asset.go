package controllers

import (
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/theme/core"
)

type AssetController struct {
	control.APIController
}

func NewAssetCtrl(ctrlMap *control.ControllerMap) *AssetController {
	result := &AssetController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetAsset
// @Description Gets the requested asset
// @Param	appID			path	string 	true		"the application requesting a service"
// @Param	serviceName		path 	string	true		"the name of the service you want to get"
// @Param	clean			path 	bool	false		"clean will return a user friendly URL and not the application's actual URL"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group/:file [get]
func (req *AssetController) Get() {
	group := req.Ctx.Input.Param(":group")
	fileName := req.Ctx.Input.Param(":file")

	res, err := core.FindAsset(group, fileName)
	req.Serve(res, err)
}
