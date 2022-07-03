package coingecko

import "github.com/turbot/steampipe-plugin-sdk/plugin"

type CoinGeckoConfig struct {
	URL        string
	VsCurrency string
}

func GetConfig(connection *plugin.Connection) CoinGeckoConfig {
	return CoinGeckoConfig{VsCurrency: "usd"}
}
