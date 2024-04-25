package main

import (
	"path/filepath"
	"strings"
)

func IsMedia(path string) bool {
	ext := Ext(path)

	if _, ok := mediaTypes[ext]; !ok {
		return false
	}

	return true
}

func Ext(path string) string {
	return strings.ToLower(filepath.Ext(path))
}
