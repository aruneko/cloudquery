// Code generated by codegen2; DO NOT EDIT.
package compute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Skus() *schema.Table {
	return &schema.Table{
		Name:      "azure_compute_skus",
		Resolver:  fetchSkus,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Compute),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "api_versions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("APIVersions"),
			},
			{
				Name:     "capabilities",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capabilities"),
			},
			{
				Name:     "capacity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Capacity"),
			},
			{
				Name:     "costs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Costs"),
			},
			{
				Name:     "family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Family"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "location_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LocationInfo"),
			},
			{
				Name:     "locations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Locations"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "restrictions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Restrictions"),
			},
			{
				Name:     "size",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Tier"),
			},
		},
	}
}

func fetchSkus(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewResourceSKUsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
