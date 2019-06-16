package models

import (
	"github.com/go-ini/ini"
)

var cfg *ini.File
var Env string

func init() {
	cfg, _ = ini.Load("conf/database.ini", "conf/app.ini")
}

func GetConfig(key string) string {
	return cfg.Section(Env).Key(key).String()
}
