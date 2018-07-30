package main

import (
	"github.com/objque/pslytics-api/pkg/api"
	"github.com/objque/pslytics-api/pkg/config"
	"github.com/objque/pslytics-api/pkg/db"
	"github.com/objque/pslytics-api/pkg/log"
)

func main() {
	// TODO (m.kalinin): replace it with a consul or external cfg
	config.Config = &config.AppConfig{
		DB: config.DBConfig{
			DBType:  "mysql",
			DBHost:  "mariadb",
			DBLogin: "root",
			DBPass:  "pslytics",
			DBName:  "pslytics",
			Log:     false,
		},
		Log: config.LogConfig{
			File:          "pslytics.log",
			Level:         "INFO",
			SyslogEnabled: false,
		},
		HTTP: config.HTTPConfig{
			IP:   "",
			Port: 5110,
		},
		Fetching: config.Fetching{
			CountOfSkippedHoursToFetch: 8,
		},
		ProxyURL: "http://172.17.0.3:3310",
	}

	log.SetLogFormatter(&log.DefaultFormatter)
	log.ConfigureStdLogger(config.Config.Log.Level)
	db.DbMgr = db.NewMainDatabaseMgr()

	// go fetcher.Run()
	panic(api.StartAPIServer(config.Config.HTTP.IP, config.Config.HTTP.Port))
}
