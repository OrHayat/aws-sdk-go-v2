// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all available details about the instance fleets in a cluster. The instance
// fleet configuration is available only in Amazon EMR versions 4.8.0 and later,
// excluding 5.0.x versions.
func (c *Client) ListInstanceFleets(ctx context.Context, params *ListInstanceFleetsInput, optFns ...func(*Options)) (*ListInstanceFleetsOutput, error) {
	if params == nil {
		params = &ListInstanceFleetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstanceFleets", params, optFns, c.addOperationListInstanceFleetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstanceFleetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstanceFleetsInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

type ListInstanceFleetsOutput struct {

	// The list of instance fleets for the cluster and given filters.
	InstanceFleets []types.InstanceFleet

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInstanceFleetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInstanceFleets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInstanceFleets{}, middleware.After)
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
	if err = addOpListInstanceFleetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstanceFleets(options.Region), middleware.Before); err != nil {
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

// ListInstanceFleetsAPIClient is a client that implements the ListInstanceFleets
// operation.
type ListInstanceFleetsAPIClient interface {
	ListInstanceFleets(context.Context, *ListInstanceFleetsInput, ...func(*Options)) (*ListInstanceFleetsOutput, error)
}

var _ ListInstanceFleetsAPIClient = (*Client)(nil)

// ListInstanceFleetsPaginatorOptions is the paginator options for
// ListInstanceFleets
type ListInstanceFleetsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstanceFleetsPaginator is a paginator for ListInstanceFleets
type ListInstanceFleetsPaginator struct {
	options   ListInstanceFleetsPaginatorOptions
	client    ListInstanceFleetsAPIClient
	params    *ListInstanceFleetsInput
	nextToken *string
	firstPage bool
}

// NewListInstanceFleetsPaginator returns a new ListInstanceFleetsPaginator
func NewListInstanceFleetsPaginator(client ListInstanceFleetsAPIClient, params *ListInstanceFleetsInput, optFns ...func(*ListInstanceFleetsPaginatorOptions)) *ListInstanceFleetsPaginator {
	if params == nil {
		params = &ListInstanceFleetsInput{}
	}

	options := ListInstanceFleetsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstanceFleetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstanceFleetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInstanceFleets page.
func (p *ListInstanceFleetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstanceFleetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListInstanceFleets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInstanceFleets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListInstanceFleets",
	}
}