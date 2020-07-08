package core

import (
	"github.com/louisevanderlith/husk"
	"log"
)

type context struct {
	Assets husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Assets: husk.NewTable(Asset{}),
	}
}

func Shutdown() {
	ctx.Assets.Save()
}

func seed() {
	if ctx.Assets.Exists(husk.Everything()) {
		return
	}

	groups := []string{"css", "fonts", "html", "ico", "js"}

	//Find Files per Group
	var many []husk.Dataer

	for _, group := range groups {
		assets, err := ListAssets(group)

		if err != nil {
			panic(err)
		}

		//Save file data
		for _, asset := range assets {
			data, err := FindAsset(group, asset)

			if err != nil {
				panic(err)
			}

			obj := Asset{
				BLOB:  data,
				Group: group,
				Name:  asset,
			}

			many = append(many, obj)
		}
	}

	tot, err := ctx.Assets.CreateMulti(many...)

	if err != nil {
		panic(err)
		return
	}

	log.Println("Rows Seeded", tot)

	err = ctx.Assets.Save()

	if err != nil {
		panic(err)
	}
}
