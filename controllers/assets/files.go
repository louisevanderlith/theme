package assets

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/louisevanderlith/theme/core"
)

// @Title GetAsset
// @Description Gets the requested asset
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Param	file		path 	string	true		"full filename /html/master.entry.html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group/:file [get]
func Get(c *gin.Context) {
	group := c.Param("group")
	fileName := c.Param("file")

	res, conLen, err := core.FindCachedAsset(group, fileName)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	ext := getExt(fileName)
	mimes := make(map[string]string)
	mimes["js"] = "text/javascript"
	mimes["css"] = "text/css"
	mimes["html"] = "text/html"
	mimes["ico"] = "image/x-icon"
	mimes["font"] = "font/" + ext
	mimes["jpeg"] = "image/jpeg"
	mimes["jpg"] = "image/jpeg"
	mimes["png"] = "image/png"

	content := mimes[ext]

	c.DataFromReader(http.StatusOK, conLen, content, res, headers(content, fileName))
}

// @Title GetAsset List
// @Description Lists the assets in a group
// @Param	group			path	string 	true		"the group name like css, js, html"
// @Success 200 {string} string
// @Failure 403 :asstype or :file is empty
// @router /:group [get]
func GetAll(c *gin.Context) {
	group := c.Param("group")
	assets, err := core.ListCachedAssets(group)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, assets)
}

func headers(contenttype, filename string) map[string]string {
	result := make(map[string]string)

	result["Strict-Transport-Security"] = "max-age=31536000; includeSubDomains"
	result["Access-Control-Allow-Credentials"] = "true"
	result["Server"] = "kettle"
	result["X-Content-Type-Options"] = "nosniff"

	result["Content-Description"] = "File Transfer"
	result["Content-Transfer-Encoding"] = "binary"
	result["Expires"] = "0"
	result["Cache-Control"] = "must-revalidate"
	result["Pragma"] = "public"

	result["Content-Disposition"] = "attachment; filename=" + filename
	result["Content-Type"] = contenttype

	return result
}

func getExt(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	return filename[dotIndex+1:]
}
