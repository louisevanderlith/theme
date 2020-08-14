package assets

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	"github.com/louisevanderlith/theme/core"
)

// Download
func Download(w http.ResponseWriter, r *http.Request) {
	group := drx.FindParam(r, "group")
	fileName := drx.FindParam(r, "file")

	res, err := core.FindCachedAsset(group, fileName)

	if err != nil {
		log.Println("FindCacheAsset Error", err)
		http.Error(w, "", http.StatusBadRequest)
	}

	err = mix.Write(w, mix.Octet(fileName, res))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// View - all
func View(w http.ResponseWriter, r *http.Request) {
	group := drx.FindParam(r, "group")
	assets, err := core.ListCachedAssets(group)

	if err != nil {
		log.Println("View Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(assets))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
