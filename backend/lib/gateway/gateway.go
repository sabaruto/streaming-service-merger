package gateway

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/sabaruto/streaming-service-merger/backend/lib/genproto/customer/v1"
)

func StartReverseProxy(ctx context.Context, opts []grpc.DialOption, authAddr string) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()
	err := customer.RegisterCustomerServiceHandlerFromEndpoint(ctx, mux, authAddr, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to register service: %v", err)
	}

	return mux, nil
}
