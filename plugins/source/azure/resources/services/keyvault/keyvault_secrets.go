package keyvault

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func keyvault_secrets() *schema.Table {
	return &schema.Table{
		Name:      "azure_keyvault_keyvault_secrets",
		Resolver:  fetchKeyvaultSecrets,
		Transform: transformers.TransformWithStruct(&armkeyvault.Secret{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
