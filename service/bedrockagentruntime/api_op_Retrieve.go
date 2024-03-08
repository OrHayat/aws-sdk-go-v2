// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Queries a knowledge base and retrieves information from it.
func (c *Client) Retrieve(ctx context.Context, params *RetrieveInput, optFns ...func(*Options)) (*RetrieveOutput, error) {
	if params == nil {
		params = &RetrieveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Retrieve", params, optFns, c.addOperationRetrieveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetrieveInput struct {

	// The unique identifier of the knowledge base to query.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The query to send the knowledge base.
	//
	// This member is required.
	RetrievalQuery *types.KnowledgeBaseQuery

	// If there are more results than can fit in the response, the response returns a
	// nextToken . Use this token in the nextToken field of another request to
	// retrieve the next batch of results.
	NextToken *string

	// Contains details about how the results should be returned.
	RetrievalConfiguration *types.KnowledgeBaseRetrievalConfiguration

	noSmithyDocumentSerde
}

type RetrieveOutput struct {

	// A list of results from querying the knowledge base.
	//
	// This member is required.
	RetrievalResults []types.KnowledgeBaseRetrievalResult

	// If there are more results than can fit in the response, the response returns a
	// nextToken . Use this token in the nextToken field of another request to
	// retrieve the next batch of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetrieveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRetrieve{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRetrieve{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Retrieve"); err != nil {
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
	if err = addOpRetrieveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieve(options.Region), middleware.Before); err != nil {
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

// RetrieveAPIClient is a client that implements the Retrieve operation.
type RetrieveAPIClient interface {
	Retrieve(context.Context, *RetrieveInput, ...func(*Options)) (*RetrieveOutput, error)
}

var _ RetrieveAPIClient = (*Client)(nil)

// RetrievePaginatorOptions is the paginator options for Retrieve
type RetrievePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// RetrievePaginator is a paginator for Retrieve
type RetrievePaginator struct {
	options   RetrievePaginatorOptions
	client    RetrieveAPIClient
	params    *RetrieveInput
	nextToken *string
	firstPage bool
}

// NewRetrievePaginator returns a new RetrievePaginator
func NewRetrievePaginator(client RetrieveAPIClient, params *RetrieveInput, optFns ...func(*RetrievePaginatorOptions)) *RetrievePaginator {
	if params == nil {
		params = &RetrieveInput{}
	}

	options := RetrievePaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &RetrievePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *RetrievePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next Retrieve page.
func (p *RetrievePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*RetrieveOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.Retrieve(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opRetrieve(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Retrieve",
	}
}
