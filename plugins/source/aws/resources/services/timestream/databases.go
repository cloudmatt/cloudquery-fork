// Code generated by codegen; DO NOT EDIT.

package timestream

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Databases() *schema.Table {
	return &schema.Table{
		Name:      "aws_timestream_databases",
		Resolver:  fetchTimestreamDatabases,
		Multiplex: client.ServiceAccountRegionMultiplexer("ingest.timestream"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: fetchDatabaseTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "database_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseName"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "table_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TableCount"),
			},
		},

		Relations: []*schema.Table{
			Tables(),
		},
	}
}
