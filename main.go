package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/theme/routers"

	"github.com/astaxie/beego"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(appName, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		panic(err)
	}

	routers.Setup(srv, host)
	beego.Run()
}
