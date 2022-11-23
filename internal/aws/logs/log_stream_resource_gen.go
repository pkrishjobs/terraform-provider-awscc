// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_logs_log_stream", logStreamResource)
}

// logStreamResource returns the Terraform awscc_logs_log_stream resource.
// This Terraform resource corresponds to the CloudFormation AWS::Logs::LogStream resource.
func logStreamResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"log_group_name": {
			// Property: LogGroupName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the log group where the log stream is created.",
			//	  "type": "string"
			//	}
			Description: "The name of the log group where the log stream is created.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"log_stream_name": {
			// Property: LogStreamName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the log stream. The name must be unique wihtin the log group.",
			//	  "type": "string"
			//	}
			Description: "The name of the log stream. The name must be unique wihtin the log group.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Logs::LogStream",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::LogStream").WithTerraformTypeName("awscc_logs_log_stream")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"log_group_name":  "LogGroupName",
		"log_stream_name": "LogStreamName",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}