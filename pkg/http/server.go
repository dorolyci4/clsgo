// Simple http web server support apis and static files both.
// GoFrame used.
package http

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	// "github.com/lovelacelee/clsgo/pkg/log"
)

type Request = ghttp.Request
type Meta = g.Meta

const (
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="API Doc"/>
  <title>APIDoc</title>
  <link rel="stylesheet" href="/css/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="/js/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api.json',
			dom_id: '#swagger-ui',
		});
	};
</script>
</body>
</html>
`
)

type DefaultHandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Self-defined Middleware
// Intermediate processing of the interface
type MiddlewareFn func(r *Request)
type Resource interface{}
type ResourceHandle struct {
	MiddlewareCallback MiddlewareFn
	// Pointer of Resource struct
	Res Resource
}

type APIS map[string]interface{}
type APIG map[string]ResourceHandle

func init() {
}

// Default Middleware
func MiddlewareDefault(r *ghttp.Request) {
	r.Middleware.Next()
	//https://goframe.org/pages/viewpage.action?pageId=1114281
	if r.Response.BufferLength() > 0 {
		return
	}

	res := r.GetHandlerResponse()

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    200,
		Message: "OK",
		Data:    res,
	})
}

// Create a http server, with default routes[/(public),/api,/swagger],
// apis is a map of api interfaces,
func App(host string, port int, apiv string, apis *APIS, apig *APIG) {
	sApi := g.Server("API")
	// When server: openapiPath: "/api.json" swaggerPath: "/swagger" not set,
	// you could enable openapi document as follow:

	openApi := sApi.GetOpenApi()
	openApi.Info.Title = "Reference"
	openApi.Info.Description = "API reference"
	openApi.Info.Version = apiv
	// sApi.SetOpenApiPath("/api.json")

	for k, v := range *apis {
		sApi.BindHandler(k, v)
	}
	sApi.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/swagger", func(r *ghttp.Request) {
			r.Response.Write(swaggerUIPageContent)
		})
	})
	for k, v := range *apig {
		sApi.Group(k, func(group *ghttp.RouterGroup) {
			group.Middleware(v.MiddlewareCallback)
			group.Bind(v.Res)
		})
	}
	sApi.SetServerRoot("public")

	sApi.SetAddr(host)
	sApi.SetPort(port)
	sApi.Start()

	g.Wait()
}
