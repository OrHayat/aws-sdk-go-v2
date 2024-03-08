// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides the feedback for an authentication event, whether it was from a valid
// user or not. This feedback is used for improving the risk evaluation decision
// for the user pool as part of Amazon Cognito advanced security. Amazon Cognito
// doesn't evaluate Identity and Access Management (IAM) policies in requests for
// this API operation. For this operation, you can't use IAM credentials to
// authorize requests, and you can't grant IAM permissions in policies. For more
// information about authorization models in Amazon Cognito, see Using the Amazon
// Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) UpdateAuthEventFeedback(ctx context.Context, params *UpdateAuthEventFeedbackInput, optFns ...func(*Options)) (*UpdateAuthEventFeedbackOutput, error) {
	if params == nil {
		params = &UpdateAuthEventFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAuthEventFeedback", params, optFns, c.addOperationUpdateAuthEventFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAuthEventFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAuthEventFeedbackInput struct {

	// The event ID.
	//
	// This member is required.
	EventId *string

	// The feedback token.
	//
	// This member is required.
	FeedbackToken *string

	// The authentication event feedback value. When you provide a FeedbackValue value
	// of valid , you tell Amazon Cognito that you trust a user session where Amazon
	// Cognito has evaluated some level of risk. When you provide a FeedbackValue
	// value of invalid , you tell Amazon Cognito that you don't trust a user session,
	// or you don't believe that Amazon Cognito evaluated a high-enough risk level.
	//
	// This member is required.
	FeedbackValue types.FeedbackValueType

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The username of the user that you want to query or modify. The value of this
	// parameter is typically your user's username, but it can be any of their alias
	// attributes. If username isn't an alias attribute in your user pool, this value
	// must be the sub of a local user or the username of a user from a third-party
	// IdP.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

type UpdateAuthEventFeedbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAuthEventFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAuthEventFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAuthEventFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAuthEventFeedback"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addOpUpdateAuthEventFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAuthEventFeedback(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAuthEventFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAuthEventFeedback",
	}
}
