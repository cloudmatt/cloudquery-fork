// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SshPublicKeys() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_ssh_public_keys",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SSHPublicKeyMetadata.html`,
		Resolver:    fetchIamSshPublicKeys,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "ssh_public_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SSHPublicKeyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "upload_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UploadDate"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
			},
		},
	}
}
