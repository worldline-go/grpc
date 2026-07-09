package health_test

import (
	"testing"

	healthv1 "github.com/worldline-go/grpc/health/v1"

	// Blank import of the grpc-go health SERVER mirrors the exact binary layout
	// from the original panic stack: its init pulls in grpc_health_v1 and
	// registers grpc.health.v1.* in the global protobuf registry.
	_ "google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"
)

// TestNoProtoRegistryConflict guards against the "proto: file
// \"grpc/health/v1/health.proto\" has a name conflict over
// grpc.health.v1.HealthListResponse" panic.
//
// The panic used to happen at package-init time whenever a binary linked BOTH
// this module's healthv1 package (which previously shipped its own generated
// *.pb.go registering grpc.health.v1.*) AND grpc-go's grpc_health_v1 (pulled in
// by any grpc health server). If a double registration were reintroduced, the
// protobuf global registry would panic before this test body ever runs, failing
// the package.
func TestNoProtoRegistryConflict(t *testing.T) {
	// Compile-time proof that our exported type is the *same* type as grpc-go's
	// original (a type alias), not a distinct message with a duplicate proto
	// full name. This only compiles when they share identity.
	var _ *grpc_health_v1.HealthCheckResponse = (*healthv1.HealthCheckResponse)(nil)
	var _ *grpc_health_v1.HealthListResponse = (*healthv1.HealthListResponse)(nil)

	// Exercise the aliased types at runtime.
	req := &healthv1.HealthCheckRequest{Service: "svc"}
	if got := proto.Clone(req).(*healthv1.HealthCheckRequest).GetService(); got != "svc" {
		t.Fatalf("unexpected service: %q", got)
	}

	if healthv1.HealthCheckResponse_SERVING != grpc_health_v1.HealthCheckResponse_SERVING {
		t.Fatal("SERVING enum value mismatch between healthv1 alias and grpc_health_v1")
	}

	// The connect bindings reference this descriptor; make sure it resolves to
	// the single registration owned by grpc-go.
	if healthv1.File_health_v1_health_proto == nil {
		t.Fatal("File_health_v1_health_proto descriptor is nil")
	}
	if name := string(healthv1.File_health_v1_health_proto.Package()); name != "grpc.health.v1" {
		t.Fatalf("unexpected proto package: %q", name)
	}
}
