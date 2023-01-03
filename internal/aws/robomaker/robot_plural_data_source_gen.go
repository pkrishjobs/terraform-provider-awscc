// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_robomaker_robots", robotsDataSource)
}

// robotsDataSource returns the Terraform awscc_robomaker_robots data source.
// This Terraform data source corresponds to the CloudFormation AWS::RoboMaker::Robot resource.
func robotsDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Description: "Uniquely identifies the data source.",
			Computed:    true,
		},
		"ids": schema.SetAttribute{
			Description: "Set of Resource Identifiers.",
			ElementType: types.StringType,
			Computed:    true,
		},
	}

	schema := schema.Schema{
		Description: "Plural Data Source schema for AWS::RoboMaker::Robot",
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::Robot").WithTerraformTypeName("awscc_robomaker_robots")
	opts = opts.WithTerraformSchema(schema)

	v, err := NewPluralDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
