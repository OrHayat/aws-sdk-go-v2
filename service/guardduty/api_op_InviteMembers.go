// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invites other Amazon Web Services accounts (created as members of the current
// Amazon Web Services account by CreateMembers) to enable GuardDuty, and allow the
// current Amazon Web Services account to view and manage these accounts' findings
// on their behalf as the GuardDuty administrator account.
func (c *Client) InviteMembers(ctx context.Context, params *InviteMembersInput, optFns ...func(*Options)) (*InviteMembersOutput, error) {
	if params == nil {
		params = &InviteMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InviteMembers", params, optFns, c.addOperationInviteMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InviteMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteMembersInput struct {

	// A list of account IDs of the accounts that you want to invite to GuardDuty as
	// members.
	//
	// This member is required.
	AccountIds []string

	// The unique ID of the detector of the GuardDuty account that you want to invite
	// members with.
	//
	// This member is required.
	DetectorId *string

	// A Boolean value that specifies whether you want to disable email notification to
	// the accounts that you are inviting to GuardDuty as members.
	DisableEmailNotification bool

	// The invitation message that you want to send to the accounts that you're
	// inviting to GuardDuty as members.
	Message *string

	noSmithyDocumentSerde
}

type InviteMembersOutput struct {

	// A list of objects that contain the unprocessed account and a result string that
	// explains why it was unprocessed.
	//
	// This member is required.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInviteMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInviteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInviteMembers{}, middleware.After)
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
	if err = addOpInviteMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInviteMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInviteMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "InviteMembers",
	}
}
