package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/theme/controllers/assets"
)

func Setup(e resins.Epoxi) {
	routr := e.Router().(*mux.Router)
	e.JoinPath(routr, "/asset/{group:[a-z]+}", "Assets by Group", http.MethodGet, roletype.Unknown, mix.JSON, assets.GetAll)
	e.JoinPath(routr, "/asset/{group:[a-z]+}/{file}", "Download Assets", http.MethodGet, roletype.Unknown, mix.Octet, assets.Get)
}
