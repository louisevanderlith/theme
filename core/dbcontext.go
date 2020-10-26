package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"reflect"
)

type context struct {
	Assets husk.Table
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
	//if ctx.Assets.Exists(op.Everything()) {
	//	ctx.Assets
	//}

	files, err := assetSeeds()

	if err != nil {
		panic(err)
		return
	}

	err = ctx.Assets.Seed(files)

	if err != nil {
		panic(err)
	}
}

func assetSeeds() (collections.Enumerable, error) {
	groups := []string{"css", "fonts", "html", "ico", "js"}

	//Find Files per Group
	var many []Asset

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

	return collections.ReadOnlyList(reflect.ValueOf(many)), nil
}
