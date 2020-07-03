package core

import (
	"github.com/louisevanderlith/husk"
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
	groups := []string{"css", "fonts", "html", "ico", "js"}

	//Find Files per Group
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

			rec := ctx.Assets.Create(obj)

			if rec.Error != nil {
				panic(rec.Error)
			}
		}

		ctx.Assets.Save()
	}
}
