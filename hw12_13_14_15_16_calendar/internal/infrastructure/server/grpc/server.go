package internalgrpc

import (
	"context"
	"errors"
	"net"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/api/pb"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cfg        config.ServerGrpcConf
	logger     logger.Logger
	grpcServer *grpc.Server
	pb.UnimplementedEventServiceServer
}

func (s *Server) CreateEvent(_ context.Context, _ *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}

func New(cfg config.ServerGrpcConf, logger logger.Logger) *Server {
	return &Server{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *Server) Start(_ context.Context) error {
	address := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	lsn, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(
	//grpc.ChainUnaryInterceptor(validate.UnaryServerRequestValidatorInterceptor(validate.ValidateReq)
	//),
	)

	pb.RegisterEventServiceServer(s.grpcServer, s)

	// Блокирующий вызов
	s.logger.Info("gRPC server starting at " + address)
	if err := s.grpcServer.Serve(lsn); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		s.logger.Error("gRPC server failed: " + err.Error())
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if s.grpcServer == nil {
		return nil
	}
	s.logger.Info("Shutting down gRPS server...")

	// Graceful shutdown с таймаутом
	done := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		s.logger.Info("gRPC server stopped gracefully")
		return nil
	case <-ctx.Done():
		s.logger.Error("Timeout during gRPC shutdown, forcing stop")
		s.grpcServer.Stop() // немедленное завершение
		return ctx.Err()
	}
}
