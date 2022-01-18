// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_appstream_application_entitlement_association", applicationEntitlementAssociationDataSourceType)
}

// applicationEntitlementAssociationDataSourceType returns the Terraform awscc_appstream_application_entitlement_association data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::AppStream::ApplicationEntitlementAssociation resource type.
func applicationEntitlementAssociationDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_identifier": {
			// Property: ApplicationIdentifier
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"entitlement_name": {
			// Property: EntitlementName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"stack_name": {
			// Property: StackName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::AppStream::ApplicationEntitlementAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::ApplicationEntitlementAssociation").WithTerraformTypeName("awscc_appstream_application_entitlement_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_identifier": "ApplicationIdentifier",
		"entitlement_name":       "EntitlementName",
		"stack_name":             "StackName",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}