package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/theme/core"
)

type AssetController struct {
	xontrols.APICtrl
}

// @Title GetAsset
// @Description Gets the requested asset
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Param	file		path 	string	true		"full filename /html/master.entry.html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group/:file [get]
func (req *AssetController) Get() {
	group := req.FindParam("group")
	fileName := req.FindParam("file")

	res, err := core.FindAsset(group, fileName)

	if err != nil {
		log.Println(err)
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	mimes := make(map[string]string)
	mimes["js"] = "text/javascript"
	mimes["css"] = "text/css"
	mimes["html"] = "text/html"
	mimes["ico"] = "image/x-icon"
	mimes["font"] = "font/" + getExt(fileName)

	req.ServeBinaryWithMIME(res, fileName, mimes[group])
}

// @Title GetAsset List
// @Description Lists the assets in a group
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group [get]
func (req *AssetController) GetAll() {
	group := req.FindParam("group")
	assests, err := core.ListAssets(group)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, assests)
}

func getExt(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	return filename[dotIndex+1:]
}
