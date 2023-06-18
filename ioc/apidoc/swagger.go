package apidoc

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

// APIDocs creates a new restful.WebService and returns an instance of restfulspec.NewOpenAPIService.
//
// Parameters:
// - apiDocPath: string representing the API documentation path.
// - docs: restfulspec.PostBuildSwaggerObjectFunc to handle the Swagger object after it's built.
//
// Returns:
// - Pointer to a restful.WebService.
func APIDocs(apiDocPath string, docs restfulspec.PostBuildSwaggerObjectFunc) *restful.WebService {

	// Define configuration object for OpenAPI service
	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       apiDocPath,
		PostBuildSwaggerObjectHandler: docs,
		DefinitionNameHandler: func(name string) string {

			// Remove certain definitions that are not relevant to API documentation
			if name == "state" || name == "sizeCache" || name == "unknownFields" {
				return ""
			}
			return name
		},
	}

	// Return pointer to new OpenAPI service
	return restfulspec.NewOpenAPIService(config)
}
