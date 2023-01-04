// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lightsail_load_balancer", loadBalancerDataSource)
}

// loadBalancerDataSource returns the Terraform awscc_lightsail_load_balancer data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lightsail::LoadBalancer resource.
func loadBalancerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachedInstances
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The names of the instances attached to the load balancer.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"attached_instances": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The names of the instances attached to the load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., \"/\").",
		//	  "type": "string"
		//	}
		"health_check_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path you provided to perform the load balancer health check. If you didn't specify a health check path, Lightsail uses the root path of your website (e.g., \"/\").",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstancePort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The instance port where you're creating your load balancer.",
		//	  "type": "integer"
		//	}
		"instance_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The instance port where you're creating your load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpAddressType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.",
		//	  "type": "string"
		//	}
		"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address type for the load balancer. The possible values are ipv4 for IPv4 only, and dualstack for IPv4 and IPv6. The default value is dualstack.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"load_balancer_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of your load balancer.",
		//	  "pattern": "\\w[\\w\\-]*\\w",
		//	  "type": "string"
		//	}
		"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of your load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SessionStickinessEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Configuration option to enable session stickiness.",
		//	  "type": "boolean"
		//	}
		"session_stickiness_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Configuration option to enable session stickiness.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SessionStickinessLBCookieDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Configuration option to adjust session stickiness cookie duration parameter.",
		//	  "type": "string"
		//	}
		"session_stickiness_lb_cookie_duration_seconds": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Configuration option to adjust session stickiness cookie duration parameter.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TlsPolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the TLS policy to apply to the load balancer.",
		//	  "type": "string"
		//	}
		"tls_policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the TLS policy to apply to the load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lightsail::LoadBalancer",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::LoadBalancer").WithTerraformTypeName("awscc_lightsail_load_balancer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attached_instances":         "AttachedInstances",
		"health_check_path":          "HealthCheckPath",
		"instance_port":              "InstancePort",
		"ip_address_type":            "IpAddressType",
		"key":                        "Key",
		"load_balancer_arn":          "LoadBalancerArn",
		"load_balancer_name":         "LoadBalancerName",
		"session_stickiness_enabled": "SessionStickinessEnabled",
		"session_stickiness_lb_cookie_duration_seconds": "SessionStickinessLBCookieDurationSeconds",
		"tags":            "Tags",
		"tls_policy_name": "TlsPolicyName",
		"value":           "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
