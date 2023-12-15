//go:build mage
// +build mage

// This project uses mage (https://magefile.org/) to run tasks.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
)

// read config from file
func ReadMageConfig() {
	// if configDir empty use default "$PWD/internal/config"
	// absoluteDir := pwd + "/../../" + "/internal/config"
	// get current working directory
	pwd, _ := os.Getwd()
	// path join
	absoluteDir := filepath.Join(pwd, "/../../", "/internal/config")

	// load and parse application envs
	config.ReadConfig(absoluteDir, false)
}

// Echo prints input Usage "mage echo MESSAGE"
func Echo(message string) {
	fmt.Println(message)
}
