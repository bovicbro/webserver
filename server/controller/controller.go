package controller

import . "webserver/server/http"

type Controller = func(req Request, res Response) Response
