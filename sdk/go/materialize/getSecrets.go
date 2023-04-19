// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package materialize

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-materialize/sdk/go/materialize"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := materialize.GetSecrets(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = materialize.GetSecrets(ctx, &GetSecretsArgs{
//				DatabaseName: pulumi.StringRef("materialize"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = materialize.GetSecrets(ctx, &GetSecretsArgs{
//				DatabaseName: pulumi.StringRef("materialize"),
//				SchemaName:   pulumi.StringRef("schema"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSecrets(ctx *pulumi.Context, args *GetSecretsArgs, opts ...pulumi.InvokeOption) (*GetSecretsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecretsResult
	err := ctx.Invoke("materialize:index/getSecrets:GetSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSecrets.
type GetSecretsArgs struct {
	// Limit secrets to a specific database
	DatabaseName *string `pulumi:"databaseName"`
	// Limit secrets to a specific schema within a specific database
	SchemaName *string `pulumi:"schemaName"`
}

// A collection of values returned by GetSecrets.
type GetSecretsResult struct {
	// Limit secrets to a specific database
	DatabaseName *string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Limit secrets to a specific schema within a specific database
	SchemaName *string `pulumi:"schemaName"`
	// The secrets in the account
	Secrets []GetSecretsSecret `pulumi:"secrets"`
}

func GetSecretsOutput(ctx *pulumi.Context, args GetSecretsOutputArgs, opts ...pulumi.InvokeOption) GetSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretsResult, error) {
			args := v.(GetSecretsArgs)
			r, err := GetSecrets(ctx, &args, opts...)
			var s GetSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretsResultOutput)
}

// A collection of arguments for invoking GetSecrets.
type GetSecretsOutputArgs struct {
	// Limit secrets to a specific database
	DatabaseName pulumi.StringPtrInput `pulumi:"databaseName"`
	// Limit secrets to a specific schema within a specific database
	SchemaName pulumi.StringPtrInput `pulumi:"schemaName"`
}

func (GetSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsArgs)(nil)).Elem()
}

// A collection of values returned by GetSecrets.
type GetSecretsResultOutput struct{ *pulumi.OutputState }

func (GetSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsResult)(nil)).Elem()
}

func (o GetSecretsResultOutput) ToGetSecretsResultOutput() GetSecretsResultOutput {
	return o
}

func (o GetSecretsResultOutput) ToGetSecretsResultOutputWithContext(ctx context.Context) GetSecretsResultOutput {
	return o
}

// Limit secrets to a specific database
func (o GetSecretsResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretsResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Limit secrets to a specific schema within a specific database
func (o GetSecretsResultOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretsResult) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

// The secrets in the account
func (o GetSecretsResultOutput) Secrets() GetSecretsSecretArrayOutput {
	return o.ApplyT(func(v GetSecretsResult) []GetSecretsSecret { return v.Secrets }).(GetSecretsSecretArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretsResultOutput{})
}
