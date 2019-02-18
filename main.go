package main

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/theme/routers"

	"github.com/astaxie/beego"
)

func main() {
	mode := beego.BConfig.RunMode

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, enums.API)

	routers.Setup(srv)
	beego.Run()
}
