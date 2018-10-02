package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

// ListConfig = Print all config to screen
func ListConfig() {
	files, err := ioutil.ReadDir(PathName)
	if err != nil {
		fmt.Println("App not initialized!")
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
	}
}

// AddConfig
func AddConfig(c *cli.Context) {
	argPath := c.Args().Get(1)
	argName := c.Args().Get(0)

	if argPath != "" {
		if _, err := os.Stat("./" + argPath); os.IsNotExist(err) {
			log.Fatal("Path not found!")
		} else {
			fmt.Println("Path found")
		}
	} else {
		fmt.Println("Default config.json")
		CopyFile("config.json", PathName+"/"+argName+".json")
	}
}

// ApplyConfig
func ApplyConfig(c *cli.Context) {
	argName := c.Args().Get(0)

	if CheckIfConfigExist(argName) {
		CopyFile(PathName+"/"+argName+".json", "config.json")
	} else {
		log.Fatal("Config not found!")
	}
}
