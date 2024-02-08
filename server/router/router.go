package router

import (
	. "webserver/server/controller"
	. "webserver/server/http"
	"webserver/utility"
)

type ControlledRoutes struct {
	route      Route
	controller Controller
}

type RouterType = func(req Request, rcs []ControlledRoutes) Response

func Router(req Request, rcs []ControlledRoutes) Response {
	index := utility.SliceIndexOf(rcs, func(rc ControlledRoutes) bool {
		return rc.route.Url == req.Url
	})
	if index == -1 {
		return "404"
	}
	rc := rcs[index]
	return rc.controller(req, createBaseResponse(req))
}

type AddControllerType = func(route Route, controller Controller, rcs []ControlledRoutes)

func AddController(route Route, controller Controller, rcs []ControlledRoutes) []ControlledRoutes {
	rcs = append(
		rcs,
		ControlledRoutes{
			route:      route,
			controller: controller,
		})
	return rcs
}

func createBaseResponse(req Request) Response {
	return "response"
}
