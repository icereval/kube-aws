package model

type Controller struct {
	AutoScalingGroup `yaml:"autoScalingGroup,omitempty"`
	PrivateSubnets   []*Subnet `yaml:"privateSubnets,omitempty"`
}

func (c Controller) TopologyPrivate() bool {
	return len(c.PrivateSubnets) > 0
}
