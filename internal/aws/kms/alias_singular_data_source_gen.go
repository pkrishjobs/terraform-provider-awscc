// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_kms_alias", aliasDataSource)
}

// aliasDataSource returns the Terraform awscc_kms_alias data source.
// This Terraform data source corresponds to the CloudFormation AWS::KMS::Alias resource.
func aliasDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AliasName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^(alias/)[a-zA-Z0-9:/_-]+$",
		//	  "type": "string"
		//	}
		"alias_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed keys.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifies the AWS KMS key to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"target_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifies the AWS KMS key to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::KMS::Alias",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::KMS::Alias").WithTerraformTypeName("awscc_kms_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias_name":    "AliasName",
		"target_key_id": "TargetKeyId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
