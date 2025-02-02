// Code generated by codegen; DO NOT EDIT.

package team

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func TeamMembers() *schema.Table {
	return &schema.Table{
		Name:     "vercel_team_members",
		Resolver: fetchTeamMembers,
		Columns: []schema.Column{
			{
				Name:     "team_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "avatar",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Avatar"),
			},
			{
				Name:     "confirmed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Confirmed"),
			},
			{
				Name:     "email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Email"),
			},
			{
				Name:     "github",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Github"),
			},
			{
				Name:     "gitlab",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Gitlab"),
			},
			{
				Name:     "bitbucket",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Bitbucket"),
			},
			{
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Role"),
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Username"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "access_requested_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccessRequestedAt"),
			},
		},
	}
}
