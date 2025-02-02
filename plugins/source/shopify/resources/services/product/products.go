// Code generated by codegen; DO NOT EDIT.

package product

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:     "shopify_products",
		Resolver: fetchProducts,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "title",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Title"),
			},
			{
				Name:     "body_html",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BodyHTML"),
			},
			{
				Name:     "vendor",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Vendor"),
			},
			{
				Name:     "product_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductType"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "handle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Handle"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "published_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PublishedAt"),
			},
			{
				Name:     "template_suffix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TemplateSuffix"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "published_scope",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublishedScope"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "image",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Image"),
			},
			{
				Name:     "admin_graphql_api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdminGraphqlAPIID"),
			},
		},

		Relations: []*schema.Table{
			ProductVariants(),
			ProductImages(),
		},
	}
}
