// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_apigateway_request_validator", requestValidatorDataSourceType)
}

// requestValidatorDataSourceType returns the Terraform awscc_apigateway_request_validator data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ApiGateway::RequestValidator resource type.
func requestValidatorDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the request validator.",
			//   "type": "string"
			// }
			Description: "Name of the request validator.",
			Type:        types.StringType,
			Computed:    true,
		},
		"request_validator_id": {
			// Property: RequestValidatorId
			// CloudFormation resource type schema:
			// {
			//   "description": "ID of the request validator.",
			//   "type": "string"
			// }
			Description: "ID of the request validator.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			// {
			//   "description": "The identifier of the targeted API entity.",
			//   "type": "string"
			// }
			Description: "The identifier of the targeted API entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"validate_request_body": {
			// Property: ValidateRequestBody
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether to validate the request body according to the configured schema for the targeted API and method. ",
			//   "type": "boolean"
			// }
			Description: "Indicates whether to validate the request body according to the configured schema for the targeted API and method. ",
			Type:        types.BoolType,
			Computed:    true,
		},
		"validate_request_parameters": {
			// Property: ValidateRequestParameters
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether to validate request parameters.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether to validate request parameters.",
			Type:        types.BoolType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ApiGateway::RequestValidator",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::RequestValidator").WithTerraformTypeName("awscc_apigateway_request_validator")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"name":                        "Name",
		"request_validator_id":        "RequestValidatorId",
		"rest_api_id":                 "RestApiId",
		"validate_request_body":       "ValidateRequestBody",
		"validate_request_parameters": "ValidateRequestParameters",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_apigateway_request_validator", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}