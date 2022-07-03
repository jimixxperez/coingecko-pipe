package coingecko

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-coingecko",
		DefaultTransform: transform.FromGo(),
		TableMap: map[string]*plugin.Table{
			"coingecko_coin": tableCoin(ctx),
			//"getml_pipeline": tablePipeline(ctx),

		},
	}
	return p
}
