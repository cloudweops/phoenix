package host

// Host 模型定义 model
type Host struct{}

type Service interface {
	CreateHost()
	QueryHost()
	UpdateHost()
	DeleteHost()
}
