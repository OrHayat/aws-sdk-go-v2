// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// We strongly recommend that all Auto Scaling groups use launch templates to
// ensure full functionality for Amazon EC2 Auto Scaling and Amazon EC2. Updates
// the configuration for the specified Auto Scaling group. To update an Auto
// Scaling group, specify the name of the group and the property that you want to
// change. Any properties that you don't specify are not changed by this update
// request. The new settings take effect on any scaling activities after this call
// returns. If you associate a new launch configuration or template with an Auto
// Scaling group, all new instances will get the updated configuration. Existing
// instances continue to run with the configuration that they were originally
// launched with. When you update a group to specify a mixed instances policy
// instead of a launch configuration or template, existing instances may be
// replaced to match the new purchasing options that you specified in the policy.
// For example, if the group currently has 100% On-Demand capacity and the policy
// specifies 50% Spot capacity, this means that half of your instances will be
// gradually terminated and relaunched as Spot Instances. When replacing instances,
// Amazon EC2 Auto Scaling launches new instances before terminating the old ones,
// so that updating your group does not compromise the performance or availability
// of your application. Note the following about changing DesiredCapacity , MaxSize
// , or MinSize :
//   - If a scale-in activity occurs as a result of a new DesiredCapacity value
//     that is lower than the current size of the group, the Auto Scaling group uses
//     its termination policy to determine which instances to terminate.
//   - If you specify a new value for MinSize without specifying a value for
//     DesiredCapacity , and the new MinSize is larger than the current size of the
//     group, this sets the group's DesiredCapacity to the new MinSize value.
//   - If you specify a new value for MaxSize without specifying a value for
//     DesiredCapacity , and the new MaxSize is smaller than the current size of the
//     group, this sets the group's DesiredCapacity to the new MaxSize value.
//
// To see which properties have been set, call the DescribeAutoScalingGroups API.
// To view the scaling policies for an Auto Scaling group, call the
// DescribePolicies API. If the group has scaling policies, you can update them by
// calling the PutScalingPolicy API.
func (c *Client) UpdateAutoScalingGroup(ctx context.Context, params *UpdateAutoScalingGroupInput, optFns ...func(*Options)) (*UpdateAutoScalingGroupOutput, error) {
	if params == nil {
		params = &UpdateAutoScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAutoScalingGroup", params, optFns, c.addOperationUpdateAutoScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAutoScalingGroupInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more Availability Zones for the group.
	AvailabilityZones []string

	// Enables or disables Capacity Rebalancing. For more information, see Use
	// Capacity Rebalancing to handle Amazon EC2 Spot Interruptions (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	CapacityRebalance *bool

	// Reserved.
	Context *string

	// Only needed if you use simple scaling policies. The amount of time, in seconds,
	// between one scaling activity ending and another one starting due to simple
	// scaling policies. For more information, see Scaling cooldowns for Amazon EC2
	// Auto Scaling (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	DefaultCooldown *int32

	// The amount of time, in seconds, until a new instance is considered to have
	// finished initializing and resource consumption to become stable after it enters
	// the InService state. During an instance refresh, Amazon EC2 Auto Scaling waits
	// for the warm-up period after it replaces an instance before it moves on to
	// replacing the next instance. Amazon EC2 Auto Scaling also waits for the warm-up
	// period before aggregating the metrics for new instances with existing instances
	// in the Amazon CloudWatch metrics that are used for scaling, resulting in more
	// reliable usage data. For more information, see Set the default instance warmup
	// for an Auto Scaling group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html)
	// in the Amazon EC2 Auto Scaling User Guide. To manage various warm-up settings at
	// the group level, we recommend that you set the default instance warmup, even if
	// it is set to 0 seconds. To remove a value that you previously set, include the
	// property but specify -1 for the value. However, we strongly recommend keeping
	// the default instance warmup enabled by specifying a value of 0 or other nominal
	// value.
	DefaultInstanceWarmup *int32

	// The desired capacity is the initial capacity of the Auto Scaling group after
	// this operation completes and the capacity it attempts to maintain. This number
	// must be greater than or equal to the minimum size of the group and less than or
	// equal to the maximum size of the group.
	DesiredCapacity *int32

	// The unit of measurement for the value specified for desired capacity. Amazon
	// EC2 Auto Scaling supports DesiredCapacityType for attribute-based instance type
	// selection only. For more information, see Creating an Auto Scaling group using
	// attribute-based instance type selection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-instance-type-requirements.html)
	// in the Amazon EC2 Auto Scaling User Guide. By default, Amazon EC2 Auto Scaling
	// specifies units , which translates into number of instances. Valid values: units
	// | vcpu | memory-mib
	DesiredCapacityType *string

	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before
	// checking the health status of an EC2 instance that has come into service and
	// marking it unhealthy due to a failed health check. This is useful if your
	// instances do not immediately pass their health checks after they enter the
	// InService state. For more information, see Set the health check grace period
	// for an Auto Scaling group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	HealthCheckGracePeriod *int32

	// A comma-separated value string of one or more health check types. The valid
	// values are EC2 , ELB , and VPC_LATTICE . EC2 is the default health check and
	// cannot be disabled. For more information, see Health checks for Auto Scaling
	// instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html)
	// in the Amazon EC2 Auto Scaling User Guide. Only specify EC2 if you must clear a
	// value that was previously set.
	HealthCheckType *string

	// An instance maintenance policy. For more information, see Set instance
	// maintenance policy (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	InstanceMaintenancePolicy *types.InstanceMaintenancePolicy

	// The name of the launch configuration. If you specify LaunchConfigurationName in
	// your update request, you can't specify LaunchTemplate or MixedInstancesPolicy .
	LaunchConfigurationName *string

	// The launch template and version to use to specify the updates. If you specify
	// LaunchTemplate in your update request, you can't specify LaunchConfigurationName
	// or MixedInstancesPolicy .
	LaunchTemplate *types.LaunchTemplateSpecification

	// The maximum amount of time, in seconds, that an instance can be in service. The
	// default is null. If specified, the value must be either 0 or a number equal to
	// or greater than 86,400 seconds (1 day). To clear a previously set value, specify
	// a new value of 0. For more information, see Replacing Auto Scaling instances
	// based on maximum instance lifetime (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MaxInstanceLifetime *int32

	// The maximum size of the Auto Scaling group. With a mixed instances policy that
	// uses instance weighting, Amazon EC2 Auto Scaling may need to go above MaxSize
	// to meet your capacity requirements. In this event, Amazon EC2 Auto Scaling will
	// never go above MaxSize by more than your largest instance weight (weights that
	// define how many units each instance contributes to the desired capacity of the
	// group).
	MaxSize *int32

	// The minimum size of the Auto Scaling group.
	MinSize *int32

	// The mixed instances policy. For more information, see Auto Scaling groups with
	// multiple instance types and purchase options (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	MixedInstancesPolicy *types.MixedInstancesPolicy

	// Indicates whether newly launched instances are protected from termination by
	// Amazon EC2 Auto Scaling when scaling in. For more information about preventing
	// instances from terminating on scale in, see Using instance scale-in protection (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	NewInstancesProtectedFromScaleIn *bool

	// The name of an existing placement group into which to launch your instances.
	// For more information, see Placement groups (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html)
	// in the Amazon EC2 User Guide for Linux Instances. A cluster placement group is a
	// logical grouping of instances within a single Availability Zone. You cannot
	// specify multiple Availability Zones and a cluster placement group.
	PlacementGroup *string

	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling
	// group uses to call other Amazon Web Services on your behalf. For more
	// information, see Service-linked roles (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	ServiceLinkedRoleARN *string

	// A policy or a list of policies that are used to select the instances to
	// terminate. The policies are executed in the order that you list them. For more
	// information, see Work with Amazon EC2 Auto Scaling termination policies (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.html)
	// in the Amazon EC2 Auto Scaling User Guide. Valid values: Default |
	// AllocationStrategy | ClosestToNextInstanceHour | NewestInstance | OldestInstance
	// | OldestLaunchConfiguration | OldestLaunchTemplate |
	// arn:aws:lambda:region:account-id:function:my-function:my-alias
	TerminationPolicies []string

	// A comma-separated list of subnet IDs for a virtual private cloud (VPC). If you
	// specify VPCZoneIdentifier with AvailabilityZones , the subnets that you specify
	// must reside in those Availability Zones.
	VPCZoneIdentifier *string

	noSmithyDocumentSerde
}

type UpdateAutoScalingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAutoScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateAutoScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateAutoScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAutoScalingGroup"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateAutoScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAutoScalingGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateAutoScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAutoScalingGroup",
	}
}
