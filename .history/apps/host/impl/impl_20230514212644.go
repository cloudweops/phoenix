package impl

import "github.com/CloudWeOps/phoenix/apps/host"

var _ host.HostService = (*HostServiceImpl)(nil)

type HostServiceImpl struct{}
