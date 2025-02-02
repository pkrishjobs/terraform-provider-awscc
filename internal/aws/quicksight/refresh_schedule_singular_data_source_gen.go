// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_quicksight_refresh_schedule", refreshScheduleDataSource)
}

// refreshScheduleDataSource returns the Terraform awscc_quicksight_refresh_schedule data source.
// This Terraform data source corresponds to the CloudFormation AWS::QuickSight::RefreshSchedule resource.
func refreshScheduleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the data source.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Resource Name (ARN) of the data source.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DataSetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"data_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "RefreshType": {
		//	      "enum": [
		//	        "FULL_REFRESH",
		//	        "INCREMENTAL_REFRESH"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ScheduleFrequency": {
		//	      "additionalProperties": false,
		//	      "description": "\u003cp\u003eInformation about the schedule frequency.\u003c/p\u003e",
		//	      "properties": {
		//	        "Interval": {
		//	          "enum": [
		//	            "MINUTE15",
		//	            "MINUTE30",
		//	            "HOURLY",
		//	            "DAILY",
		//	            "WEEKLY",
		//	            "MONTHLY"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "RefreshOnDay": {
		//	          "additionalProperties": false,
		//	          "description": "\u003cp\u003eThe day scheduled for refresh.\u003c/p\u003e",
		//	          "properties": {
		//	            "DayOfMonth": {
		//	              "description": "\u003cp\u003eThe Day Of Month for scheduled refresh.\u003c/p\u003e",
		//	              "maxLength": 128,
		//	              "minLength": 1,
		//	              "type": "string"
		//	            },
		//	            "DayOfWeek": {
		//	              "enum": [
		//	                "SUNDAY",
		//	                "MONDAY",
		//	                "TUESDAY",
		//	                "WEDNESDAY",
		//	                "THURSDAY",
		//	                "FRIDAY",
		//	                "SATURDAY"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "TimeOfTheDay": {
		//	          "description": "\u003cp\u003eThe time of the day for scheduled refresh.\u003c/p\u003e",
		//	          "maxLength": 128,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "TimeZone": {
		//	          "description": "\u003cp\u003eThe timezone for scheduled refresh.\u003c/p\u003e",
		//	          "maxLength": 128,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ScheduleId": {
		//	      "description": "\u003cp\u003eAn unique identifier for the refresh schedule.\u003c/p\u003e",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "StartAfterDateTime": {
		//	      "description": "\u003cp\u003eThe date time after which refresh is to be scheduled\u003c/p\u003e",
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schedule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RefreshType
				"refresh_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ScheduleFrequency
				"schedule_frequency": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Interval
						"interval": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: RefreshOnDay
						"refresh_on_day": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DayOfMonth
								"day_of_month": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "<p>The Day Of Month for scheduled refresh.</p>",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: DayOfWeek
								"day_of_week": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "<p>The day scheduled for refresh.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: TimeOfTheDay
						"time_of_the_day": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "<p>The time of the day for scheduled refresh.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: TimeZone
						"time_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "<p>The timezone for scheduled refresh.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "<p>Information about the schedule frequency.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ScheduleId
				"schedule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>An unique identifier for the refresh schedule.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StartAfterDateTime
				"start_after_date_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The date time after which refresh is to be scheduled</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::QuickSight::RefreshSchedule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::QuickSight::RefreshSchedule").WithTerraformTypeName("awscc_quicksight_refresh_schedule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"aws_account_id":        "AwsAccountId",
		"data_set_id":           "DataSetId",
		"day_of_month":          "DayOfMonth",
		"day_of_week":           "DayOfWeek",
		"interval":              "Interval",
		"refresh_on_day":        "RefreshOnDay",
		"refresh_type":          "RefreshType",
		"schedule":              "Schedule",
		"schedule_frequency":    "ScheduleFrequency",
		"schedule_id":           "ScheduleId",
		"start_after_date_time": "StartAfterDateTime",
		"time_of_the_day":       "TimeOfTheDay",
		"time_zone":             "TimeZone",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
