package api

import (
	"github.com/cloudweops/phoenix/app/health"
	"github.com/cloudweops/phoenix/http/response"
	"github.com/emicklei/go-restful/v3"
)

func (h *handler) Check(r *restful.Request, w *restful.Response) {
	req := health.NewHealthCheckRequest()
	resp, err := h.service.Check(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, NewHealth(resp))
}
