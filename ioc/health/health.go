package health

import (
	"github.com/cloudweops/phoenix/http/restful/response"
	"github.com/cloudweops/phoenix/logger"
	"github.com/cloudweops/phoenix/logger/zap"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct {
	service           healthgrpc.HealthServer
	log               logger.Logger
	HealthCheckerPath string
}

// NewHealthChecker returns a pointer to a HealthChecker instance, initialized with the
// provided service, default HealthCheckerPath, and a logger.
//
// service: a grpc HealthServer implementation.
// Returns a pointer to a HealthChecker instance.
func NewHealthChecker(service healthgrpc.HealthServer) *HealthChecker {
	return &HealthChecker{
		service:           service,
		HealthCheckerPath: "/healthz",
		log:               zap.L().Named("health_checker"),
	}
}

// NewDefaultHealthChecker returns a new instance of HealthChecker with a
// default health server.
//
// Returns a pointer to HealthChecker.
func NewDefaultHealthChecker() *HealthChecker {
	return NewHealthChecker(health.NewServer())
}

// WebService returns a pointer to a restful.WebService instance.
//
// This method receives no parameters.
// It returns a *restful.WebService instance.
func (h *HealthChecker) WebService() *restful.WebService {
	ws := new(restful.WebService)
	tags := []string{"健康检查"}

	ws.Route(ws.GET(h.HealthCheckerPath).To(h.Check).
		Doc("查询服务当前状态").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "success", HealthCheckResponse{}))
	return ws
}

// Check is a method of the HealthChecker struct that checks the health of a service.
//
// It takes in two parameters, a pointer to a restful.Request and a pointer to a restful.Response.
func (h *HealthChecker) Check(r *restful.Request, w *restful.Response) {
	req := NewHealthCheckRequest()
	resp, err := h.service.Check(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	err = w.WriteAsJson(NewHealthCheckResponse(resp))
	if err != nil {
		zap.L().Errorf("failed to write success response: %v", err)
	}
}

// NewHealthCheckRequest returns a new health check request.
func NewHealthCheckRequest() *healthgrpc.HealthCheckRequest {
	return &healthgrpc.HealthCheckRequest{}
}

// NewHealthCheckResponse creates a new HealthCheckResponse from a healthgrpc.HealthCheckResponse.
func NewHealthCheckResponse(h *healthgrpc.HealthCheckResponse) HealthCheckResponse {
	return HealthCheckResponse{
		Status: h.Status.String(),
	}
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}
