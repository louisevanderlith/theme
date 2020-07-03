package core

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/louisevanderlith/husk"
)

type Asset struct {
	Group string `hsk:"size(5)"`
	Name  string `hsk:"size(128)"`
	BLOB  []byte `json:"-"` //Blob shouldn't be returned in JSON result sets.
}

func (a Asset) Valid() error {
	if len(a.Group) < 2 {
		return errors.New("group too short")
	}

	if len(a.Name) < 3 || !strings.Contains(a.Name, ".") {
		return errors.New("name is invalid")
	}

	return husk.ValidateStruct(&a)
}

func FindCachedAsset(group, name string) (io.Reader, error) {
	if len(group) < 2 {
		return nil, errors.New("group too short")
	}

	if len(name) < 3 || !strings.Contains(name, ".") {
		return nil, errors.New("name is invalid")
	}

	upload, err := ctx.Assets.FindFirst(byGroupAndName(group, name))

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(upload.Data().(Asset).BLOB), nil
}

func ListCachedAssets(group string) ([]string, error) {
	coll, err := ctx.Assets.Find(1, 100, byGroup(group))

	if err != nil {
		return nil, err
	}

	if !coll.Any() {
		return nil, errors.New("nothing found")
	}

	enumer := coll.GetEnumerator()

	var result []string
	for enumer.MoveNext() {
		obj := enumer.Current().Data().(Asset)
		result = append(result, obj.Name)
	}

	return result, nil
}

//ListAssets returns a collection of files in a group
func ListAssets(group string) ([]string, error) {
	var result []string

	if len(group) < 2 {
		return nil, errors.New("group too short")
	}

	fullPath := filepath.Join("dist", group)

	info, err := ioutil.ReadDir(fullPath)

	if err != nil {
		return result, err
	}

	for _, v := range info {
		result = append(result, v.Name())
	}

	return result, nil
}

//FindAsset finds the requested asset
func FindAsset(group, name string) ([]byte, error) {
	if len(group) < 2 {
		return nil, errors.New("group too short")
	}

	if len(name) < 3 || !strings.Contains(name, ".") {
		return nil, errors.New("name is invalid")
	}

	fullPath := filepath.Join("dist", group, name)

	bits, err := ioutil.ReadFile(fullPath)

	if err != nil {
		return nil, err
	}

	return bits, nil
}
