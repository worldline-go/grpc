// Copyright 2015 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package healthv1 re-exports the canonical grpc.health.v1 message types from
// google.golang.org/grpc/health/grpc_health_v1 via type aliases.
//
// This package intentionally does NOT generate its own *.pb.go. Doing so would
// register the "grpc.health.v1.*" proto messages a second time in the global
// protobuf registry and panic with a namespace conflict whenever a binary also
// links google.golang.org/grpc/health/grpc_health_v1 (e.g. a grpc-go health
// server). Aliasing shares the single registration owned by grpc-go while
// keeping this import path and its API backward compatible.
package healthv1

import grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"

// Message and enum types (aliases share identity with the grpc-go originals).
type (
	HealthCheckRequest                = grpc_health_v1.HealthCheckRequest
	HealthCheckResponse               = grpc_health_v1.HealthCheckResponse
	HealthCheckResponse_ServingStatus = grpc_health_v1.HealthCheckResponse_ServingStatus
	HealthListRequest                 = grpc_health_v1.HealthListRequest
	HealthListResponse                = grpc_health_v1.HealthListResponse
)

// ServingStatus enum values.
const (
	HealthCheckResponse_UNKNOWN         = grpc_health_v1.HealthCheckResponse_UNKNOWN
	HealthCheckResponse_SERVING         = grpc_health_v1.HealthCheckResponse_SERVING
	HealthCheckResponse_NOT_SERVING     = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	HealthCheckResponse_SERVICE_UNKNOWN = grpc_health_v1.HealthCheckResponse_SERVICE_UNKNOWN
)

// HealthCheckResponse_ServingStatus_name / _value maps.
var (
	HealthCheckResponse_ServingStatus_name  = grpc_health_v1.HealthCheckResponse_ServingStatus_name
	HealthCheckResponse_ServingStatus_value = grpc_health_v1.HealthCheckResponse_ServingStatus_value
)

// File_health_v1_health_proto is the file descriptor used by the generated
// connect bindings. It aliases the descriptor owned by grpc-go so both refer to
// the same underlying registration.
var File_health_v1_health_proto = grpc_health_v1.File_grpc_health_v1_health_proto
