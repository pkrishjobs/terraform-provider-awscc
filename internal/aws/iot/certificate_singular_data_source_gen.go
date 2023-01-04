// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_certificate", certificateDataSource)
}

// certificateDataSource returns the Terraform awscc_iot_certificate data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::Certificate resource.
func certificateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CACertificatePem
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 65536,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"ca_certificate_pem": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DEFAULT",
		//	    "SNI_ONLY"
		//	  ],
		//	  "type": "string"
		//	}
		"certificate_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CertificatePem
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 65536,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"certificate_pem": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateSigningRequest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"certificate_signing_request": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ACTIVE",
		//	    "INACTIVE",
		//	    "REVOKED",
		//	    "PENDING_TRANSFER",
		//	    "PENDING_ACTIVATION"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoT::Certificate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Certificate").WithTerraformTypeName("awscc_iot_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"ca_certificate_pem":          "CACertificatePem",
		"certificate_mode":            "CertificateMode",
		"certificate_pem":             "CertificatePem",
		"certificate_signing_request": "CertificateSigningRequest",
		"id":                          "Id",
		"status":                      "Status",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
