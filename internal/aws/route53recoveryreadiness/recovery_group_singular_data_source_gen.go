// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53recoveryreadiness_recovery_group", recoveryGroupDataSource)
}

// recoveryGroupDataSource returns the Terraform awscc_route53recoveryreadiness_recovery_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53RecoveryReadiness::RecoveryGroup resource.
func recoveryGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Cells
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "type": "array"
		//	}
		"cells": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of the cell Amazon Resource Names (ARNs) in the recovery group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RecoveryGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"recovery_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A collection of tags associated with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RecoveryGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the recovery group to create.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_]+",
		//	  "type": "string"
		//	}
		"recovery_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the recovery group to create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53RecoveryReadiness::RecoveryGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::RecoveryGroup").WithTerraformTypeName("awscc_route53recoveryreadiness_recovery_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cells":               "Cells",
		"key":                 "Key",
		"recovery_group_arn":  "RecoveryGroupArn",
		"recovery_group_name": "RecoveryGroupName",
		"tags":                "Tags",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
