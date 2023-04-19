// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package materialize

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A sink describes an external system you want Materialize to write data to, and provides details about how to encode that data.
//
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
//			_, err := materialize.NewSinkKafka(ctx, "exampleSinkKafka", &materialize.SinkKafkaArgs{
//				Envelope: &SinkKafkaEnvelopeArgs{
//					Upsert: pulumi.Bool(true),
//				},
//				Format: &SinkKafkaFormatArgs{
//					Avro: &SinkKafkaFormatAvroArgs{
//						SchemaRegistryConnection: &SinkKafkaFormatAvroSchemaRegistryConnectionArgs{
//							DatabaseName: pulumi.String("database"),
//							Name:         pulumi.String("csr_connection"),
//							SchemaName:   pulumi.String("schema"),
//						},
//					},
//				},
//				From: &SinkKafkaFromArgs{
//					Name: pulumi.String("table"),
//				},
//				KafkaConnection: &SinkKafkaKafkaConnectionArgs{
//					Name: pulumi.String("kafka_connection"),
//				},
//				SchemaName: pulumi.String("schema"),
//				Size:       pulumi.String("3xsmall"),
//				Topic:      pulumi.String("test_avro_topic"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Sinks can be imported using the sink id
//
// ```sh
//
//	$ pulumi import materialize:index/sinkKafka:SinkKafka example_sink_kafka <sink_id>
//
// ```
type SinkKafka struct {
	pulumi.CustomResourceState

	// The cluster to maintain this sink. If not specified, the size option must be specified.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The identifier for the sink database.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// How to interpret records (e.g. Debezium, Upsert).
	Envelope SinkKafkaEnvelopePtrOutput `pulumi:"envelope"`
	// How to decode raw bytes from different formats into data structures it can understand at runtime.
	Format SinkKafkaFormatPtrOutput `pulumi:"format"`
	// The name of the source, table or materialized view you want to send to the sink.
	From SinkKafkaFromOutput `pulumi:"from"`
	// The name of the Kafka connection to use in the sink.
	KafkaConnection SinkKafkaKafkaConnectionOutput `pulumi:"kafkaConnection"`
	// An optional list of columns to use for the Kafka key. If unspecified, the Kafka key is left unset.
	Keys pulumi.StringArrayOutput `pulumi:"keys"`
	// The identifier for the sink.
	Name pulumi.StringOutput `pulumi:"name"`
	// The fully qualified name of the sink.
	QualifiedSqlName pulumi.StringOutput `pulumi:"qualifiedSqlName"`
	// The identifier for the sink schema.
	SchemaName pulumi.StringPtrOutput `pulumi:"schemaName"`
	// The type of sink.
	SinkType pulumi.StringOutput `pulumi:"sinkType"`
	// The size of the sink.
	Size pulumi.StringOutput `pulumi:"size"`
	// Whether to emit the consolidated results of the query before the sink was created at the start of the sink.
	Snapshot pulumi.BoolPtrOutput `pulumi:"snapshot"`
	// The Kafka topic you want to subscribe to.
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewSinkKafka registers a new resource with the given unique name, arguments, and options.
func NewSinkKafka(ctx *pulumi.Context,
	name string, args *SinkKafkaArgs, opts ...pulumi.ResourceOption) (*SinkKafka, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.From == nil {
		return nil, errors.New("invalid value for required argument 'From'")
	}
	if args.KafkaConnection == nil {
		return nil, errors.New("invalid value for required argument 'KafkaConnection'")
	}
	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SinkKafka
	err := ctx.RegisterResource("materialize:index/sinkKafka:SinkKafka", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSinkKafka gets an existing SinkKafka resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSinkKafka(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SinkKafkaState, opts ...pulumi.ResourceOption) (*SinkKafka, error) {
	var resource SinkKafka
	err := ctx.ReadResource("materialize:index/sinkKafka:SinkKafka", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SinkKafka resources.
type sinkKafkaState struct {
	// The cluster to maintain this sink. If not specified, the size option must be specified.
	ClusterName *string `pulumi:"clusterName"`
	// The identifier for the sink database.
	DatabaseName *string `pulumi:"databaseName"`
	// How to interpret records (e.g. Debezium, Upsert).
	Envelope *SinkKafkaEnvelope `pulumi:"envelope"`
	// How to decode raw bytes from different formats into data structures it can understand at runtime.
	Format *SinkKafkaFormat `pulumi:"format"`
	// The name of the source, table or materialized view you want to send to the sink.
	From *SinkKafkaFrom `pulumi:"from"`
	// The name of the Kafka connection to use in the sink.
	KafkaConnection *SinkKafkaKafkaConnection `pulumi:"kafkaConnection"`
	// An optional list of columns to use for the Kafka key. If unspecified, the Kafka key is left unset.
	Keys []string `pulumi:"keys"`
	// The identifier for the sink.
	Name *string `pulumi:"name"`
	// The fully qualified name of the sink.
	QualifiedSqlName *string `pulumi:"qualifiedSqlName"`
	// The identifier for the sink schema.
	SchemaName *string `pulumi:"schemaName"`
	// The type of sink.
	SinkType *string `pulumi:"sinkType"`
	// The size of the sink.
	Size *string `pulumi:"size"`
	// Whether to emit the consolidated results of the query before the sink was created at the start of the sink.
	Snapshot *bool `pulumi:"snapshot"`
	// The Kafka topic you want to subscribe to.
	Topic *string `pulumi:"topic"`
}

type SinkKafkaState struct {
	// The cluster to maintain this sink. If not specified, the size option must be specified.
	ClusterName pulumi.StringPtrInput
	// The identifier for the sink database.
	DatabaseName pulumi.StringPtrInput
	// How to interpret records (e.g. Debezium, Upsert).
	Envelope SinkKafkaEnvelopePtrInput
	// How to decode raw bytes from different formats into data structures it can understand at runtime.
	Format SinkKafkaFormatPtrInput
	// The name of the source, table or materialized view you want to send to the sink.
	From SinkKafkaFromPtrInput
	// The name of the Kafka connection to use in the sink.
	KafkaConnection SinkKafkaKafkaConnectionPtrInput
	// An optional list of columns to use for the Kafka key. If unspecified, the Kafka key is left unset.
	Keys pulumi.StringArrayInput
	// The identifier for the sink.
	Name pulumi.StringPtrInput
	// The fully qualified name of the sink.
	QualifiedSqlName pulumi.StringPtrInput
	// The identifier for the sink schema.
	SchemaName pulumi.StringPtrInput
	// The type of sink.
	SinkType pulumi.StringPtrInput
	// The size of the sink.
	Size pulumi.StringPtrInput
	// Whether to emit the consolidated results of the query before the sink was created at the start of the sink.
	Snapshot pulumi.BoolPtrInput
	// The Kafka topic you want to subscribe to.
	Topic pulumi.StringPtrInput
}

func (SinkKafkaState) ElementType() reflect.Type {
	return reflect.TypeOf((*sinkKafkaState)(nil)).Elem()
}

type sinkKafkaArgs struct {
	// The cluster to maintain this sink. If not specified, the size option must be specified.
	ClusterName *string `pulumi:"clusterName"`
	// The identifier for the sink database.
	DatabaseName *string `pulumi:"databaseName"`
	// How to interpret records (e.g. Debezium, Upsert).
	Envelope *SinkKafkaEnvelope `pulumi:"envelope"`
	// How to decode raw bytes from different formats into data structures it can understand at runtime.
	Format *SinkKafkaFormat `pulumi:"format"`
	// The name of the source, table or materialized view you want to send to the sink.
	From SinkKafkaFrom `pulumi:"from"`
	// The name of the Kafka connection to use in the sink.
	KafkaConnection SinkKafkaKafkaConnection `pulumi:"kafkaConnection"`
	// An optional list of columns to use for the Kafka key. If unspecified, the Kafka key is left unset.
	Keys []string `pulumi:"keys"`
	// The identifier for the sink.
	Name *string `pulumi:"name"`
	// The identifier for the sink schema.
	SchemaName *string `pulumi:"schemaName"`
	// The size of the sink.
	Size *string `pulumi:"size"`
	// Whether to emit the consolidated results of the query before the sink was created at the start of the sink.
	Snapshot *bool `pulumi:"snapshot"`
	// The Kafka topic you want to subscribe to.
	Topic string `pulumi:"topic"`
}

// The set of arguments for constructing a SinkKafka resource.
type SinkKafkaArgs struct {
	// The cluster to maintain this sink. If not specified, the size option must be specified.
	ClusterName pulumi.StringPtrInput
	// The identifier for the sink database.
	DatabaseName pulumi.StringPtrInput
	// How to interpret records (e.g. Debezium, Upsert).
	Envelope SinkKafkaEnvelopePtrInput
	// How to decode raw bytes from different formats into data structures it can understand at runtime.
	Format SinkKafkaFormatPtrInput
	// The name of the source, table or materialized view you want to send to the sink.
	From SinkKafkaFromInput
	// The name of the Kafka connection to use in the sink.
	KafkaConnection SinkKafkaKafkaConnectionInput
	// An optional list of columns to use for the Kafka key. If unspecified, the Kafka key is left unset.
	Keys pulumi.StringArrayInput
	// The identifier for the sink.
	Name pulumi.StringPtrInput
	// The identifier for the sink schema.
	SchemaName pulumi.StringPtrInput
	// The size of the sink.
	Size pulumi.StringPtrInput
	// Whether to emit the consolidated results of the query before the sink was created at the start of the sink.
	Snapshot pulumi.BoolPtrInput
	// The Kafka topic you want to subscribe to.
	Topic pulumi.StringInput
}

func (SinkKafkaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sinkKafkaArgs)(nil)).Elem()
}

type SinkKafkaInput interface {
	pulumi.Input

	ToSinkKafkaOutput() SinkKafkaOutput
	ToSinkKafkaOutputWithContext(ctx context.Context) SinkKafkaOutput
}

func (*SinkKafka) ElementType() reflect.Type {
	return reflect.TypeOf((**SinkKafka)(nil)).Elem()
}

func (i *SinkKafka) ToSinkKafkaOutput() SinkKafkaOutput {
	return i.ToSinkKafkaOutputWithContext(context.Background())
}

func (i *SinkKafka) ToSinkKafkaOutputWithContext(ctx context.Context) SinkKafkaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkKafkaOutput)
}

// SinkKafkaArrayInput is an input type that accepts SinkKafkaArray and SinkKafkaArrayOutput values.
// You can construct a concrete instance of `SinkKafkaArrayInput` via:
//
//	SinkKafkaArray{ SinkKafkaArgs{...} }
type SinkKafkaArrayInput interface {
	pulumi.Input

	ToSinkKafkaArrayOutput() SinkKafkaArrayOutput
	ToSinkKafkaArrayOutputWithContext(context.Context) SinkKafkaArrayOutput
}

type SinkKafkaArray []SinkKafkaInput

func (SinkKafkaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SinkKafka)(nil)).Elem()
}

func (i SinkKafkaArray) ToSinkKafkaArrayOutput() SinkKafkaArrayOutput {
	return i.ToSinkKafkaArrayOutputWithContext(context.Background())
}

func (i SinkKafkaArray) ToSinkKafkaArrayOutputWithContext(ctx context.Context) SinkKafkaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkKafkaArrayOutput)
}

// SinkKafkaMapInput is an input type that accepts SinkKafkaMap and SinkKafkaMapOutput values.
// You can construct a concrete instance of `SinkKafkaMapInput` via:
//
//	SinkKafkaMap{ "key": SinkKafkaArgs{...} }
type SinkKafkaMapInput interface {
	pulumi.Input

	ToSinkKafkaMapOutput() SinkKafkaMapOutput
	ToSinkKafkaMapOutputWithContext(context.Context) SinkKafkaMapOutput
}

type SinkKafkaMap map[string]SinkKafkaInput

func (SinkKafkaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SinkKafka)(nil)).Elem()
}

func (i SinkKafkaMap) ToSinkKafkaMapOutput() SinkKafkaMapOutput {
	return i.ToSinkKafkaMapOutputWithContext(context.Background())
}

func (i SinkKafkaMap) ToSinkKafkaMapOutputWithContext(ctx context.Context) SinkKafkaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SinkKafkaMapOutput)
}

type SinkKafkaOutput struct{ *pulumi.OutputState }

func (SinkKafkaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SinkKafka)(nil)).Elem()
}

func (o SinkKafkaOutput) ToSinkKafkaOutput() SinkKafkaOutput {
	return o
}

func (o SinkKafkaOutput) ToSinkKafkaOutputWithContext(ctx context.Context) SinkKafkaOutput {
	return o
}

// The cluster to maintain this sink. If not specified, the size option must be specified.
func (o SinkKafkaOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// The identifier for the sink database.
func (o SinkKafkaOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// How to interpret records (e.g. Debezium, Upsert).
func (o SinkKafkaOutput) Envelope() SinkKafkaEnvelopePtrOutput {
	return o.ApplyT(func(v *SinkKafka) SinkKafkaEnvelopePtrOutput { return v.Envelope }).(SinkKafkaEnvelopePtrOutput)
}

// How to decode raw bytes from different formats into data structures it can understand at runtime.
func (o SinkKafkaOutput) Format() SinkKafkaFormatPtrOutput {
	return o.ApplyT(func(v *SinkKafka) SinkKafkaFormatPtrOutput { return v.Format }).(SinkKafkaFormatPtrOutput)
}

// The name of the source, table or materialized view you want to send to the sink.
func (o SinkKafkaOutput) From() SinkKafkaFromOutput {
	return o.ApplyT(func(v *SinkKafka) SinkKafkaFromOutput { return v.From }).(SinkKafkaFromOutput)
}

// The name of the Kafka connection to use in the sink.
func (o SinkKafkaOutput) KafkaConnection() SinkKafkaKafkaConnectionOutput {
	return o.ApplyT(func(v *SinkKafka) SinkKafkaKafkaConnectionOutput { return v.KafkaConnection }).(SinkKafkaKafkaConnectionOutput)
}

// An optional list of columns to use for the Kafka key. If unspecified, the Kafka key is left unset.
func (o SinkKafkaOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringArrayOutput { return v.Keys }).(pulumi.StringArrayOutput)
}

// The identifier for the sink.
func (o SinkKafkaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The fully qualified name of the sink.
func (o SinkKafkaOutput) QualifiedSqlName() pulumi.StringOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringOutput { return v.QualifiedSqlName }).(pulumi.StringOutput)
}

// The identifier for the sink schema.
func (o SinkKafkaOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringPtrOutput { return v.SchemaName }).(pulumi.StringPtrOutput)
}

// The type of sink.
func (o SinkKafkaOutput) SinkType() pulumi.StringOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringOutput { return v.SinkType }).(pulumi.StringOutput)
}

// The size of the sink.
func (o SinkKafkaOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// Whether to emit the consolidated results of the query before the sink was created at the start of the sink.
func (o SinkKafkaOutput) Snapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.BoolPtrOutput { return v.Snapshot }).(pulumi.BoolPtrOutput)
}

// The Kafka topic you want to subscribe to.
func (o SinkKafkaOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *SinkKafka) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

type SinkKafkaArrayOutput struct{ *pulumi.OutputState }

func (SinkKafkaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SinkKafka)(nil)).Elem()
}

func (o SinkKafkaArrayOutput) ToSinkKafkaArrayOutput() SinkKafkaArrayOutput {
	return o
}

func (o SinkKafkaArrayOutput) ToSinkKafkaArrayOutputWithContext(ctx context.Context) SinkKafkaArrayOutput {
	return o
}

func (o SinkKafkaArrayOutput) Index(i pulumi.IntInput) SinkKafkaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SinkKafka {
		return vs[0].([]*SinkKafka)[vs[1].(int)]
	}).(SinkKafkaOutput)
}

type SinkKafkaMapOutput struct{ *pulumi.OutputState }

func (SinkKafkaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SinkKafka)(nil)).Elem()
}

func (o SinkKafkaMapOutput) ToSinkKafkaMapOutput() SinkKafkaMapOutput {
	return o
}

func (o SinkKafkaMapOutput) ToSinkKafkaMapOutputWithContext(ctx context.Context) SinkKafkaMapOutput {
	return o
}

func (o SinkKafkaMapOutput) MapIndex(k pulumi.StringInput) SinkKafkaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SinkKafka {
		return vs[0].(map[string]*SinkKafka)[vs[1].(string)]
	}).(SinkKafkaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SinkKafkaInput)(nil)).Elem(), &SinkKafka{})
	pulumi.RegisterInputType(reflect.TypeOf((*SinkKafkaArrayInput)(nil)).Elem(), SinkKafkaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SinkKafkaMapInput)(nil)).Elem(), SinkKafkaMap{})
	pulumi.RegisterOutputType(SinkKafkaOutput{})
	pulumi.RegisterOutputType(SinkKafkaArrayOutput{})
	pulumi.RegisterOutputType(SinkKafkaMapOutput{})
}
