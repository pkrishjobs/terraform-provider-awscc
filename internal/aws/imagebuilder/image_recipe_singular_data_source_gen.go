// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_imagebuilder_image_recipe", imageRecipeDataSource)
}

// imageRecipeDataSource returns the Terraform awscc_imagebuilder_image_recipe data source.
// This Terraform data source corresponds to the CloudFormation AWS::ImageBuilder::ImageRecipe resource.
func imageRecipeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdditionalInstanceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specify additional settings and launch scripts for your build instances.",
		//	  "properties": {
		//	    "SystemsManagerAgent": {
		//	      "additionalProperties": false,
		//	      "description": "Contains settings for the SSM agent on your build instance.",
		//	      "properties": {
		//	        "UninstallAfterBuild": {
		//	          "description": "Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI. If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "UserDataOverride": {
		//	      "description": "Use this property to provide commands or a command script to run when you launch your build instance.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"additional_instance_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SystemsManagerAgent
				"systems_manager_agent": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: UninstallAfterBuild
						"uninstall_after_build": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI. If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Contains settings for the SSM agent on your build instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UserDataOverride
				"user_data_override": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Use this property to provide commands or a command script to run when you launch your build instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specify additional settings and launch scripts for your build instances.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the image recipe.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BlockDeviceMappings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The block device mappings to apply when creating images from this recipe.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Defines block device mappings for the instance used to configure your image. ",
		//	    "properties": {
		//	      "DeviceName": {
		//	        "description": "The device to which these mappings apply.",
		//	        "type": "string"
		//	      },
		//	      "Ebs": {
		//	        "additionalProperties": false,
		//	        "description": "Use to manage Amazon EBS-specific configuration for this mapping.",
		//	        "properties": {
		//	          "DeleteOnTermination": {
		//	            "description": "Use to configure delete on termination of the associated device.",
		//	            "type": "boolean"
		//	          },
		//	          "Encrypted": {
		//	            "description": "Use to configure device encryption.",
		//	            "type": "boolean"
		//	          },
		//	          "Iops": {
		//	            "description": "Use to configure device IOPS.",
		//	            "type": "integer"
		//	          },
		//	          "KmsKeyId": {
		//	            "description": "Use to configure the KMS key to use when encrypting the device.",
		//	            "type": "string"
		//	          },
		//	          "SnapshotId": {
		//	            "description": "The snapshot that defines the device contents.",
		//	            "type": "string"
		//	          },
		//	          "Throughput": {
		//	            "description": "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
		//	            "type": "integer"
		//	          },
		//	          "VolumeSize": {
		//	            "description": "Use to override the device's volume size.",
		//	            "type": "integer"
		//	          },
		//	          "VolumeType": {
		//	            "description": "Use to override the device's volume type.",
		//	            "enum": [
		//	              "standard",
		//	              "io1",
		//	              "io2",
		//	              "gp2",
		//	              "gp3",
		//	              "sc1",
		//	              "st1"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "NoDevice": {
		//	        "description": "Use to remove a mapping from the parent image.",
		//	        "type": "string"
		//	      },
		//	      "VirtualName": {
		//	        "description": "Use to manage instance ephemeral devices.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"block_device_mappings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DeviceName
					"device_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The device to which these mappings apply.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Ebs
					"ebs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DeleteOnTermination
							"delete_on_termination": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Use to configure delete on termination of the associated device.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Encrypted
							"encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
								Description: "Use to configure device encryption.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Iops
							"iops": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Use to configure device IOPS.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: KmsKeyId
							"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Use to configure the KMS key to use when encrypting the device.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SnapshotId
							"snapshot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The snapshot that defines the device contents.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Throughput
							"throughput": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: VolumeSize
							"volume_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "Use to override the device's volume size.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: VolumeType
							"volume_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Use to override the device's volume type.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Use to manage Amazon EBS-specific configuration for this mapping.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: NoDevice
					"no_device": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Use to remove a mapping from the parent image.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: VirtualName
					"virtual_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Use to manage instance ephemeral devices.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The block device mappings to apply when creating images from this recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Components
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The components of the image recipe.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Configuration details of the component.",
		//	    "properties": {
		//	      "ComponentArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the component.",
		//	        "type": "string"
		//	      },
		//	      "Parameters": {
		//	        "description": "A group of parameter settings that are used to configure the component for a specific recipe.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "Contains a key/value pair that sets the named component parameter.",
		//	          "properties": {
		//	            "Name": {
		//	              "description": "The name of the component parameter to set.",
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "description": "Sets the value for the named component parameter.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "Name",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"components": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ComponentArn
					"component_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the component.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Name
								"name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the component parameter to set.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "Sets the value for the named component parameter.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A group of parameter settings that are used to configure the component for a specific recipe.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The components of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the image recipe.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the image recipe.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ParentImage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parent image of the image recipe.",
		//	  "type": "string"
		//	}
		"parent_image": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The parent image of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The tags of the image recipe.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The tags of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the image recipe.",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the image recipe.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkingDirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The working directory to be used during build and test workflows.",
		//	  "type": "string"
		//	}
		"working_directory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The working directory to be used during build and test workflows.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ImageBuilder::ImageRecipe",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ImageRecipe").WithTerraformTypeName("awscc_imagebuilder_image_recipe")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_instance_configuration": "AdditionalInstanceConfiguration",
		"arn":                               "Arn",
		"block_device_mappings":             "BlockDeviceMappings",
		"component_arn":                     "ComponentArn",
		"components":                        "Components",
		"delete_on_termination":             "DeleteOnTermination",
		"description":                       "Description",
		"device_name":                       "DeviceName",
		"ebs":                               "Ebs",
		"encrypted":                         "Encrypted",
		"iops":                              "Iops",
		"kms_key_id":                        "KmsKeyId",
		"name":                              "Name",
		"no_device":                         "NoDevice",
		"parameters":                        "Parameters",
		"parent_image":                      "ParentImage",
		"snapshot_id":                       "SnapshotId",
		"systems_manager_agent":             "SystemsManagerAgent",
		"tags":                              "Tags",
		"throughput":                        "Throughput",
		"uninstall_after_build":             "UninstallAfterBuild",
		"user_data_override":                "UserDataOverride",
		"value":                             "Value",
		"version":                           "Version",
		"virtual_name":                      "VirtualName",
		"volume_size":                       "VolumeSize",
		"volume_type":                       "VolumeType",
		"working_directory":                 "WorkingDirectory",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
