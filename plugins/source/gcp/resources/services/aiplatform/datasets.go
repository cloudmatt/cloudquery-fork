// Code generated by codegen; DO NOT EDIT.

package aiplatform

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"google.golang.org/genproto/googleapis/cloud/location"

	"cloud.google.com/go/aiplatform/apiv1"
)

func Datasets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_datasets",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.datasets#Dataset`,
		Resolver:    fetchDatasets,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "metadata_schema_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataSchemaUri"),
			},
			{
				Name:          "metadata",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Metadata"),
				IgnoreInTests: true,
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "encryption_spec",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionSpec"),
			},
			{
				Name:     "metadata_artifact",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataArtifact"),
			},
		},
	}
}

func fetchDatasets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListDatasetsRequest{
		Parent: parent.Item.(*location.Location).Name,
	}
	if filterLocation(parent) {
		return nil
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewDatasetClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListDatasets(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
