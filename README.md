# gRPC proto

Common gRPC proto files interaction in golang.

## Health Check

This is directly using `google.golang.org/grpc/health/grpc_health_v1` package with wrapper `connectRPC`.

```go
// import "github.com/worldline-go/grpc/health"

if err := health.Check(ctx, k.HTTP, awesomev1connect.AwesomeRouterServiceName); err != nil {
    return nil, fmt.Errorf("could not connect to awesome router: %w", err)
}
```
