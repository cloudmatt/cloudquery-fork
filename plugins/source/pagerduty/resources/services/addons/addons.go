// Code generated by codegen; DO NOT EDIT.

package addons

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Addons() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_addons",
		Description: `https://developer.pagerduty.com/api-reference/e58b140202a57-list-installed-add-ons`,
		Resolver:    fetchAddons,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Self"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "src",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Src"),
			},
			{
				Name:     "services",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Services"),
			},
		},
	}
}
