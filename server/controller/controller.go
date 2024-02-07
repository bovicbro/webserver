package controller

import . "webserver/http"

type Controller = func(req Request, res Response) Response
