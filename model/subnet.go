package model

import (
	"strings"
	"fmt"
)

type Subnet struct {
	Identifier       `yaml:",inline"`
	AvailabilityZone string     `yaml:"availabilityZone,omitempty"`
	InstanceCIDR     string     `yaml:"instanceCIDR,omitempty"`
	NatGateway       NatGateway `yaml:"natGateway,omitempty"`
}

func (c Subnet) AvailabilityZoneLogicalName() string {
	return strings.Replace(strings.Title(c.AvailabilityZone), "-", "", -1)
}

func (c Subnet) NatGatewayLogicalName() string {
	return "NatGateway" + c.AvailabilityZoneLogicalName()
}

func (c Subnet) NatGatewayRef() string {
	return c.NatGateway.Identifier.Ref(c.NatGatewayLogicalName())
}

func (c Subnet) NatGatewayEIPAllocationLogicalName() string {
	return c.NatGatewayLogicalName() + "EIPAllocation"
}

func (c Subnet) NatGatewayEIPAllocationRef() string {
	if c.NatGateway.EIPAllocation.StackID != "" || c.NatGateway.EIPAllocation.ID != "" {
		return c.NatGateway.EIPAllocation.Identifier.Ref("")
	} else {
		return fmt.Sprintf(`{ "Fn::GetAtt": [%q, "AllocationId"] }`, c.NatGatewayEIPAllocationLogicalName())
	}
}

type PublicSubnet struct {
	Subnet            `yaml:",inline"`
	MapPublicIp       bool       `yaml:"mapPublicIP,omitempty"`
	PrivateRouteTable RouteTable `yaml:"privateRouteTable,omitempty"`
}

func (c PublicSubnet) LogicalName() string {
	return "Subnet" + c.AvailabilityZoneLogicalName()
}

func (c PublicSubnet) Ref() string {
	return c.Identifier.Ref(c.LogicalName())
}

func (c PublicSubnet) PrivateRouteTableLogicalName() string {
	return "PrivateRouteTable" + c.AvailabilityZoneLogicalName()
}

func (c PublicSubnet) PrivateRouteTableRef() string {
	return c.PrivateRouteTable.Identifier.Ref(c.PrivateRouteTableLogicalName())
}

type PrivateSubnet struct {
	Subnet     `yaml:",inline"`
	RouteTable RouteTable `yaml:"routeTable,omitempty"`
}

func (c PrivateSubnet) LogicalName(prefix string) string {
	return prefix + "PrivateSubnet" + c.AvailabilityZoneLogicalName()
}

func (c PrivateSubnet) Ref(prefix string) string {
	return c.Identifier.Ref(c.LogicalName(prefix))
}

func (c PrivateSubnet) RouteTableLogicalName() string {
	return "PrivateRouteTable" + c.AvailabilityZoneLogicalName()
}

func (c PrivateSubnet) RouteTableRef() string {
	return c.RouteTable.Identifier.Ref(c.RouteTableLogicalName())
}

func (c PrivateSubnet) PrivateRouteTableRef(publicSubnets []*PublicSubnet) string {
	for _, subnet := range publicSubnets {
		if subnet.AvailabilityZone == c.AvailabilityZone {
			return subnet.PrivateRouteTableRef()
		}
	}
	return ""
}

type NatGateway struct {
	Identifier `yaml:",inline"`
	EIPAllocation EIPAllocation `yaml:"eip,omitempty"`
}

type EIPAllocation struct {
	Identifier    `yaml:",inline"`
}
