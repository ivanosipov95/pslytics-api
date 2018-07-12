package main

import (
	"github.com/objque/pslytics-api/pkg/api"
	"github.com/objque/pslytics-api/pkg/config"
	"github.com/objque/pslytics-api/pkg/db"
)

func main() {
	// TODO (m.kalinin): replace it with a consul or external cfg
	config.Config = &config.AppConfig{
		DB: config.DBConfig{
			DBType:  "mysql",
			DBHost:  "127.0.0.1",
			DBLogin: "pslytics",
			DBPass:  "pslytics",
			DBName:  "pslytics",
			Log:     true,
		},
		Log: config.LogConfig{
			File:          "pslytics.log",
			Level:         "INFO",
			SyslogEnabled: false,
		},
		HTTP: config.HTTPConfig{
			IP:   "127.0.0.1",
			Port: 5110,
		},
	}

	db.DbMgr = db.NewMainDatabaseMgr()
	panic(api.StartAPIServer(config.Config.HTTP.IP, config.Config.HTTP.Port))
}
