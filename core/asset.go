package core

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//FindAsset finds the requested asset
func FindAsset(group, name string) ([]byte, error) {
	if len(group) < 2 {
		return nil, errors.New("group too short")
	}

	if len(name) < 3 || !strings.Contains(name, ".") {
		return nil, errors.New("name is invalid")
	}

	fullPath := filepath.Join("dist", group, name)
	return ioutil.ReadFile(fullPath)
}
