// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Indicates whether the specified access point currently has a policy that allows
// public access. For more information about public access through access points,
// see Managing Data Access with Amazon S3 Access Points
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-points.html) in the
// Amazon Simple Storage Service Developer Guide.
func (c *Client) GetAccessPointPolicyStatus(ctx context.Context, params *GetAccessPointPolicyStatusInput, optFns ...func(*Options)) (*GetAccessPointPolicyStatusOutput, error) {
	if params == nil {
		params = &GetAccessPointPolicyStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessPointPolicyStatus", params, optFns, addOperationGetAccessPointPolicyStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessPointPolicyStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessPointPolicyStatusInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy status you want to retrieve.
	//
	// This member is required.
	Name *string
}

type GetAccessPointPolicyStatusOutput struct {

	// Indicates the current policy status of the specified access point.
	PolicyStatus *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAccessPointPolicyStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetAccessPointPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetAccessPointPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetAccessPointPolicyStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessPointPolicyStatus(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAccessPointPolicyStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetAccessPointPolicyStatus",
	}
}
