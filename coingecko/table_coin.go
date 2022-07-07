package coingecko

import (
	"context"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

type Coin struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

func tableCoin(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "coingecko_coin",
		Description: "List all coins.",
		List: &plugin.ListConfig{
			Hydrate: listCoin,
		},
		//Get: &plugin.GetConfig{
		//	Hydrate:    getCoin,
		//	KeyColumns: plugin.SingleColumn("id"),
		//},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "Coin ID"},
			{Name: "symbol", Type: proto.ColumnType_STRING, Description: "Coin Symbol"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Coin name"},
			{Name: "platforms", Type: proto.ColumnType_JSON, Description: "Coin platforms"},
		},
	}
}

func listCoin(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	cgClient, err := connect(ctx, d)
	if err != nil {
		log.Fatal(err)
	}
	var coins []Coin
	err = cgClient.Get("/coins/list?include_platform=true", &coins)
	if err != nil {
		return nil, err
	}
	for _, coin := range coins {
		d.StreamListItem(ctx, coin)
	}
	return nil, nil
}

//func getCoin(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error){
//	cgClient, err := connect(ctx, d)
//	if err != nil {
//		log.Fatal(err)
//	}
//	equalQuals := d.KeyColumnQuals
//	var name string
//	if equalQuals["name"] = nil {
//		name = equalQuals["name"].GetStringValue()
//	}
//
//	obj := Coin{}
//	data, err := cgClient.Get(fmt.Sprintf("/coins/%s", name))
//
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(data, &obj)
//
//	if err != nil {
//		plugin.Logger(ctx).Error("coingecko.getCoin", "query_error", err)
//		return nil, err
//	}
//
//	return obj, nil
//
//}
