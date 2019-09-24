package cli

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/logger"
	"net"
)

// Config is the global configuration varaible
var Config = viper.New()

// InitConfig set configuration
func InitConfig() {
	Config.SetConfigName("boxmetrics")
	Config.SetEnvPrefix("boxmetrics")
	Config.AutomaticEnv()

	Config.AddConfigPath("/etc/boxmetrics/")
	Config.AddConfigPath("$HOME/.boxmetrics/")
	Config.AddConfigPath("./configs/")
	Config.AddConfigPath(".")
	setDefault()

	err := Config.ReadInConfig()
	initLogger()

	if err != nil {
		Log.Warn("no configuration file found")
	} else {
		Log.Info("configuration file loaded")
		Config.WatchConfig()
		Config.OnConfigChange(func(e fsnotify.Event) {
			Log.Info("configuration file changed:", e.Name)
		})
	}
}

func setDefault() {
	consoleLog := Logger{Type: "console", Format: "text", Level: "debug"}
	Config.SetDefault("loggers", []Logger{consoleLog})
}
