package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("spider", map[string]any{
		"thread":   config.Env("SPIDER_THREAD", 30),
		"path":     config.Env("SPIDER_PATH", "download"),
		"live_api": config.Env("SPIDER_LIVE_URL", "http://api.hclyz.com:81"),
	})
}
