package config

import (
	"log"
	"want_to_go/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

// 外から読み取れるようにGolobalに宣言
var Config ConfigList

// main関数より前に読み込む
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

// iniファイルを読み込む
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("5000"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
