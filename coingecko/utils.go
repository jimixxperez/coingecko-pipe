package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type coinGeckoClient struct {
	HTTPClient *http.Client
	URL        string
}

func (c *coinGeckoClient) Get(subpath string, target interface{}) error {
	//u, err := url.Parse(c.URL)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//u.Path = path.Join(u.Path, subpath)
	//r, err := c.HTTPClient.Get(u.String())
	r, err := c.HTTPClient.Get(fmt.Sprintf("%s/%s", c.URL, subpath))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return json.Unmarshal(body, target)
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
