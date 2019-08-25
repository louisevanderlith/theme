package assets

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/theme/core"
)

// @Title GetAsset
// @Description Gets the requested asset
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Param	file		path 	string	true		"full filename /html/master.entry.html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group/:file [get]
func Get(ctx context.Contexer) (int, interface{}) {
	group := ctx.FindParam("group")
	fileName := ctx.FindParam("file")

	res, err := core.FindCachedAsset(group, fileName)

	if err != nil {
		return http.StatusNotFound, err
	}

	//mimes := make(map[string]string)
	//mimes["js"] = "text/javascript"
	//mimes["css"] = "text/css"
	//mimes["html"] = "text/html"
	//mimes["ico"] = "image/x-icon"
	//mimes["font"] = "font/" + getExt(fileName)

	return http.StatusOK, res
}

// @Title GetAsset List
// @Description Lists the assets in a group
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group [get]
func GetAll(ctx context.Contexer) (int, interface{}) {
	group := ctx.FindParam("group")
	assets, err := core.ListCachedAssets(group)

	if err != nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, assets
}
