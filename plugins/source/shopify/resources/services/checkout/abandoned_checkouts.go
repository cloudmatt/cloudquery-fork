// Code generated by codegen; DO NOT EDIT.

package checkout

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func AbandonedCheckouts() *schema.Table {
	return &schema.Table{
		Name:     "shopify_abandoned_checkouts",
		Resolver: fetchAbandonedCheckouts,
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
				Name:     "token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Token"),
			},
			{
				Name:     "cart_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CartToken"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "gateway",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Gateway"),
			},
			{
				Name:     "buyer_accepts_marketing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BuyerAcceptsMarketing"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "landing_site",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LandingSite"),
			},
			{
				Name:     "note_attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NoteAttributes"),
			},
			{
				Name:     "referring_site",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReferringSite"),
			},
			{
				Name:     "shipping_lines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ShippingLines"),
			},
			{
				Name:     "taxes_included",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TaxesIncluded"),
			},
			{
				Name:     "total_weight",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalWeight"),
			},
			{
				Name:     "currency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Currency"),
			},
			{
				Name:     "completed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CompletedAt"),
			},
			{
				Name:     "closed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClosedAt"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UserID"),
			},
			{
				Name:     "customer_locale",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerLocale"),
			},
			{
				Name:     "line_items",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LineItems"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "abandoned_checkout_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AbandonedCheckoutURL"),
			},
			{
				Name:     "discount_codes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DiscountCodes"),
			},
			{
				Name:     "tax_lines",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TaxLines"),
			},
			{
				Name:     "source_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceName"),
			},
			{
				Name:     "presentment_currency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PresentmentCurrency"),
			},
			{
				Name:     "buyer_accepts_sms_marketing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BuyerAcceptsSmsMarketing"),
			},
			{
				Name:     "total_discounts",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalDiscounts"),
			},
			{
				Name:     "total_line_items_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalLineItemsPrice"),
			},
			{
				Name:     "total_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalPrice"),
			},
			{
				Name:     "total_tax",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TotalTax"),
			},
			{
				Name:     "subtotal_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubtotalPrice"),
			},
			{
				Name:     "customer",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Customer"),
			},
		},
	}
}
