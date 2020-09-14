package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type assetFilter func(obj Asset) bool

func (f assetFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Asset))
}

func byGroupAndName(group, name string) assetFilter {
	return func(obj Asset) bool {
		return obj.Group == group && obj.Name == name
	}
}

func byGroup(group string) assetFilter {
	return func(obj Asset) bool {
		return obj.Group == group
	}
}
