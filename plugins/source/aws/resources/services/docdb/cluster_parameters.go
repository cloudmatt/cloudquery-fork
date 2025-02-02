// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterParameters() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_cluster_parameters",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Parameter.html`,
		Resolver:    fetchDocdbClusterParameters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
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
				Name:     "allowed_values",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllowedValues"),
			},
			{
				Name:     "apply_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplyMethod"),
			},
			{
				Name:     "apply_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplyType"),
			},
			{
				Name:     "data_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataType"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "is_modifiable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsModifiable"),
			},
			{
				Name:     "minimum_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumEngineVersion"),
			},
			{
				Name:     "parameter_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterName"),
			},
			{
				Name:     "parameter_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterValue"),
			},
			{
				Name:     "source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Source"),
			},
		},
	}
}
