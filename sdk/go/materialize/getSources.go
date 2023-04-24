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
//			_, err := materialize.GetSources(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = materialize.GetSources(ctx, &GetSourcesArgs{
//				DatabaseName: pulumi.StringRef("materialize"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = materialize.GetSources(ctx, &GetSourcesArgs{
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
func GetSources(ctx *pulumi.Context, args *GetSourcesArgs, opts ...pulumi.InvokeOption) (*GetSourcesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSourcesResult
	err := ctx.Invoke("materialize:index/getSources:GetSources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetSources.
type GetSourcesArgs struct {
	// Limit sources to a specific database
	DatabaseName *string `pulumi:"databaseName"`
	// Limit sources to a specific schema within a specific database
	SchemaName *string `pulumi:"schemaName"`
}

// A collection of values returned by GetSources.
type GetSourcesResult struct {
	// Limit sources to a specific database
	DatabaseName *string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Limit sources to a specific schema within a specific database
	SchemaName *string `pulumi:"schemaName"`
	// The sources in the account
	Sources []GetSourcesSource `pulumi:"sources"`
}

func GetSourcesOutput(ctx *pulumi.Context, args GetSourcesOutputArgs, opts ...pulumi.InvokeOption) GetSourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSourcesResult, error) {
			args := v.(GetSourcesArgs)
			r, err := GetSources(ctx, &args, opts...)
			var s GetSourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSourcesResultOutput)
}

// A collection of arguments for invoking GetSources.
type GetSourcesOutputArgs struct {
	// Limit sources to a specific database
	DatabaseName pulumi.StringPtrInput `pulumi:"databaseName"`
	// Limit sources to a specific schema within a specific database
	SchemaName pulumi.StringPtrInput `pulumi:"schemaName"`
}

func (GetSourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSourcesArgs)(nil)).Elem()
}

// A collection of values returned by GetSources.
type GetSourcesResultOutput struct{ *pulumi.OutputState }

func (GetSourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSourcesResult)(nil)).Elem()
}

func (o GetSourcesResultOutput) ToGetSourcesResultOutput() GetSourcesResultOutput {
	return o
}

func (o GetSourcesResultOutput) ToGetSourcesResultOutputWithContext(ctx context.Context) GetSourcesResultOutput {
	return o
}

// Limit sources to a specific database
func (o GetSourcesResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSourcesResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSourcesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourcesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Limit sources to a specific schema within a specific database
func (o GetSourcesResultOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSourcesResult) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

// The sources in the account
func (o GetSourcesResultOutput) Sources() GetSourcesSourceArrayOutput {
	return o.ApplyT(func(v GetSourcesResult) []GetSourcesSource { return v.Sources }).(GetSourcesSourceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSourcesResultOutput{})
}
