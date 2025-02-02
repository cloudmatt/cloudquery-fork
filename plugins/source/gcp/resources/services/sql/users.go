// Code generated by codegen; DO NOT EDIT.

package sql

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "gcp_sql_users",
		Description: `https://cloud.google.com/sql/docs/mysql/admin-api/rest/v1beta4/users#User`,
		Resolver:    fetchUsers,
		Multiplex:   client.ProjectMultiplexEnabledServices("sqladmin.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "dual_password_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DualPasswordType"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "host",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Host"),
			},
			{
				Name:     "instance",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Instance"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "password",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Password"),
			},
			{
				Name:     "password_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PasswordPolicy"),
			},
			{
				Name:     "project",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Project"),
			},
			{
				Name:     "sqlserver_user_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SqlserverUserDetails"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
