package coingecko

import "github.com/turbot/steampipe-plugin-sdk/v3/plugin"

type CoinGeckoConfig struct {
	URL        string
	VsCurrency string
}

func GetConfig(connection *plugin.Connection) CoinGeckoConfig {
	return CoinGeckoConfig{URL: "https://api.coingecko.com/api/v3", VsCurrency: "usd"}

}
