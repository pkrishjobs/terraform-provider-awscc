// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53_health_check", healthCheckDataSource)
}

// healthCheckDataSource returns the Terraform awscc_route53_health_check data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53::HealthCheck resource.
func healthCheckDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: HealthCheckConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A complex type that contains information about the health check.",
		//	  "properties": {
		//	    "AlarmIdentifier": {
		//	      "additionalProperties": false,
		//	      "description": "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
		//	      "properties": {
		//	        "Name": {
		//	          "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
		//	          "maxLength": 256,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "Region": {
		//	          "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Name",
		//	        "Region"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ChildHealthChecks": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "maxItems": 256,
		//	      "type": "array"
		//	    },
		//	    "EnableSNI": {
		//	      "type": "boolean"
		//	    },
		//	    "FailureThreshold": {
		//	      "maximum": 10,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "FullyQualifiedDomainName": {
		//	      "maxLength": 255,
		//	      "type": "string"
		//	    },
		//	    "HealthThreshold": {
		//	      "maximum": 256,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "IPAddress": {
		//	      "maxLength": 45,
		//	      "pattern": "^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))$|^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$",
		//	      "type": "string"
		//	    },
		//	    "InsufficientDataHealthStatus": {
		//	      "enum": [
		//	        "Healthy",
		//	        "LastKnownStatus",
		//	        "Unhealthy"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Inverted": {
		//	      "type": "boolean"
		//	    },
		//	    "MeasureLatency": {
		//	      "type": "boolean"
		//	    },
		//	    "Port": {
		//	      "maximum": 65535,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Regions": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "maxItems": 64,
		//	      "type": "array"
		//	    },
		//	    "RequestInterval": {
		//	      "maximum": 30,
		//	      "minimum": 10,
		//	      "type": "integer"
		//	    },
		//	    "ResourcePath": {
		//	      "maxLength": 255,
		//	      "type": "string"
		//	    },
		//	    "RoutingControlArn": {
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "SearchString": {
		//	      "maxLength": 255,
		//	      "type": "string"
		//	    },
		//	    "Type": {
		//	      "enum": [
		//	        "CALCULATED",
		//	        "CLOUDWATCH_METRIC",
		//	        "HTTP",
		//	        "HTTP_STR_MATCH",
		//	        "HTTPS",
		//	        "HTTPS_STR_MATCH",
		//	        "TCP",
		//	        "RECOVERY_CONTROL"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"health_check_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AlarmIdentifier
				"alarm_identifier": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Name
						"name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Region
						"region": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ChildHealthChecks
				"child_health_checks": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EnableSNI
				"enable_sni": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: FailureThreshold
				"failure_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: FullyQualifiedDomainName
				"fully_qualified_domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: HealthThreshold
				"health_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IPAddress
				"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: InsufficientDataHealthStatus
				"insufficient_data_health_status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Inverted
				"inverted": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MeasureLatency
				"measure_latency": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Regions
				"regions": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RequestInterval
				"request_interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ResourcePath
				"resource_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RoutingControlArn
				"routing_control_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SearchString
				"search_string": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A complex type that contains information about the health check.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"health_check_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag.",
		//	        "maxLength": 128,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"health_check_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53::HealthCheck",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::HealthCheck").WithTerraformTypeName("awscc_route53_health_check")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_identifier":                "AlarmIdentifier",
		"child_health_checks":             "ChildHealthChecks",
		"enable_sni":                      "EnableSNI",
		"failure_threshold":               "FailureThreshold",
		"fully_qualified_domain_name":     "FullyQualifiedDomainName",
		"health_check_config":             "HealthCheckConfig",
		"health_check_id":                 "HealthCheckId",
		"health_check_tags":               "HealthCheckTags",
		"health_threshold":                "HealthThreshold",
		"insufficient_data_health_status": "InsufficientDataHealthStatus",
		"inverted":                        "Inverted",
		"ip_address":                      "IPAddress",
		"key":                             "Key",
		"measure_latency":                 "MeasureLatency",
		"name":                            "Name",
		"port":                            "Port",
		"region":                          "Region",
		"regions":                         "Regions",
		"request_interval":                "RequestInterval",
		"resource_path":                   "ResourcePath",
		"routing_control_arn":             "RoutingControlArn",
		"search_string":                   "SearchString",
		"type":                            "Type",
		"value":                           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
