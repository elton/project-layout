package configs

import (
	"testing"
	"time"

	"github.com/elton/project-layout/app/myapp/global"
)

func TestReadConfig(t *testing.T) {
	if err := ReadConfig(global.CfgMap); err != nil {
		t.Errorf("ReadConfig(%v) failed: %v", global.CfgMap, err)
	}
	t.Logf("Server Name = %s", AppCfg.Server.Name)
	t.Logf("Server Port = %s", AppCfg.Server.Port)
	t.Logf("Server Prefork = %#v", AppCfg.Server.Prefork)
	t.Logf("Server ReadTimeout = %d", AppCfg.Server.ReadTimeout)
	t.Logf("Server WriteTimeout = %v", time.Duration(AppCfg.Server.WriteTimeout)*time.Second)
	t.Logf("database connection string = %s", AppCfg.Database.Dsn)
}
