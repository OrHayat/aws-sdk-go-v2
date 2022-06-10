// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//
func (c *Client) PutDedicatedIpWarmupAttributes(ctx context.Context, params *PutDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutDedicatedIpWarmupAttributesOutput, error) {
	if params == nil {
		params = &PutDedicatedIpWarmupAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutDedicatedIpWarmupAttributes", params, optFns, c.addOperationPutDedicatedIpWarmupAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutDedicatedIpWarmupAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change the warm-up attributes for a dedicated IP address. This
// operation is useful when you want to resume the warm-up process for an existing
// IP address.
type PutDedicatedIpWarmupAttributesInput struct {

	// The dedicated IP address that you want to update the warm-up attributes for.
	//
	// This member is required.
	Ip *string

	// The warm-up percentage that you want to associate with the dedicated IP address.
	//
	// This member is required.
	WarmupPercentage *int32

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutDedicatedIpWarmupAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutDedicatedIpWarmupAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutDedicatedIpWarmupAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutDedicatedIpWarmupAttributes{}, middleware.After)
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
	if err = addOpPutDedicatedIpWarmupAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutDedicatedIpWarmupAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutDedicatedIpWarmupAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutDedicatedIpWarmupAttributes",
	}
}