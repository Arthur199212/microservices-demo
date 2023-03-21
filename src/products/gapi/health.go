package gapi

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func (s *Server) Check(
	ctx context.Context,
	req *grpc_health_v1.HealthCheckRequest,
) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Server) Watch(
	req *grpc_health_v1.HealthCheckRequest,
	ws grpc_health_v1.Health_WatchServer,
) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
