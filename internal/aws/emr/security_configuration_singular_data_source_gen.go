// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package emr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_emr_security_configuration", securityConfigurationDataSource)
}

// securityConfigurationDataSource returns the Terraform awscc_emr_security_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::EMR::SecurityConfiguration resource.
func securityConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the security configuration.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the security configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The security configuration details in JSON format.",
		//	  "type": "string"
		//	}
		"security_configuration": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The security configuration details in JSON format.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EMR::SecurityConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMR::SecurityConfiguration").WithTerraformTypeName("awscc_emr_security_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"name":                   "Name",
		"security_configuration": "SecurityConfiguration",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
