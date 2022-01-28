// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a geofence collection, which manages and stores geofences.
func (c *Client) CreateGeofenceCollection(ctx context.Context, params *CreateGeofenceCollectionInput, optFns ...func(*Options)) (*CreateGeofenceCollectionOutput, error) {
	if params == nil {
		params = &CreateGeofenceCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGeofenceCollection", params, optFns, c.addOperationCreateGeofenceCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGeofenceCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGeofenceCollectionInput struct {

	// A custom name for the geofence collection. Requirements:
	//
	// * Contain only
	// alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods (.), and
	// underscores (_).
	//
	// * Must be a unique geofence collection name.
	//
	// * No spaces
	// allowed. For example, ExampleGeofenceCollection.
	//
	// This member is required.
	CollectionName *string

	// An optional description for the geofence collection.
	Description *string

	// A key identifier for an AWS KMS customer managed key
	// (https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html). Enter
	// a key ID, key ARN, alias name, or alias ARN.
	KmsKeyId *string

	// No longer used. If included, the only allowed value is RequestBasedUsage.
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	// This parameter is no longer used.
	//
	// Deprecated: Deprecated. No longer allowed.
	PricingPlanDataSource *string

	// Applies one or more tags to the geofence collection. A tag is a key-value pair
	// helps manage, identify, search, and filter your resources by labelling them.
	// Format: "key" : "value" Restrictions:
	//
	// * Maximum 50 tags per resource
	//
	// * Each
	// resource tag must be unique with a maximum of one value.
	//
	// * Maximum key length:
	// 128 Unicode characters in UTF-8
	//
	// * Maximum value length: 256 Unicode characters
	// in UTF-8
	//
	// * Can use alphanumeric characters (A–Z, a–z, 0–9), and the following
	// characters: + - = . _ : / @.
	//
	// * Cannot use "aws:" as a prefix for a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateGeofenceCollectionOutput struct {

	// The Amazon Resource Name (ARN) for the geofence collection resource. Used when
	// you need to specify a resource across all AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:geofence-collection/ExampleGeofenceCollection
	//
	// This member is required.
	CollectionArn *string

	// The name for the geofence collection.
	//
	// This member is required.
	CollectionName *string

	// The timestamp for when the geofence collection was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ
	//
	// This member is required.
	CreateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGeofenceCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGeofenceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGeofenceCollection{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateGeofenceCollectionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateGeofenceCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGeofenceCollection(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateGeofenceCollectionMiddleware struct {
}

func (*endpointPrefix_opCreateGeofenceCollectionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateGeofenceCollectionMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "geofencing." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateGeofenceCollectionMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateGeofenceCollectionMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGeofenceCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CreateGeofenceCollection",
	}
}
