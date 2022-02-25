package config

import (
	"testing"
	"time"
)

func TestReadConfig(t *testing.T) {

	const configPath = "/app/myapp/etc/server.yaml"
	if err := ReadConfig(configPath); err != nil {
		t.Errorf("ReadConfig(%s) failed: %v", configPath, err)
	}
	t.Logf("Server Name = %s", AppCfg.Server.Name)
	t.Logf("Server Port = %s", AppCfg.Server.Port)
	t.Logf("Server Prefork = %#v", AppCfg.Server.Prefork)
	t.Logf("Server ReadTimeout = %d", AppCfg.Server.ReadTimeout)
	t.Logf("Server WriteTimeout = %v", time.Duration(AppCfg.Server.WriteTimeout)*time.Second)
}
