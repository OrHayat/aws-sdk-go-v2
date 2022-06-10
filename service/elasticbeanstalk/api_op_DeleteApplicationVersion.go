// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified version from the specified application. You cannot delete
// an application version that is associated with a running environment.
func (c *Client) DeleteApplicationVersion(ctx context.Context, params *DeleteApplicationVersionInput, optFns ...func(*Options)) (*DeleteApplicationVersionOutput, error) {
	if params == nil {
		params = &DeleteApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationVersion", params, optFns, c.addOperationDeleteApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to delete an application version.
type DeleteApplicationVersionInput struct {

	// The name of the application to which the version belongs.
	//
	// This member is required.
	ApplicationName *string

	// The label of the version to delete.
	//
	// This member is required.
	VersionLabel *string

	// Set to true to delete the source bundle from your storage bucket. Otherwise, the
	// application version is deleted only from Elastic Beanstalk and the source bundle
	// remains in Amazon S3.
	DeleteSourceBundle *bool

	noSmithyDocumentSerde
}

type DeleteApplicationVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteApplicationVersion{}, middleware.After)
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
	if err = addOpDeleteApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DeleteApplicationVersion",
	}
}