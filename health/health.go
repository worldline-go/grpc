package health

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/worldline-go/grpc/health/v1/healthv1connect"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Check checks the health of the service.
//   - If service not return SERVING status, it will return an error.
func Check(ctx context.Context, client connect.HTTPClient, serviceName string, opts ...Option) error {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}

	healthClient := healthv1connect.NewHealthClient(client, o.BaseURL, o.ClientOptions...)
	healthRequest := &healthpb.HealthCheckRequest{Service: serviceName}
	resp, err := healthClient.Check(ctx, connect.NewRequest(healthRequest))
	if err != nil {
		return err
	} else if resp.Msg.GetStatus() != healthpb.HealthCheckResponse_SERVING {
		return fmt.Errorf("agent health check failed: %s", resp.Msg.Status.String())
	}

	return nil
}
