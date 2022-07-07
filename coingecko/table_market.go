package coingecko

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type Market struct {
	ID                           string  `json:"id"`
	CurrentPrice                 float32 `json:"current_price"`
	MarketCap                    int     `json:"market_cap"`
	MarketCapRank                int     `json:"market_cap_rank"`
	FullyDilutedValuation        int     `json:"fully_diluted_valuation"`
	TotalVolume                  float32 `json:"total_volume"`
	High24h                      float32 `json:"high_24h"`
	Low24h                       float32 `json:"low_24h"`
	PriceChange24h               float32 `json:"price_change_24h"`
	PriceChangePercentage24h     float32 `json:"price_change_percentage_24h"`
	MarketCapChange24h           float32 `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float32 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float32 `json:"circulating_supply"`
	TotalSupply                  float32 `json:"total_supply"`
	MaxSupply                    float32 `json:"max_supply"`
	ATH                          float32 `json:"ath"`
	ATHChangePercentage          float32 `json:"ath_change_percentage"`
}

func tableMarket(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "coingecko_market",
		Description: "all the coins market data (price, market cap, volume)",
		List: &plugin.ListConfig{
			//KeyColumns: plugin.OptionalColumns([]string{"id"}),
			KeyColumns: []*plugin.KeyColumn{
				{Name: "id", Require: plugin.Optional, Operators: []string{"="}},
			},
			Hydrate: listMarket,
		},
		//Get: &plugin.GetConfig{
		//	KeyColumns: plugin.SingleColumn("id"),
		//	Hydrate:    getMarket,
		//},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Coin ID"},
			{Name: "current_price", Type: proto.ColumnType_DOUBLE, Description: "Current price"},
			{Name: "market_cap", Type: proto.ColumnType_INT, Description: "Market cap"},
			{Name: "market_cap_rank", Type: proto.ColumnType_INT, Description: "Market cap rank"},
			{Name: "fully_diluted_valuation", Type: proto.ColumnType_DOUBLE, Description: "Fully diluated valuation"},
			{Name: "total_volume", Type: proto.ColumnType_INT, Description: "Total volume"},
			{Name: "high_24h", Type: proto.ColumnType_DOUBLE, Description: "High 24h"},
			{Name: "low_24h", Type: proto.ColumnType_DOUBLE, Description: "Low 24h"},
			{Name: "price_change_24h", Type: proto.ColumnType_DOUBLE, Description: "Price change 24h"},
			{Name: "price_change_percentage_24h", Type: proto.ColumnType_DOUBLE, Description: "Price change percentage 24h"},
			{Name: "market_cap_change_24h", Type: proto.ColumnType_DOUBLE, Description: "Market cap change 24h"},
			{Name: "market_cap_change_percentage_24h", Type: proto.ColumnType_DOUBLE, Description: "Market cap change percentage 24h"},
			{Name: "circulating_supply", Type: proto.ColumnType_DOUBLE, Description: "Circulating supply"},
			{Name: "total_supply", Type: proto.ColumnType_DOUBLE, Description: "Total supply"},
			{Name: "max_supply", Type: proto.ColumnType_DOUBLE, Description: "Max supply"},
			{Name: "ath", Type: proto.ColumnType_DOUBLE, Description: "ATH"},
			{Name: "ath_percentage_change", Type: proto.ColumnType_DOUBLE, Description: "ATH percentage change"},
		},
	}
}

func listMarket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cgClient, err := connect(ctx, d)
	if err != nil {
		log.Fatal(err)
	}
	var markets []Market

	var name string
	equalQuals := d.KeyColumnQuals
	//f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//defer f.Close()
	//f.WriteString("test")
	//for k := range d.Quals {
	//	f.WriteString(k + "\n")
	//}
	if equalQuals["id"] != nil {
		name = equalQuals["id"].GetStringValue()
		os.WriteFile("/tmp/dat3", []byte(name), 0644)
	}
	pageSize := 250
	page := 1
	for true {
		url := fmt.Sprintf("/coins/markets?vs_currency=usd&per_page=%d&page=%d", pageSize, page)
		err = cgClient.Get(url, &markets)
		if err != nil {
			plugin.Logger(ctx).Error("coingecko.listMarket", "query_error", err)
			return nil, err
		}
		if page > 3 {
			break
		}
		for _, market := range markets {
			d.StreamListItem(ctx, market)
		}
		page++
	}

	return nil, nil
}

//func getMarket(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
//	cgClient, err := connect(ctx, d)
//	if err != nil {
//		log.Fatal(err)
//	}
//	equalQuals := d.KeyColumnQuals
//	var name string
//	if equalQuals["id"] != nil {
//		name = equalQuals["id"].GetStringValue()
//		os.WriteFile("/tmp/dat2", []byte(name), 0644)
//	}
//	markets := []Market{}
//	url := fmt.Sprintf("/coins/markets?vs_currency=usd&ids=%s", name)
//	os.WriteFile("/tmp/dat1", []byte(url), 0644)
//	err = cgClient.Get(url, &markets)
//
//	if err != nil {
//		plugin.Logger(ctx).Error("coingecko.getCoin", "query_error", err)
//		return nil, err
//	}
//
//	return markets, nil
//
//}
