// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_documentation_part", documentationPartDataSource)
}

// documentationPartDataSource returns the Terraform awscc_apigateway_documentation_part data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::DocumentationPart resource.
func documentationPartDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DocumentationPartId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the documentation Part.",
		//	  "type": "string"
		//	}
		"documentation_part_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the documentation Part.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The location of the API entity that the documentation applies to.",
		//	  "properties": {
		//	    "Method": {
		//	      "description": "The HTTP verb of a method.",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "The name of the targeted API entity.",
		//	      "type": "string"
		//	    },
		//	    "Path": {
		//	      "description": "The URL path of the target.",
		//	      "type": "string"
		//	    },
		//	    "StatusCode": {
		//	      "description": "The HTTP status code of a response.",
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "description": "The type of API entity that the documentation content applies to.",
		//	      "enum": [
		//	        "API",
		//	        "AUTHORIZER",
		//	        "MODEL",
		//	        "RESOURCE",
		//	        "METHOD",
		//	        "PATH_PARAMETER",
		//	        "QUERY_PARAMETER",
		//	        "REQUEST_HEADER",
		//	        "REQUEST_BODY",
		//	        "RESPONSE",
		//	        "RESPONSE_HEADER",
		//	        "RESPONSE_BODY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Method
				"method": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTTP verb of a method.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the targeted API entity.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Path
				"path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The URL path of the target.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StatusCode
				"status_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTTP status code of a response.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of API entity that the documentation content applies to.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The location of the API entity that the documentation applies to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Properties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The documentation content map of the targeted API entity.",
		//	  "type": "string"
		//	}
		"properties": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The documentation content map of the targeted API entity.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RestApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifier of the targeted API entity",
		//	  "type": "string"
		//	}
		"rest_api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifier of the targeted API entity",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::DocumentationPart",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationPart").WithTerraformTypeName("awscc_apigateway_documentation_part")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"documentation_part_id": "DocumentationPartId",
		"location":              "Location",
		"method":                "Method",
		"name":                  "Name",
		"path":                  "Path",
		"properties":            "Properties",
		"rest_api_id":           "RestApiId",
		"status_code":           "StatusCode",
		"type":                  "Type",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
