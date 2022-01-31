// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wordle

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Wordle struct {
	pulumi.CustomResourceState

	Date   pulumi.StringOutput      `pulumi:"date"`
	Result pulumi.StringArrayOutput `pulumi:"result"`
	Word   pulumi.StringOutput      `pulumi:"word"`
}

// NewWordle registers a new resource with the given unique name, arguments, and options.
func NewWordle(ctx *pulumi.Context,
	name string, args *WordleArgs, opts ...pulumi.ResourceOption) (*Wordle, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Word == nil {
		return nil, errors.New("invalid value for required argument 'Word'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Wordle
	err := ctx.RegisterResource("wordle:index:Wordle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWordle gets an existing Wordle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWordle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WordleState, opts ...pulumi.ResourceOption) (*Wordle, error) {
	var resource Wordle
	err := ctx.ReadResource("wordle:index:Wordle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wordle resources.
type wordleState struct {
}

type WordleState struct {
}

func (WordleState) ElementType() reflect.Type {
	return reflect.TypeOf((*wordleState)(nil)).Elem()
}

type wordleArgs struct {
	Word string `pulumi:"word"`
}

// The set of arguments for constructing a Wordle resource.
type WordleArgs struct {
	Word pulumi.StringInput
}

func (WordleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wordleArgs)(nil)).Elem()
}

type WordleInput interface {
	pulumi.Input

	ToWordleOutput() WordleOutput
	ToWordleOutputWithContext(ctx context.Context) WordleOutput
}

func (*Wordle) ElementType() reflect.Type {
	return reflect.TypeOf((**Wordle)(nil)).Elem()
}

func (i *Wordle) ToWordleOutput() WordleOutput {
	return i.ToWordleOutputWithContext(context.Background())
}

func (i *Wordle) ToWordleOutputWithContext(ctx context.Context) WordleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WordleOutput)
}

// WordleArrayInput is an input type that accepts WordleArray and WordleArrayOutput values.
// You can construct a concrete instance of `WordleArrayInput` via:
//
//          WordleArray{ WordleArgs{...} }
type WordleArrayInput interface {
	pulumi.Input

	ToWordleArrayOutput() WordleArrayOutput
	ToWordleArrayOutputWithContext(context.Context) WordleArrayOutput
}

type WordleArray []WordleInput

func (WordleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wordle)(nil)).Elem()
}

func (i WordleArray) ToWordleArrayOutput() WordleArrayOutput {
	return i.ToWordleArrayOutputWithContext(context.Background())
}

func (i WordleArray) ToWordleArrayOutputWithContext(ctx context.Context) WordleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WordleArrayOutput)
}

// WordleMapInput is an input type that accepts WordleMap and WordleMapOutput values.
// You can construct a concrete instance of `WordleMapInput` via:
//
//          WordleMap{ "key": WordleArgs{...} }
type WordleMapInput interface {
	pulumi.Input

	ToWordleMapOutput() WordleMapOutput
	ToWordleMapOutputWithContext(context.Context) WordleMapOutput
}

type WordleMap map[string]WordleInput

func (WordleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wordle)(nil)).Elem()
}

func (i WordleMap) ToWordleMapOutput() WordleMapOutput {
	return i.ToWordleMapOutputWithContext(context.Background())
}

func (i WordleMap) ToWordleMapOutputWithContext(ctx context.Context) WordleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WordleMapOutput)
}

type WordleOutput struct{ *pulumi.OutputState }

func (WordleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wordle)(nil)).Elem()
}

func (o WordleOutput) ToWordleOutput() WordleOutput {
	return o
}

func (o WordleOutput) ToWordleOutputWithContext(ctx context.Context) WordleOutput {
	return o
}

type WordleArrayOutput struct{ *pulumi.OutputState }

func (WordleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wordle)(nil)).Elem()
}

func (o WordleArrayOutput) ToWordleArrayOutput() WordleArrayOutput {
	return o
}

func (o WordleArrayOutput) ToWordleArrayOutputWithContext(ctx context.Context) WordleArrayOutput {
	return o
}

func (o WordleArrayOutput) Index(i pulumi.IntInput) WordleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wordle {
		return vs[0].([]*Wordle)[vs[1].(int)]
	}).(WordleOutput)
}

type WordleMapOutput struct{ *pulumi.OutputState }

func (WordleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wordle)(nil)).Elem()
}

func (o WordleMapOutput) ToWordleMapOutput() WordleMapOutput {
	return o
}

func (o WordleMapOutput) ToWordleMapOutputWithContext(ctx context.Context) WordleMapOutput {
	return o
}

func (o WordleMapOutput) MapIndex(k pulumi.StringInput) WordleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wordle {
		return vs[0].(map[string]*Wordle)[vs[1].(string)]
	}).(WordleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WordleInput)(nil)).Elem(), &Wordle{})
	pulumi.RegisterInputType(reflect.TypeOf((*WordleArrayInput)(nil)).Elem(), WordleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WordleMapInput)(nil)).Elem(), WordleMap{})
	pulumi.RegisterOutputType(WordleOutput{})
	pulumi.RegisterOutputType(WordleArrayOutput{})
	pulumi.RegisterOutputType(WordleMapOutput{})
}
