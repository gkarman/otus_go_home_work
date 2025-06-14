package internalgrpc

import (
	"context"
	"errors"
	"net"

	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/api/pb"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/application/requestdto"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/domain/logger"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cfg        config.ServerGrpcConf
	logger     logger.Logger
	app        application.Calendar
	grpcServer *grpc.Server
	pb.UnimplementedEventServiceServer
}

func (s *Server) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	requestDto := requestdto.CreateEvent{
		UserID:       req.GetUserId(),
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		TimeStart:    req.GetStartTime().AsTime(),
		TimeEnd:      req.GetEndTime().AsTime(),
		NotifyBefore: req.GetNotifyBefore().AsDuration(),
	}

	response, err := s.app.CreateEvent(ctx, requestDto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "some error")
	}

	return &pb.CreateEventResponse{
		Id: response.ID,
	}, nil
}

func (s *Server) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	requestDto := requestdto.DeleteEvent{
		ID:     req.GetId(),
		UserID: req.GetUserId(),
	}

	err := s.app.DeleteEvent(ctx, requestDto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "some error")
	}

	return &pb.DeleteEventResponse{
		Status:  true,
		Message: "done",
	}, nil
}

func (s *Server) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.UpdateEventResponse, error) {
	requestDto := requestdto.UpdateEvent{
		ID:           req.Event.GetId(),
		UserID:       req.Event.GetUserId(),
		Title:        req.Event.GetTitle(),
		Description:  req.Event.GetDescription(),
		TimeStart:    req.Event.GetTimeStart().AsTime(),
		TimeEnd:      req.Event.GetTimeEnd().AsTime(),
		NotifyBefore: req.Event.GetNotifyBefore().AsDuration(),
	}

	err := s.app.UpdateEvent(ctx, requestDto)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "some error")
	}

	return &pb.UpdateEventResponse{
		Status:  true,
		Message: "done",
	}, nil
}

func New(cfg config.ServerGrpcConf, logger logger.Logger, app application.Calendar) *Server {
	return &Server{
		logger: logger,
		cfg:    cfg,
		app:    app,
	}
}

func (s *Server) Start(_ context.Context) error {
	address := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	lsn, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			UnaryLoggingInterceptor(s.logger),
		),
	)

	pb.RegisterEventServiceServer(s.grpcServer, s)

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
