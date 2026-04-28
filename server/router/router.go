package router

import (
	"strings"

	. "webserver/server/controller"
	. "webserver/server/http"
)

type ControlledRoutes struct {
	route      Route
	controller Controller
}

type RouterType = func(req Request, rcs []ControlledRoutes) Response

func Router(req Request, controlledRoutes []ControlledRoutes) Response {
	for _, controllerRoute := range controlledRoutes {
		if req.HttpMethod != controllerRoute.route.Method {
			continue
		}
		matched, params := matchRoute(string(controllerRoute.route.Url), string(req.Url))
		if matched {
			req.PathParams = params
			return controllerRoute.controller(req, CreateBaseResponse(req))
		}
	}
	return Response{Body: "404 Not found", Status: NOT_FOUND, Content: HTML}
}

func matchRoute(pattern string, target string) (bool, map[string]string) {
	if before, ok := strings.CutSuffix(pattern, "/*"); ok {
		prefix := before
		if prefix == "" || strings.HasPrefix(target, prefix) {
			return true, nil
		}
		return false, nil
	}

	if !strings.Contains(pattern, ":") {
		return pattern == target, nil
	}

	patternParts := strings.Split(pattern, "/")
	targetParts := strings.Split(target, "/")

	if len(patternParts) != len(targetParts) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range patternParts {
		if strings.HasPrefix(patternParts[i], ":") {
			paramName := patternParts[i][1:]
			params[paramName] = targetParts[i]
		} else if patternParts[i] != targetParts[i] {
			return false, nil
		}
	}

	return true, params
}

func AddController(route Route, controller Controller, rcs []ControlledRoutes) []ControlledRoutes {
	rcs = append(
		rcs,
		ControlledRoutes{
			route:      route,
			controller: controller,
		})
	return rcs
}
