package hybridcompute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func PrivateLinkScopes() *schema.Table {
	return &schema.Table{
		Name:        "azure_hybridcompute_private_link_scopes",
		Resolver:    fetchPrivateLinkScopes,
		Description: "https://learn.microsoft.com/en-us/rest/api/hybridcompute/private-link-scopes/list?tabs=HTTP#hybridcomputeprivatelinkscope",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_hybridcompute_private_link_scopes", client.Namespacemicrosoft_hybridcompute),
		Transform:   transformers.TransformWithStruct(&armhybridcompute.PrivateLinkScope{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateLinkScopes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armhybridcompute.NewPrivateLinkScopesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
