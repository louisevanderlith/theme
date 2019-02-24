package core

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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