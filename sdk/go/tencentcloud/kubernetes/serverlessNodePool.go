// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create serverless node pool of cluster.
//
// ## Example Usage
// ### Add serverless node pool to a cluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		availabilityZone := "ap-guangzhou-3"
// 		if param := cfg.Get("availabilityZone"); param != "" {
// 			availabilityZone = param
// 		}
// 		exampleClusterCidr := "10.31.0.0/16"
// 		if param := cfg.Get("exampleClusterCidr"); param != "" {
// 			exampleClusterCidr = param
// 		}
// 		vpc, err := Vpc.GetSubnets(ctx, &vpc.GetSubnetsArgs{
// 			IsDefault:        pulumi.BoolRef(true),
// 			AvailabilityZone: pulumi.StringRef(availabilityZone),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		vpcId := vpc.InstanceLists[0].VpcId
// 		subnetId := vpc.InstanceLists[0].SubnetId
// 		sg, err := Security.GetGroups(ctx, &security.GetGroupsArgs{
// 			Name: pulumi.StringRef("default"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		sgId := sg.SecurityGroups[0].SecurityGroupId
// 		exampleCluster, err := Kubernetes.NewCluster(ctx, "exampleCluster", &Kubernetes.ClusterArgs{
// 			VpcId:                pulumi.String(vpcId),
// 			ClusterCidr:          pulumi.String(exampleClusterCidr),
// 			ClusterMaxPodNum:     pulumi.Int(32),
// 			ClusterName:          pulumi.String("tf_example_cluster"),
// 			ClusterDesc:          pulumi.String("tf example cluster"),
// 			ClusterMaxServiceNum: pulumi.Int(32),
// 			ClusterVersion:       pulumi.String("1.18.4"),
// 			ClusterDeployType:    pulumi.String("MANAGED_CLUSTER"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Kubernetes.NewServerlessNodePool(ctx, "exampleServerlessNodePool", &Kubernetes.ServerlessNodePoolArgs{
// 			ClusterId: exampleCluster.ID(),
// 			ServerlessNodes: kubernetes.ServerlessNodePoolServerlessNodeArray{
// 				&kubernetes.ServerlessNodePoolServerlessNodeArgs{
// 					DisplayName: pulumi.String("tf_example_serverless_node1"),
// 					SubnetId:    pulumi.String(subnetId),
// 				},
// 				&kubernetes.ServerlessNodePoolServerlessNodeArgs{
// 					DisplayName: pulumi.String("tf_example_serverless_node2"),
// 					SubnetId:    pulumi.String(subnetId),
// 				},
// 			},
// 			SecurityGroupIds: pulumi.StringArray{
// 				pulumi.String(sgId),
// 			},
// 			Labels: pulumi.AnyMap{
// 				"label1": pulumi.Any("value1"),
// 				"label2": pulumi.Any("value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Adding taints to the virtual nodes under this node pool
//
// The pods without appropriate tolerations will not be scheduled on this node. Refer [taint-and-toleration](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) for more details.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Kubernetes.NewServerlessNodePool(ctx, "example", &Kubernetes.ServerlessNodePoolArgs{
// 			ClusterId: pulumi.Any(tencentcloud_kubernetes_cluster.Example.Id),
// 			ServerlessNodes: kubernetes.ServerlessNodePoolServerlessNodeArray{
// 				&kubernetes.ServerlessNodePoolServerlessNodeArgs{
// 					DisplayName: pulumi.String("tf_example_serverless_node1"),
// 					SubnetId:    pulumi.Any(local.Subnet_id),
// 				},
// 				&kubernetes.ServerlessNodePoolServerlessNodeArgs{
// 					DisplayName: pulumi.String("tf_example_serverless_node2"),
// 					SubnetId:    pulumi.Any(local.Subnet_id),
// 				},
// 			},
// 			SecurityGroupIds: pulumi.StringArray{
// 				pulumi.Any(local.Sg_id),
// 			},
// 			Labels: pulumi.AnyMap{
// 				"label1": pulumi.Any("value1"),
// 				"label2": pulumi.Any("value2"),
// 			},
// 			Taints: kubernetes.ServerlessNodePoolTaintArray{
// 				&kubernetes.ServerlessNodePoolTaintArgs{
// 					Key:    pulumi.String("key1"),
// 					Value:  pulumi.String("value1"),
// 					Effect: pulumi.String("NoSchedule"),
// 				},
// 				&kubernetes.ServerlessNodePoolTaintArgs{
// 					Key:    pulumi.String("key1"),
// 					Value:  pulumi.String("value1"),
// 					Effect: pulumi.String("NoExecute"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// serverless node pool can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool test cls-xxx#np-xxx
// ```
type ServerlessNodePool struct {
	pulumi.CustomResourceState

	// cluster id of serverless node pool.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// labels of serverless node.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// life state of serverless node pool.
	LifeState pulumi.StringOutput `pulumi:"lifeState"`
	// serverless node pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// security groups of serverless node pool.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// node list of serverless node pool.
	ServerlessNodes ServerlessNodePoolServerlessNodeArrayOutput `pulumi:"serverlessNodes"`
	// taints of serverless node.
	Taints ServerlessNodePoolTaintArrayOutput `pulumi:"taints"`
}

// NewServerlessNodePool registers a new resource with the given unique name, arguments, and options.
func NewServerlessNodePool(ctx *pulumi.Context,
	name string, args *ServerlessNodePoolArgs, opts ...pulumi.ResourceOption) (*ServerlessNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ServerlessNodes == nil {
		return nil, errors.New("invalid value for required argument 'ServerlessNodes'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServerlessNodePool
	err := ctx.RegisterResource("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerlessNodePool gets an existing ServerlessNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerlessNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerlessNodePoolState, opts ...pulumi.ResourceOption) (*ServerlessNodePool, error) {
	var resource ServerlessNodePool
	err := ctx.ReadResource("tencentcloud:Kubernetes/serverlessNodePool:ServerlessNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerlessNodePool resources.
type serverlessNodePoolState struct {
	// cluster id of serverless node pool.
	ClusterId *string `pulumi:"clusterId"`
	// labels of serverless node.
	Labels map[string]interface{} `pulumi:"labels"`
	// life state of serverless node pool.
	LifeState *string `pulumi:"lifeState"`
	// serverless node pool name.
	Name *string `pulumi:"name"`
	// security groups of serverless node pool.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// node list of serverless node pool.
	ServerlessNodes []ServerlessNodePoolServerlessNode `pulumi:"serverlessNodes"`
	// taints of serverless node.
	Taints []ServerlessNodePoolTaint `pulumi:"taints"`
}

type ServerlessNodePoolState struct {
	// cluster id of serverless node pool.
	ClusterId pulumi.StringPtrInput
	// labels of serverless node.
	Labels pulumi.MapInput
	// life state of serverless node pool.
	LifeState pulumi.StringPtrInput
	// serverless node pool name.
	Name pulumi.StringPtrInput
	// security groups of serverless node pool.
	SecurityGroupIds pulumi.StringArrayInput
	// node list of serverless node pool.
	ServerlessNodes ServerlessNodePoolServerlessNodeArrayInput
	// taints of serverless node.
	Taints ServerlessNodePoolTaintArrayInput
}

func (ServerlessNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessNodePoolState)(nil)).Elem()
}

type serverlessNodePoolArgs struct {
	// cluster id of serverless node pool.
	ClusterId string `pulumi:"clusterId"`
	// labels of serverless node.
	Labels map[string]interface{} `pulumi:"labels"`
	// serverless node pool name.
	Name *string `pulumi:"name"`
	// security groups of serverless node pool.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// node list of serverless node pool.
	ServerlessNodes []ServerlessNodePoolServerlessNode `pulumi:"serverlessNodes"`
	// taints of serverless node.
	Taints []ServerlessNodePoolTaint `pulumi:"taints"`
}

// The set of arguments for constructing a ServerlessNodePool resource.
type ServerlessNodePoolArgs struct {
	// cluster id of serverless node pool.
	ClusterId pulumi.StringInput
	// labels of serverless node.
	Labels pulumi.MapInput
	// serverless node pool name.
	Name pulumi.StringPtrInput
	// security groups of serverless node pool.
	SecurityGroupIds pulumi.StringArrayInput
	// node list of serverless node pool.
	ServerlessNodes ServerlessNodePoolServerlessNodeArrayInput
	// taints of serverless node.
	Taints ServerlessNodePoolTaintArrayInput
}

func (ServerlessNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverlessNodePoolArgs)(nil)).Elem()
}

type ServerlessNodePoolInput interface {
	pulumi.Input

	ToServerlessNodePoolOutput() ServerlessNodePoolOutput
	ToServerlessNodePoolOutputWithContext(ctx context.Context) ServerlessNodePoolOutput
}

func (*ServerlessNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessNodePool)(nil)).Elem()
}

func (i *ServerlessNodePool) ToServerlessNodePoolOutput() ServerlessNodePoolOutput {
	return i.ToServerlessNodePoolOutputWithContext(context.Background())
}

func (i *ServerlessNodePool) ToServerlessNodePoolOutputWithContext(ctx context.Context) ServerlessNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessNodePoolOutput)
}

// ServerlessNodePoolArrayInput is an input type that accepts ServerlessNodePoolArray and ServerlessNodePoolArrayOutput values.
// You can construct a concrete instance of `ServerlessNodePoolArrayInput` via:
//
//          ServerlessNodePoolArray{ ServerlessNodePoolArgs{...} }
type ServerlessNodePoolArrayInput interface {
	pulumi.Input

	ToServerlessNodePoolArrayOutput() ServerlessNodePoolArrayOutput
	ToServerlessNodePoolArrayOutputWithContext(context.Context) ServerlessNodePoolArrayOutput
}

type ServerlessNodePoolArray []ServerlessNodePoolInput

func (ServerlessNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessNodePool)(nil)).Elem()
}

func (i ServerlessNodePoolArray) ToServerlessNodePoolArrayOutput() ServerlessNodePoolArrayOutput {
	return i.ToServerlessNodePoolArrayOutputWithContext(context.Background())
}

func (i ServerlessNodePoolArray) ToServerlessNodePoolArrayOutputWithContext(ctx context.Context) ServerlessNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessNodePoolArrayOutput)
}

// ServerlessNodePoolMapInput is an input type that accepts ServerlessNodePoolMap and ServerlessNodePoolMapOutput values.
// You can construct a concrete instance of `ServerlessNodePoolMapInput` via:
//
//          ServerlessNodePoolMap{ "key": ServerlessNodePoolArgs{...} }
type ServerlessNodePoolMapInput interface {
	pulumi.Input

	ToServerlessNodePoolMapOutput() ServerlessNodePoolMapOutput
	ToServerlessNodePoolMapOutputWithContext(context.Context) ServerlessNodePoolMapOutput
}

type ServerlessNodePoolMap map[string]ServerlessNodePoolInput

func (ServerlessNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessNodePool)(nil)).Elem()
}

func (i ServerlessNodePoolMap) ToServerlessNodePoolMapOutput() ServerlessNodePoolMapOutput {
	return i.ToServerlessNodePoolMapOutputWithContext(context.Background())
}

func (i ServerlessNodePoolMap) ToServerlessNodePoolMapOutputWithContext(ctx context.Context) ServerlessNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessNodePoolMapOutput)
}

type ServerlessNodePoolOutput struct{ *pulumi.OutputState }

func (ServerlessNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessNodePool)(nil)).Elem()
}

func (o ServerlessNodePoolOutput) ToServerlessNodePoolOutput() ServerlessNodePoolOutput {
	return o
}

func (o ServerlessNodePoolOutput) ToServerlessNodePoolOutputWithContext(ctx context.Context) ServerlessNodePoolOutput {
	return o
}

// cluster id of serverless node pool.
func (o ServerlessNodePoolOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessNodePool) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// labels of serverless node.
func (o ServerlessNodePoolOutput) Labels() pulumi.MapOutput {
	return o.ApplyT(func(v *ServerlessNodePool) pulumi.MapOutput { return v.Labels }).(pulumi.MapOutput)
}

// life state of serverless node pool.
func (o ServerlessNodePoolOutput) LifeState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessNodePool) pulumi.StringOutput { return v.LifeState }).(pulumi.StringOutput)
}

// serverless node pool name.
func (o ServerlessNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerlessNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// security groups of serverless node pool.
func (o ServerlessNodePoolOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerlessNodePool) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// node list of serverless node pool.
func (o ServerlessNodePoolOutput) ServerlessNodes() ServerlessNodePoolServerlessNodeArrayOutput {
	return o.ApplyT(func(v *ServerlessNodePool) ServerlessNodePoolServerlessNodeArrayOutput { return v.ServerlessNodes }).(ServerlessNodePoolServerlessNodeArrayOutput)
}

// taints of serverless node.
func (o ServerlessNodePoolOutput) Taints() ServerlessNodePoolTaintArrayOutput {
	return o.ApplyT(func(v *ServerlessNodePool) ServerlessNodePoolTaintArrayOutput { return v.Taints }).(ServerlessNodePoolTaintArrayOutput)
}

type ServerlessNodePoolArrayOutput struct{ *pulumi.OutputState }

func (ServerlessNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerlessNodePool)(nil)).Elem()
}

func (o ServerlessNodePoolArrayOutput) ToServerlessNodePoolArrayOutput() ServerlessNodePoolArrayOutput {
	return o
}

func (o ServerlessNodePoolArrayOutput) ToServerlessNodePoolArrayOutputWithContext(ctx context.Context) ServerlessNodePoolArrayOutput {
	return o
}

func (o ServerlessNodePoolArrayOutput) Index(i pulumi.IntInput) ServerlessNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerlessNodePool {
		return vs[0].([]*ServerlessNodePool)[vs[1].(int)]
	}).(ServerlessNodePoolOutput)
}

type ServerlessNodePoolMapOutput struct{ *pulumi.OutputState }

func (ServerlessNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerlessNodePool)(nil)).Elem()
}

func (o ServerlessNodePoolMapOutput) ToServerlessNodePoolMapOutput() ServerlessNodePoolMapOutput {
	return o
}

func (o ServerlessNodePoolMapOutput) ToServerlessNodePoolMapOutputWithContext(ctx context.Context) ServerlessNodePoolMapOutput {
	return o
}

func (o ServerlessNodePoolMapOutput) MapIndex(k pulumi.StringInput) ServerlessNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerlessNodePool {
		return vs[0].(map[string]*ServerlessNodePool)[vs[1].(string)]
	}).(ServerlessNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessNodePoolInput)(nil)).Elem(), &ServerlessNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessNodePoolArrayInput)(nil)).Elem(), ServerlessNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerlessNodePoolMapInput)(nil)).Elem(), ServerlessNodePoolMap{})
	pulumi.RegisterOutputType(ServerlessNodePoolOutput{})
	pulumi.RegisterOutputType(ServerlessNodePoolArrayOutput{})
	pulumi.RegisterOutputType(ServerlessNodePoolMapOutput{})
}
