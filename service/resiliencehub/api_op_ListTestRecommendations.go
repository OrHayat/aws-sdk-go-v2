// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the test recommendations for the Resilience Hub application.
func (c *Client) ListTestRecommendations(ctx context.Context, params *ListTestRecommendationsInput, optFns ...func(*Options)) (*ListTestRecommendationsOutput, error) {
	if params == nil {
		params = &ListTestRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTestRecommendations", params, optFns, c.addOperationListTestRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTestRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTestRecommendationsInput struct {

	// The Amazon Resource Name (ARN) of the assessment. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app-assessment/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	AssessmentArn *string

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTestRecommendationsOutput struct {

	// The test recommendations for the Resilience Hub application.
	//
	// This member is required.
	TestRecommendations []types.TestRecommendation

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTestRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTestRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTestRecommendations{}, middleware.After)
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
	if err = addOpListTestRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTestRecommendations(options.Region), middleware.Before); err != nil {
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

// ListTestRecommendationsAPIClient is a client that implements the
// ListTestRecommendations operation.
type ListTestRecommendationsAPIClient interface {
	ListTestRecommendations(context.Context, *ListTestRecommendationsInput, ...func(*Options)) (*ListTestRecommendationsOutput, error)
}

var _ ListTestRecommendationsAPIClient = (*Client)(nil)

// ListTestRecommendationsPaginatorOptions is the paginator options for
// ListTestRecommendations
type ListTestRecommendationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTestRecommendationsPaginator is a paginator for ListTestRecommendations
type ListTestRecommendationsPaginator struct {
	options   ListTestRecommendationsPaginatorOptions
	client    ListTestRecommendationsAPIClient
	params    *ListTestRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewListTestRecommendationsPaginator returns a new
// ListTestRecommendationsPaginator
func NewListTestRecommendationsPaginator(client ListTestRecommendationsAPIClient, params *ListTestRecommendationsInput, optFns ...func(*ListTestRecommendationsPaginatorOptions)) *ListTestRecommendationsPaginator {
	if params == nil {
		params = &ListTestRecommendationsInput{}
	}

	options := ListTestRecommendationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTestRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTestRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTestRecommendations page.
func (p *ListTestRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTestRecommendationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListTestRecommendations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTestRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "ListTestRecommendations",
	}
}