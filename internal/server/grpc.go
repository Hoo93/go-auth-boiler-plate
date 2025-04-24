package server

import (
	"auth-server-boiler-plate/api/v1"
	"auth-server-boiler-plate/internal/conf"
	"auth-server-boiler-plate/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/vetching-corporation/vetching-infra/sdk/go/telemetry"

	originGRPC "google.golang.org/grpc"
)

// gRPC Interceptor → Kratos 미들웨어로 변환
func InterceptorConvertor(interceptor originGRPC.UnaryServerInterceptor) middleware.Middleware {
	return func(next middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				// transport 정보가 없으면 interceptor 없이 직접 실행
				return next(ctx, req)
			}

			info := &originGRPC.UnaryServerInfo{
				FullMethod: tr.Operation(),
			}

			// gRPC Interceptor 형태로 wrapping
			return interceptor(ctx, req, info, func(ctx context.Context, req interface{}) (interface{}, error) {
				return next(ctx, req)
			})
		}
	}
}

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, user *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			InterceptorConvertor(telemetry.GRPCMiddleware()),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	v1.RegisterGreeterServer(srv, greeter)
	v1.RegisterUserServer(srv, user)
	return srv
}
