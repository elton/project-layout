package config

import "testing"

func TestReadConfig(t *testing.T) {
	if err := ReadConfig(); err != nil {
		t.Errorf("ReadConfig() error = %v", err)
	}
	t.Logf("Server Name = %s", AppCfg.Server.Name)
	t.Logf("Server Port = %s", AppCfg.Server.Port)
}
