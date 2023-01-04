// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_personalize_schema", schemaDataSource)
}

// schemaDataSource returns the Terraform awscc_personalize_schema data source.
// This Terraform data source corresponds to the CloudFormation AWS::Personalize::Schema resource.
func schemaDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain of a Domain dataset group.",
		//	  "enum": [
		//	    "ECOMMERCE",
		//	    "VIDEO_ON_DEMAND"
		//	  ],
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain of a Domain dataset group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for the schema.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A schema in Avro JSON format.",
		//	  "maxLength": 10000,
		//	  "type": "string"
		//	}
		"schema": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A schema in Avro JSON format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn for the schema.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn for the schema.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Personalize::Schema",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Schema").WithTerraformTypeName("awscc_personalize_schema")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain":     "Domain",
		"name":       "Name",
		"schema":     "Schema",
		"schema_arn": "SchemaArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
