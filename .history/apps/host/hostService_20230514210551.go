package host

// Host 模型定义 model
type Host struct {
	*Resource // 公共属性
	*Describe // 独有属性
}

type Resource struct{}

type Describe struct{}

// host app 接口定义
type Service interface {
	CreateHost()
	QueryHost()
	UpdateHost()
	DeleteHost()
}
