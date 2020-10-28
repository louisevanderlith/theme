package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/louisevanderlith/theme/handles/assets"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()
	mw := open.BearerMiddleware(audience, issuer)

	r.Handle("/asset/{group:[a-z]+}", mw.Handler(http.HandlerFunc(assets.View))).Methods(http.MethodGet)

	//dwnld := ins.Middleware("theme.assets.download", scrt, authUrl, assets.Download)
	r.HandleFunc("/asset/{group:[a-z]+}/{file}", assets.Download).Methods(http.MethodGet)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "theme.assets.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
