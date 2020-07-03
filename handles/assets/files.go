package assets

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/theme/core"
)

// Download
func Download(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	group := ctx.FindParam("group")
	fileName := ctx.FindParam("file")
	
	res, err := core.FindCachedAsset(group, fileName)

	if err != nil {
		log.Println("FindCacheAsset Error", err)
		http.Error(w, "", http.StatusBadRequest)
	}

	err = ctx.Serve(http.StatusOK, mix.Octet(fileName, res))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// View - all
func View(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	group := ctx.FindParam("group")
	assets, err := core.ListCachedAssets(group)

	if err != nil {
		log.Println("View Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(assets))

	if err != nil {
		log.Println(err)
	}
}
