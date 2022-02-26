package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/elton/project-layout/pkg/logger"
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

// DatabaseCfg present the configuration of database.
type DatabaseCfg struct {
	Dsn string `yaml:"dsn"`
}

// Cfg represents the configuration of application.
type Cfg struct {
	Server   ServerCfg   `yaml:"server"`
	Database DatabaseCfg `yaml:"database"`
}

// AppCfg represents the configuration of application.
var AppCfg *Cfg

// ReadConfig reads the configuration file.
func ReadConfig(cfgMap map[string]string) error {
	var filePath string
	if os.Getenv("APP_ENV") == "test" { // in test mode.
		_, file, _, ok := runtime.Caller(0)
		if !ok {
			fmt.Fprintf(os.Stderr, "Unable to identify current directory")
			os.Exit(1)
		}

		// return the root of the project.
		basepath := filepath.Dir(filepath.Dir(file))
		filePath = filepath.Join(basepath, cfgMap[os.Getenv("APP_ENV")])

	} else {
		filePath = filepath.Join("./", cfgMap[os.Getenv("APP_ENV")])
	}

	f, err := os.Open(filePath)
	if err != nil {
		logger.Sugar.Errorf("Unable to open config file: %s", filePath)
		return err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&AppCfg); err != nil {
		logger.Sugar.Errorf("Unable to decode config file: %s", filePath)
		return err
	}
	logger.Sugar.Infof("Successfully loaded config file: %s", filePath)
	return nil
}
