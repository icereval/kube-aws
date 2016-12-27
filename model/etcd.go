package model

type Etcd struct {
	PrivateSubnets   []*Subnet `yaml:"privateSubnets,omitempty"`
}

func (c Etcd) TopologyPrivate() bool {
	return len(c.PrivateSubnets) > 0
}
