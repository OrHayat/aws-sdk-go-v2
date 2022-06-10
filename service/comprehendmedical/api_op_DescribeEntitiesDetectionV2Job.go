// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with a medical entities detection job. Use this
// operation to get the status of a detection job.
func (c *Client) DescribeEntitiesDetectionV2Job(ctx context.Context, params *DescribeEntitiesDetectionV2JobInput, optFns ...func(*Options)) (*DescribeEntitiesDetectionV2JobOutput, error) {
	if params == nil {
		params = &DescribeEntitiesDetectionV2JobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEntitiesDetectionV2Job", params, optFns, c.addOperationDescribeEntitiesDetectionV2JobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEntitiesDetectionV2JobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEntitiesDetectionV2JobInput struct {

	// The identifier that Comprehend Medical; generated for the job. The
	// StartEntitiesDetectionV2Job operation returns this identifier in its response.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type DescribeEntitiesDetectionV2JobOutput struct {

	// An object that contains the properties associated with a detection job.
	ComprehendMedicalAsyncJobProperties *types.ComprehendMedicalAsyncJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEntitiesDetectionV2JobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEntitiesDetectionV2Job{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEntitiesDetectionV2Job{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeEntitiesDetectionV2JobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEntitiesDetectionV2Job(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDescribeEntitiesDetectionV2Job(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "DescribeEntitiesDetectionV2Job",
	}
}