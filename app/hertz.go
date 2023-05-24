package app

import (
	"fmt"
	"path"
	"strings"

	"github.com/cloudwego/hertz/pkg/route"
)

var (
	hertzApps = map[string]HertzApp{}
)

// HertzApp Hertz服务的实例
type HertzApp interface {
	InternalApp
	Registry(route.IRouter)
	Version() string
}

// RegistryHertzApp 服务实例注册
func RegistryHertzApp(app HertzApp) {
	// 已经注册的服务禁止再次注册
	_, ok := hertzApps[app.Name()]
	if ok {
		panic(fmt.Sprintf("gin app %s has registered", app.Name()))
	}

	hertzApps[app.Name()] = app
}

// LoadedHertzApp 查询加载成功的服务
func LoadedHertzApp() (apps []string) {
	for k := range hertzApps {
		apps = append(apps, k)
	}
	return
}

func GetHertzApp(name string) HertzApp {
	app, ok := hertzApps[name]
	if !ok {
		panic(fmt.Sprintf("http app %s not registered", name))
	}

	return app
}

// LoadHertzApp 装载所有的 hertz app
func LoadHertzApp(pathPrefix string, root route.IRouter) {
	for _, api := range hertzApps {
		if pathPrefix != "" && !strings.HasPrefix(pathPrefix, "/") {
			pathPrefix = "/" + pathPrefix
		}
		api.Registry(root.Group(path.Join(pathPrefix, api.Version(), api.Name()))) // 路由拼接并注册
	}
}
