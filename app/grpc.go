package app

import (
	"fmt"

	"google.golang.org/grpc"
)

var grpcApps = map[string]GRPCApp{}

// GrpcApp GRPC服务的实例
type GRPCApp interface {
	InternalApp
	Registry(*grpc.Server)
}

// RegistryGRPCApp 服务实例注册
func RegistryGRPCApp(app GRPCApp) {
	// 已经注册的服务禁止再次注册
	_, ok := grpcApps[app.Name()]
	if ok {
		panic(fmt.Sprintf("grpc app %s has registered", app.Name()))
	}

	grpcApps[app.Name()] = app
}

// LoadedGrpcApp 查询加载成功的服务
func LoadedGrpcApp() (apps []string) {
	for k := range grpcApps {
		apps = append(apps, k)
	}
	return
}

func GetGRPCApp(name string) GRPCApp {
	app, ok := grpcApps[name]
	if !ok {
		panic(fmt.Sprintf("grpc app %s not registered", name))
	}

	return app
}

// LoadGrpcApp 加载所有的Grpc app
func LoadGrpcApp(server *grpc.Server) {
	for _, app := range grpcApps {
		app.Registry(server)
	}
}
