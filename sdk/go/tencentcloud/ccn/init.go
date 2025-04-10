// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ccn

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
	case "tencentcloud:Ccn/attachment:Attachment":
		r = &Attachment{}
	case "tencentcloud:Ccn/bandwidthLimit:BandwidthLimit":
		r = &BandwidthLimit{}
	case "tencentcloud:Ccn/instance:Instance":
		r = &Instance{}
	case "tencentcloud:Ccn/instancesAcceptAttach:InstancesAcceptAttach":
		r = &InstancesAcceptAttach{}
	case "tencentcloud:Ccn/instancesRejectAttach:InstancesRejectAttach":
		r = &InstancesRejectAttach{}
	case "tencentcloud:Ccn/instancesResetAttach:InstancesResetAttach":
		r = &InstancesResetAttach{}
	case "tencentcloud:Ccn/routes:Routes":
		r = &Routes{}
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
		"Ccn/attachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ccn/bandwidthLimit",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ccn/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ccn/instancesAcceptAttach",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ccn/instancesRejectAttach",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ccn/instancesResetAttach",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Ccn/routes",
		&module{version},
	)
}
