package impl

import "github.com/CloudWeOps/phoenix/apps/host"

// 检测 HostServiceImpl 是否实现了 接口 host.HostService
var _ host.HostService = (*HostServiceImpl)(nil)

type HostServiceImpl struct{}
