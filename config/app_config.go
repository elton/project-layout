package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

// ServerCfg present the configuration of server.
type ServerCfg struct {
	Name         string `yaml:"name"`
	Port         string `yaml:"port"`
	Prefork      bool   `yaml:"prefork"`
	ReadTimeout  uint   `yaml:"readTimeout"`
	WriteTimeout uint   `yaml:"writeTimeout"`
}

// Cfg represents the configuration of application.
type Cfg struct {
	Server ServerCfg `yaml:"server"`
}

// AppCfg represents the configuration of application.
var AppCfg *Cfg

// ReadConfig reads the configuration file.
func ReadConfig(cfgPath string) error {
	var filePath string
	if os.Getenv("APP_ENV") == "test" { // in test mode.
		_, file, _, ok := runtime.Caller(0)
		if !ok {
			fmt.Fprintf(os.Stderr, "Unable to identify current directory")
			os.Exit(1)
		}

		// return the root of the project.
		basepath := filepath.Dir(filepath.Dir(file))
		filePath = filepath.Join(basepath, cfgPath)

	} else {
		filePath = filepath.Join("./", cfgPath)
	}
	// fmt.Println("filePath: ", filePath)

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
