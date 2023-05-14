package host

// Host 模型定义 model
type Host struct{}

// 主机接口定义
type Service interface {
	CreateHost()
	QueryHost()
	UpdateHost()
	DeleteHost()
}
