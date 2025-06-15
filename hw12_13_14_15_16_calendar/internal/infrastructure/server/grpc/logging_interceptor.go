package internalgrpc

import (
	"context"
	"fmt"
	"time"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"google.golang.org/grpc"
)

func UnaryLoggingInterceptor(logg logger.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		duration := time.Since(start)

		logMessage := fmt.Sprintf(
			`gRPC [%s] Duration: %s Error: %v`,
			info.FullMethod,
			duration,
			err,
		)

		logg.LogToFile(logMessage)
		return resp, err
	}
}
