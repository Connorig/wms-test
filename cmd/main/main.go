package main

import (
	"context"
	"github.com/Domingor/go-blackbox"
	"github.com/Domingor/go-blackbox/server/apploader"
	"github.com/Domingor/go-blackbox/version"
	"github.com/Domingor/wms-test/config"
)

/**
* @Author: Connor
* @Date:   24.6.7 14:10
* @Description:
 */

func main() {
	var config config.Config
	version.AppName = "wms-test"
	version.Version = "v1.1.0"

	debugLevel := "debug"
	port := ":8811"
	timeFormate := "2006-01-02 15:04:05"

	appbox.New().Start(func(ctx context.Context, builder *appbox.ApplicationBuild) error {

		builder.LoadConfig(&config, func(loader apploader.Loader) {
			loader.SetConfigFileSearcher("config", ".")
		})

		builder.
			InitLog(".", debugLevel).
			EnableWeb(timeFormate, port, debugLevel, nil)
		return nil
	})
}
