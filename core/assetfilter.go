package core

import (
	"github.com/louisevanderlith/husk"
)

type assetFilter func(obj *Asset) bool

func (f assetFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(*Asset))
}

func byGroupAndName(group, name string) assetFilter {
	return func(obj *Asset) bool {
		return obj.Group == group && obj.Name == name
	}
}

func byGroup(group string) assetFilter {
	return func(obj *Asset) bool {
		return obj.Group == group
	}
}
