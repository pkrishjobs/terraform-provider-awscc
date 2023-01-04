// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_config_organization_conformance_pack", organizationConformancePackDataSource)
}

// organizationConformancePackDataSource returns the Terraform awscc_config_organization_conformance_pack data source.
// This Terraform data source corresponds to the CloudFormation AWS::Config::OrganizationConformancePack resource.
func organizationConformancePackDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConformancePackInputParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of ConformancePackInputParameter objects.",
		//	  "items": {
		//	    "description": "Input parameters in the form of key-value pairs for the conformance pack.",
		//	    "properties": {
		//	      "ParameterName": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "maxLength": 4096,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ParameterName",
		//	      "ParameterValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 60,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"conformance_pack_input_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ParameterName
					"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of ConformancePackInputParameter objects.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeliveryS3Bucket
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS Config stores intermediate files while processing conformance pack template.",
		//	  "maxLength": 63,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"delivery_s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AWS Config stores intermediate files while processing conformance pack template.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeliveryS3KeyPrefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The prefix for the delivery S3 bucket.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"delivery_s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The prefix for the delivery S3 bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ExcludedAccounts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1000,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"excluded_accounts": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OrganizationConformancePackName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the organization conformance pack.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z][-a-zA-Z0-9]*",
		//	  "type": "string"
		//	}
		"organization_conformance_pack_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the organization conformance pack.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string containing full conformance pack template body.",
		//	  "maxLength": 51200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"template_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string containing full conformance pack template body.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TemplateS3Uri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Location of file containing the template body.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": "s3://.*",
		//	  "type": "string"
		//	}
		"template_s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Location of file containing the template body.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Config::OrganizationConformancePack",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::OrganizationConformancePack").WithTerraformTypeName("awscc_config_organization_conformance_pack")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"conformance_pack_input_parameters":  "ConformancePackInputParameters",
		"delivery_s3_bucket":                 "DeliveryS3Bucket",
		"delivery_s3_key_prefix":             "DeliveryS3KeyPrefix",
		"excluded_accounts":                  "ExcludedAccounts",
		"organization_conformance_pack_name": "OrganizationConformancePackName",
		"parameter_name":                     "ParameterName",
		"parameter_value":                    "ParameterValue",
		"template_body":                      "TemplateBody",
		"template_s3_uri":                    "TemplateS3Uri",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
