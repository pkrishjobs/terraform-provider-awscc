// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package memorydb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_memorydb_subnet_group", subnetGroupDataSource)
}

// subnetGroupDataSource returns the Terraform awscc_memorydb_subnet_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::MemoryDB::SubnetGroup resource.
func subnetGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the subnet group.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the subnet group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional description of the subnet group.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An optional description of the subnet group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.",
		//	  "pattern": "[a-z][a-z0-9\\-]*",
		//	  "type": "string"
		//	}
		"subnet_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of VPC subnet IDs for the subnet group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of VPC subnet IDs for the subnet group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this subnet group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for the tag. May not be null.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value. May be null.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for the tag. May not be null.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value. May be null.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this subnet group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MemoryDB::SubnetGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::SubnetGroup").WithTerraformTypeName("awscc_memorydb_subnet_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "ARN",
		"description":       "Description",
		"key":               "Key",
		"subnet_group_name": "SubnetGroupName",
		"subnet_ids":        "SubnetIds",
		"tags":              "Tags",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
