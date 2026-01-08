package grpc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"github.com/dormitory-life/auth/internal/constants"
	rmodel "github.com/dormitory-life/auth/internal/server/request_models"
	auth "github.com/dormitory-life/auth/internal/service"
	pb "github.com/dormitory-life/auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	pb.UnimplementedAuthProtoServiceServer
	authService auth.AuthServiceClient
	logger      *slog.Logger
	grpcServer  *grpc.Server
	port        string
}

type GRPCServerConfig struct {
	AuthService auth.AuthServiceClient
	Logger      *slog.Logger
	Port        string
}

func NewServer(cfg GRPCServerConfig) *GRPCServer {
	return &GRPCServer{
		authService: cfg.AuthService,
		logger:      cfg.Logger,
		port:        cfg.Port,
	}
}

func (s *GRPCServer) Start() error {
	s.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(s.loggingInterceptor),
	)

	pb.RegisterAuthProtoServiceServer(s.grpcServer, s)

	reflection.Register(s.grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.logger.Info("gRPC server starting", slog.String("address", ":"+s.port))

	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("gRPC server failed: %w", err)
	}

	return nil
}

func (s *GRPCServer) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
		s.logger.Info("gRPC server stopped")
	}
}

func (s *GRPCServer) loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	s.logger.Debug("gRPC request",
		slog.String("method", info.FullMethod),
		slog.Any("request", req))

	resp, err := handler(ctx, req)

	if err != nil {
		s.logger.Error("gRPC error",
			slog.String("method", info.FullMethod),
			slog.String("error", err.Error()))
	} else {
		s.logger.Debug("gRPC response",
			slog.String("method", info.FullMethod),
			slog.Any("response", resp))
	}

	return resp, err
}

func (s *GRPCServer) CheckAccess(
	ctx context.Context,
	req *pb.CheckAccessRequest,
) (*pb.CheckAccessResponse, error) {
	switch req.RoleRequired {
	case true:
		s.logger.Debug("gRPC CheckAccess for admin called",
			slog.String("user_id", req.GetUserId()),
			slog.String("dormitory_id", req.GetDormitoryId()),
			slog.Any("roleRequired", req.RoleRequired))
	case false:
		s.logger.Debug("gRPC CheckAccess for student called",
			slog.String("user_id", req.GetUserId()),
			slog.String("dormitory_id", req.GetDormitoryId()),
			slog.Any("roleRequired", req.RoleRequired))
	}

	res, err := s.authService.GetUserInfoById(ctx, &rmodel.GetUserByIdRequest{
		UserId: req.UserId,
	})
	if err != nil {
		if errors.Is(err, auth.ErrNotFound) {
			return &pb.CheckAccessResponse{
				Allowed:  false,
				Reason:   "",
				UserRole: auth.ErrNotFound.Error(),
			}, err
		}

		return &pb.CheckAccessResponse{
			Allowed:  false,
			Reason:   "",
			UserRole: err.Error(),
		}, err
	}

	result, err := s.checkFieldsEquality(req, res)

	return result, err
}

func (s *GRPCServer) checkFieldsEquality(
	req *pb.CheckAccessRequest,
	resp *rmodel.GetUserByIdResponse,
) (*pb.CheckAccessResponse, error) {
	s.logger.Debug("Checking equality: dormitory ids compare", slog.String("requested", req.DormitoryId), slog.String("target", resp.Info.DormitoryId))
	if req.DormitoryId != resp.Info.DormitoryId {
		return &pb.CheckAccessResponse{
			Allowed:  false,
			Reason:   "Permission for another dormitory is denied",
			UserRole: resp.Info.Role,
		}, nil
	}

	if req.RoleRequired && resp.Info.Role != constants.UserAdminRole {
		return &pb.CheckAccessResponse{
			Allowed:  false,
			Reason:   fmt.Sprintf("Permission for student is deniedUser role is '%s', required 'admin'", resp.Info.Role),
			UserRole: resp.Info.Role,
		}, nil
	}

	return &pb.CheckAccessResponse{
		Allowed:  true,
		Reason:   "Allowed",
		UserRole: resp.Info.Role,
	}, nil
}
