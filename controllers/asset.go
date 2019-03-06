package controllers

import (
	"strings"

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
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Param	file		path 	string	true		"full filename /html/master.entry.html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group/:file [get]
func (req *AssetController) Get() {
	group := req.Ctx.Input.Param(":group")
	fileName := req.Ctx.Input.Param(":file")

	res, err := core.FindAsset(group, fileName)

	if err != nil {
		panic(err)
	}

	mimes := make(map[string]string)
	mimes["js"] = "text/javascript"
	mimes["css"] = "text/css"
	mimes["html"] = "text/html"
	mimes["ico"] = "image/x-icon"
	mimes["font"] = "font/" + getExt(fileName)

	//TODO: Remove, and rather use Artifact.API
	mimes["img"] = "image/png"

	req.ServeBinaryWithMIME(res, fileName, mimes[group])
}

// @Title GetAsset List
// @Description Lists the assets in a group
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group [get]
func (req *AssetController) GetAll() {
	group := req.Ctx.Input.Param(":group")

	req.Serve(core.ListAssets(group))
}

func getExt(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	return filename[dotIndex+1:]
}
