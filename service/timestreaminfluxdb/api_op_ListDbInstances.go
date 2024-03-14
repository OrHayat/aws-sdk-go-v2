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

// Returns a list of Timestream for InfluxDB DB instances.
func (c *Client) ListDbInstances(ctx context.Context, params *ListDbInstancesInput, optFns ...func(*Options)) (*ListDbInstancesOutput, error) {
	if params == nil {
		params = &ListDbInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDbInstances", params, optFns, c.addOperationListDbInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDbInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDbInstancesInput struct {

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

type ListDbInstancesOutput struct {

	// A list of Timestream for InfluxDB DB instance summaries.
	//
	// This member is required.
	Items []types.DbInstanceSummary

	// Token from a previous call of the operation. When this value is provided, the
	// service returns results from where the previous response left off.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDbInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDbInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDbInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDbInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDbInstances(options.Region), middleware.Before); err != nil {
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

// ListDbInstancesAPIClient is a client that implements the ListDbInstances
// operation.
type ListDbInstancesAPIClient interface {
	ListDbInstances(context.Context, *ListDbInstancesInput, ...func(*Options)) (*ListDbInstancesOutput, error)
}

var _ ListDbInstancesAPIClient = (*Client)(nil)

// ListDbInstancesPaginatorOptions is the paginator options for ListDbInstances
type ListDbInstancesPaginatorOptions struct {
	// The maximum number of items to return in the output. If the total number of
	// items available is more than the value specified, a NextToken is provided in the
	// output. To resume pagination, provide the NextToken value as argument of a
	// subsequent API invocation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDbInstancesPaginator is a paginator for ListDbInstances
type ListDbInstancesPaginator struct {
	options   ListDbInstancesPaginatorOptions
	client    ListDbInstancesAPIClient
	params    *ListDbInstancesInput
	nextToken *string
	firstPage bool
}

// NewListDbInstancesPaginator returns a new ListDbInstancesPaginator
func NewListDbInstancesPaginator(client ListDbInstancesAPIClient, params *ListDbInstancesInput, optFns ...func(*ListDbInstancesPaginatorOptions)) *ListDbInstancesPaginator {
	if params == nil {
		params = &ListDbInstancesInput{}
	}

	options := ListDbInstancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDbInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDbInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDbInstances page.
func (p *ListDbInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDbInstancesOutput, error) {
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

	result, err := p.client.ListDbInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDbInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDbInstances",
	}
}
