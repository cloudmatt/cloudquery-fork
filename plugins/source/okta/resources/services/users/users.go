// Code generated by codegen; DO NOT EDIT.

package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:     "okta_users",
		Resolver: fetchUsers,
		Columns: []schema.Column{
			{
				Name:     "activated",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveNullableTime("Activated"),
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "credentials",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Credentials"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "last_login",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveNullableTime("LastLogin"),
			},
			{
				Name:     "last_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdated"),
			},
			{
				Name:     "password_changed",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveNullableTime("PasswordChanged"),
			},
			{
				Name:     "profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Profile"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_changed",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveNullableTime("StatusChanged"),
			},
			{
				Name:     "transitioning_to_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitioningToStatus"),
			},
			{
				Name:     "type",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "_embedded",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Embedded"),
			},
			{
				Name:     "_links",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Links"),
			},
			{
				Name:     "additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalProperties"),
			},
		},
	}
}
