// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specify the default version of an extension. The default version of an
// extension will be used in CloudFormation operations.
func (c *Client) SetTypeDefaultVersion(ctx context.Context, params *SetTypeDefaultVersionInput, optFns ...func(*Options)) (*SetTypeDefaultVersionOutput, error) {
	if params == nil {
		params = &SetTypeDefaultVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTypeDefaultVersion", params, optFns, c.addOperationSetTypeDefaultVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTypeDefaultVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetTypeDefaultVersionInput struct {

	// The Amazon Resource Name (ARN) of the extension for which you want version
	// summary information.
	//
	// Conditional: You must specify either TypeName and Type , or Arn .
	Arn *string

	// The kind of extension.
	//
	// Conditional: You must specify either TypeName and Type , or Arn .
	Type types.RegistryType

	// The name of the extension.
	//
	// Conditional: You must specify either TypeName and Type , or Arn .
	TypeName *string

	// The ID of a specific version of the extension. The version ID is the value at
	// the end of the Amazon Resource Name (ARN) assigned to the extension version when
	// it is registered.
	VersionId *string

	noSmithyDocumentSerde
}

type SetTypeDefaultVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTypeDefaultVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetTypeDefaultVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTypeDefaultVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetTypeDefaultVersion"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTypeDefaultVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTypeDefaultVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetTypeDefaultVersion",
	}
}