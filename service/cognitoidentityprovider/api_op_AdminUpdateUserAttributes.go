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

// This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with Amazon Pinpoint (https://console.aws.amazon.com/pinpoint/home/)
// . Amazon Cognito uses the registered number automatically. Otherwise, Amazon
// Cognito users who must receive SMS messages might not be able to sign up,
// activate their accounts, or sign in. If you have never used SMS text messages
// with Amazon Cognito or any other Amazon Web Service, Amazon Simple Notification
// Service might place your account in the SMS sandbox. In sandbox mode (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// , you can send messages only to verified phone numbers. After you test your app
// while in the sandbox environment, you can move out of the sandbox and into
// production. For more information, see SMS message settings for Amazon Cognito
// user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html)
// in the Amazon Cognito Developer Guide. Updates the specified user's attributes,
// including developer attributes, as an administrator. Works on any user. To
// delete an attribute from your user, submit the attribute in your API request
// with a blank value. For custom attributes, you must prepend the custom: prefix
// to the attribute name. In addition to updating user attributes, this API can
// also be used to mark phone and email as verified. Amazon Cognito evaluates
// Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you must use IAM credentials to authorize
// requests, and you must grant yourself the corresponding IAM permission in a
// policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) AdminUpdateUserAttributes(ctx context.Context, params *AdminUpdateUserAttributesInput, optFns ...func(*Options)) (*AdminUpdateUserAttributesOutput, error) {
	if params == nil {
		params = &AdminUpdateUserAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminUpdateUserAttributes", params, optFns, c.addOperationAdminUpdateUserAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminUpdateUserAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update the user's attributes as an administrator.
type AdminUpdateUserAttributesInput struct {

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name. If your
	// user pool requires verification before Amazon Cognito updates an attribute value
	// that you specify in this request, Amazon Cognito doesn’t immediately update the
	// value of that attribute. After your user receives and responds to a verification
	// message to verify the new value, Amazon Cognito updates the attribute value.
	// Your user can sign in and receive messages with the original attribute value
	// until they verify the new value. To update the value of an attribute that
	// requires verification in the same API request, include the email_verified or
	// phone_number_verified attribute, with a value of true . If you set the
	// email_verified or phone_number_verified value for an email or phone_number
	// attribute that requires verification to true , Amazon Cognito doesn’t send a
	// verification message to your user.
	//
	// This member is required.
	UserAttributes []types.AttributeType

	// The user pool ID for the user pool where you want to update user attributes.
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

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the
	// AdminUpdateUserAttributes API action, Amazon Cognito invokes the function that
	// is assigned to the custom message trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your AdminUpdateUserAttributes
	// request. In your function code in Lambda, you can process the clientMetadata
	// value to enhance your workflow for your specific needs. For more information,
	// see Customizing user pool Workflows with Lambda Triggers (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//   - Validate the ClientMetadata value.
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	ClientMetadata map[string]string

	noSmithyDocumentSerde
}

// Represents the response from the server for the request to update user
// attributes as an administrator.
type AdminUpdateUserAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminUpdateUserAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminUpdateUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminUpdateUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AdminUpdateUserAttributes"); err != nil {
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
	if err = addComputePayloadSHA256(stack); err != nil {
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
	if err = addOpAdminUpdateUserAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminUpdateUserAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminUpdateUserAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AdminUpdateUserAttributes",
	}
}
