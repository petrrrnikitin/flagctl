package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authv1 "github.com/petrrrnikitin/flagctl/services/auth/gen"
)

type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer
}

func (h *AuthHandler) Register(_ context.Context, _ *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (h *AuthHandler) Login(_ context.Context, _ *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (h *AuthHandler) ValidateToken(_ context.Context, _ *authv1.ValidateTokenRequest) (*authv1.ValidateTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (h *AuthHandler) RefreshToken(_ context.Context, _ *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
