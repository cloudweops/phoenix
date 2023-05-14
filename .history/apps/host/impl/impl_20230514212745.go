package impl

import "github.com/CloudWeOps/phoenix/apps/host"

// 接口静态检测，看 HostServiceImpl 是否实现了接口 host.HostService
var _ host.HostService = (*HostServiceImpl)(nil)

type HostServiceImpl struct{}
