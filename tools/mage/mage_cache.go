//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
)

type Cache mg.Namespace

// // Get get from cache Usage "cache:get KEY"
// func (Cache) Get(key string) error {
// 	if key == "" {
// 		return fmt.Errorf("key cannot be empty e.g. cache get --key=foo")
// 	}
// 	fmt.Println("Cache get", key)
// 	return nil
// }

// // Set sets to cache Usage "cache:set KEY VALUE"
// func (Cache) Set(key string, value string) error {
// 	fmt.Println("Cache set", key, value)

// 	return nil
// }
