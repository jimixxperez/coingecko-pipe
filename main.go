package main

import (
	"github.com/jimixxperez/coingecko-pipe/coingecko"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: coingecko.Plugin})
}
