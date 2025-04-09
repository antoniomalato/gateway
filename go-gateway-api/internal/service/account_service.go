package service

import (
	"github.com/antoniomalato/gateway/gateway-api/internal/domain"
	"github.com/antoniomalato/gateway/gateway-api/internal/repository"
)

type AccountService struct {
	repository domain.AccountRepository
}
