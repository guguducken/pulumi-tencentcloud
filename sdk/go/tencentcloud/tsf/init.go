// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:Tsf/apiGroup:ApiGroup":
		r = &ApiGroup{}
	case "tencentcloud:Tsf/apiRateLimitRule:ApiRateLimitRule":
		r = &ApiRateLimitRule{}
	case "tencentcloud:Tsf/application:Application":
		r = &Application{}
	case "tencentcloud:Tsf/applicationConfig:ApplicationConfig":
		r = &ApplicationConfig{}
	case "tencentcloud:Tsf/applicationFileConfig:ApplicationFileConfig":
		r = &ApplicationFileConfig{}
	case "tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease":
		r = &ApplicationFileConfigRelease{}
	case "tencentcloud:Tsf/applicationPublicConfig:ApplicationPublicConfig":
		r = &ApplicationPublicConfig{}
	case "tencentcloud:Tsf/applicationPublicConfigRelease:ApplicationPublicConfigRelease":
		r = &ApplicationPublicConfigRelease{}
	case "tencentcloud:Tsf/applicationReleaseConfig:ApplicationReleaseConfig":
		r = &ApplicationReleaseConfig{}
	case "tencentcloud:Tsf/bindApiGroup:BindApiGroup":
		r = &BindApiGroup{}
	case "tencentcloud:Tsf/cluster:Cluster":
		r = &Cluster{}
	case "tencentcloud:Tsf/configTemplate:ConfigTemplate":
		r = &ConfigTemplate{}
	case "tencentcloud:Tsf/deployContainerGroup:DeployContainerGroup":
		r = &DeployContainerGroup{}
	case "tencentcloud:Tsf/deployVmGroup:DeployVmGroup":
		r = &DeployVmGroup{}
	case "tencentcloud:Tsf/enableUnitRule:EnableUnitRule":
		r = &EnableUnitRule{}
	case "tencentcloud:Tsf/group:Group":
		r = &Group{}
	case "tencentcloud:Tsf/instancesAttachment:InstancesAttachment":
		r = &InstancesAttachment{}
	case "tencentcloud:Tsf/lane:Lane":
		r = &Lane{}
	case "tencentcloud:Tsf/laneRule:LaneRule":
		r = &LaneRule{}
	case "tencentcloud:Tsf/microservice:Microservice":
		r = &Microservice{}
	case "tencentcloud:Tsf/namespace:Namespace":
		r = &Namespace{}
	case "tencentcloud:Tsf/operateContainerGroup:OperateContainerGroup":
		r = &OperateContainerGroup{}
	case "tencentcloud:Tsf/operateGroup:OperateGroup":
		r = &OperateGroup{}
	case "tencentcloud:Tsf/pathRewrite:PathRewrite":
		r = &PathRewrite{}
	case "tencentcloud:Tsf/releaseApiGroup:ReleaseApiGroup":
		r = &ReleaseApiGroup{}
	case "tencentcloud:Tsf/repository:Repository":
		r = &Repository{}
	case "tencentcloud:Tsf/task:Task":
		r = &Task{}
	case "tencentcloud:Tsf/unitNamespace:UnitNamespace":
		r = &UnitNamespace{}
	case "tencentcloud:Tsf/unitRule:UnitRule":
		r = &UnitRule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := tencentcloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/apiGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/apiRateLimitRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/application",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/applicationConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/applicationFileConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/applicationFileConfigRelease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/applicationPublicConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/applicationPublicConfigRelease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/applicationReleaseConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/bindApiGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/configTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/deployContainerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/deployVmGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/enableUnitRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/instancesAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/lane",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/laneRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/microservice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/namespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/operateContainerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/operateGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/pathRewrite",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/releaseApiGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/task",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/unitNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tsf/unitRule",
		&module{version},
	)
}
