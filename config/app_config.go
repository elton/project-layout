package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

const configPath = "/app/myapp/etc/server.yaml"

type ServerCfg struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type Cfg struct {
	Server ServerCfg `yaml:"server"`
}

var AppCfg *Cfg

func ReadConfig() error {
	var filePath string
	if os.Getenv("APP_ENV") == "test" { // in test mode.
		_, file, _, ok := runtime.Caller(0)
		if !ok {
			fmt.Fprintf(os.Stderr, "Unable to identify current directory")
			os.Exit(1)
		}

		// return the root of the project.
		basepath := filepath.Dir(filepath.Dir(file))
		filePath = filepath.Join(basepath, configPath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppCfg)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
