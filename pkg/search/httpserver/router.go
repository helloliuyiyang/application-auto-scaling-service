package httpserver

import (
	"net/http"

	"github.com/emicklei/go-restful"
)

// Route define four basic info of the northbound interface
// which is used to invoke web service.
type Route struct {
	Name      string
	Method    string
	Pattern   string
	RouteFunc restful.RouteFunction
}

// Register routes
func Register() {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON, restful.MIME_XML).Produces(restful.MIME_JSON, restful.MIME_XML)
	registerRoute(ws)
	restful.Add(ws)
}

func registerRoute(ws *restful.WebService) {
	for _, route := range routes {
		ws.Route(ws.Method(route.Method).Path(route.Pattern).To(route.RouteFunc).Operation(route.Name))
	}
}

// Routes is an array of Route
type Routes []Route

var routes = Routes{
	// 伸缩组Fleet
	Route{
		"CreateFleet",
		http.MethodPost,
		"/fleet",
		CreateFleet,
	},
	Route{
		"DeleteFleet",
		http.MethodDelete,
		"/fleet",
		CreateFleet,
	},

	// 伸缩组策略
	Route{
		"CreateAutoScalingPolicy",
		http.MethodPost,
		"/as_policy",
		CreateAutoScalingPolicy,
	},
	Route{
		"DeleteAutoScalingPolicy",
		http.MethodDelete,
		"/as_policy",
		DeleteAutoScalingPolicy,
	},

	// 告警webhook
	Route{
		"HandleAlerts",
		http.MethodPost,
		"/alerts",
		Alerts,
	},

	// mock监控指标
	Route{
		"MockCache",
		http.MethodPost,
		"/cache",
		MockCache,
	},
}
