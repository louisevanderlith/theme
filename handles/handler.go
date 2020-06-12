package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/theme/handles/assets"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	view := kong.ResourceMiddleware("theme.assets.view", scrt, secureUrl, assets.View)
	r.HandleFunc("/asset/{group:[a-z]+}", view).Methods(http.MethodGet)

	//dwnld := kong.ResourceMiddleware("theme.assets.download", scrt, authUrl, assets.Download)
	r.HandleFunc("/asset/{group:[a-z]+}/{file}", assets.Download).Methods(http.MethodGet)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "theme.assets.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst, //you service is available and allowed for this base url
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
