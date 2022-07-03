package coingecko

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

type coinGeckoClient struct {
	HTTPClient *http.Client
	URL        string
}

func (c *coinGeckoClient) Get(subpath string) (*http.Response, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, subpath)
	return c.HTTPClient.Get(u.String())
}

func connect(_ context.Context, d *plugin.QueryData) (*coinGeckoClient, error) {
	cgConfig := GetConfig(d.Connection)
	duration := 30
	client := &http.Client{
		Timeout: time.Second * time.Duration(duration),
	}
	return &coinGeckoClient{
		URL:        cgConfig.URL,
		HTTPClient: client,
	}, nil
}
