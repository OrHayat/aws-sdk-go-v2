// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreaminfluxdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Timestream for InfluxDB DB parameter groups.
func (c *Client) ListDbParameterGroups(ctx context.Context, params *ListDbParameterGroupsInput, optFns ...func(*Options)) (*ListDbParameterGroupsOutput, error) {
	if params == nil {
		params = &ListDbParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDbParameterGroups", params, optFns, c.addOperationListDbParameterGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDbParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDbParameterGroupsInput struct {

	// The maximum number of items to return in the output. If the total number of
	// items available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	MaxResults *int32

	// The pagination token. To resume pagination, provide the NextToken value as
	// argument of a subsequent API invocation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDbParameterGroupsOutput struct {

	// A list of Timestream for InfluxDB DB parameter group summaries.
	//
	// This member is required.
	Items []types.DbParameterGroupSummary

	// Token from a previous call of the operation. When this value is provided, the
	// service returns results from where the previous response left off.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDbParameterGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDbParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDbParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDbParameterGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDbParameterGroups(options.Region), middleware.Before); err != nil {
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

// ListDbParameterGroupsAPIClient is a client that implements the
// ListDbParameterGroups operation.
type ListDbParameterGroupsAPIClient interface {
	ListDbParameterGroups(context.Context, *ListDbParameterGroupsInput, ...func(*Options)) (*ListDbParameterGroupsOutput, error)
}

var _ ListDbParameterGroupsAPIClient = (*Client)(nil)

// ListDbParameterGroupsPaginatorOptions is the paginator options for
// ListDbParameterGroups
type ListDbParameterGroupsPaginatorOptions struct {
	// The maximum number of items to return in the output. If the total number of
	// items available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDbParameterGroupsPaginator is a paginator for ListDbParameterGroups
type ListDbParameterGroupsPaginator struct {
	options   ListDbParameterGroupsPaginatorOptions
	client    ListDbParameterGroupsAPIClient
	params    *ListDbParameterGroupsInput
	nextToken *string
	firstPage bool
}

// NewListDbParameterGroupsPaginator returns a new ListDbParameterGroupsPaginator
func NewListDbParameterGroupsPaginator(client ListDbParameterGroupsAPIClient, params *ListDbParameterGroupsInput, optFns ...func(*ListDbParameterGroupsPaginatorOptions)) *ListDbParameterGroupsPaginator {
	if params == nil {
		params = &ListDbParameterGroupsInput{}
	}

	options := ListDbParameterGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDbParameterGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDbParameterGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDbParameterGroups page.
func (p *ListDbParameterGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDbParameterGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListDbParameterGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDbParameterGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDbParameterGroups",
	}
}
