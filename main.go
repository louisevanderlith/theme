package main

import (
	"os"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/theme/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := os.Getenv("RUNMODE")
	pubPath := os.Getenv("KEYPATH")

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		panic(err)
	}

	routers.Setup(srv)
	beego.Run()
}
