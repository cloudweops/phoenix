package host

// Host 模型定义 model
type Host struct{}

// host app 接口定义
type Service interface {
	CreateHost()
	QueryHost()
	UpdateHost()
	DeleteHost()
}
