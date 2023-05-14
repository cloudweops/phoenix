package impl

import "github.com/CloudWeOps/phoenix/apps/host"

// 接口静态检测，编译器做的检查
// 看 HostServiceImpl 是否实现了接口 host.HostService
var _ host.HostService = (*HostServiceImpl)(nil)

// host 模块的功能都在 HostServiceImpl 这里定义
type HostServiceImpl struct {
	// log myzap.logger 1. zap自带全局logger，2. 子log实现，3. level 动态调整
}
