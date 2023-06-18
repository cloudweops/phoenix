package ioc

import (
	"path"
	"strings"

	"github.com/cloudweops/phoenix/http/restful/accessor/form"
	"github.com/cloudweops/phoenix/http/restful/accessor/yaml"
	"github.com/cloudweops/phoenix/http/restful/accessor/yamlk8s"
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin"
)

const ApiNamespace = "apis"

// HTTPService Http服务的实例
type ApiObject interface {
	IocObject
	Version() string
}

func RegistryApiObject(obj ApiObject) {
	RegistryObjectWithNamespace(ApiNamespace, obj)
}

func GetApi(name string) IocObject {
	return GetObjectWithNamespace(ApiNamespace, name)
}

func ListApiObjectNames() (names []string) {
	return store.Namespace(ApiNamespace).ObjectNames()
}

type GinApiObject interface {
	ApiObject
	Registry(gin.IRouter)
}

func LoadGinApiObject(pathPrefix string, root gin.IRouter) {
	objects := store.Namespace(ApiNamespace)
	for _, obj := range objects.Items {
		api, ok := obj.(GinApiObject)
		if !ok {
			continue
		}
		if pathPrefix != "" && !strings.HasPrefix(pathPrefix, "/") {
			pathPrefix = "/" + pathPrefix
		}
		api.Registry(root.Group(path.Join(pathPrefix, api.Version(), api.Name())))
	}
}

type GoRestfulApiObject interface {
	ApiObject
	Registry(*restful.WebService)
}

func LoadGoRestfulApiObject(pathPrefix string, root *restful.Container) {
	objects := store.Namespace(ApiNamespace)
	for _, obj := range objects.Items {
		api, ok := obj.(GoRestfulApiObject)
		if !ok {
			continue
		}
		pathPrefix = strings.TrimSuffix(pathPrefix, "/")
		ws := new(restful.WebService)
		ws.Path(path.Join(pathPrefix, api.Version(), api.Name())).Consumes(restful.MIME_JSON).Consumes(restful.MIME_JSON, form.MIME_POST_FORM, form.MIME_MULTIPART_FORM, yaml.MIME_YAML, yamlk8s.MIME_YAML).Produces(restful.MIME_JSON, yaml.MIME_YAML, yamlk8s.MIME_YAML)
		api.Registry(ws)
		root.Add(ws)
	}
}
