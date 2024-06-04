package delivery

import (
	"awesomeProject/internal/user/service"
	"go.uber.org/zap"
)

type UserDelivery struct {
	service service.UserService
	logger  *zap.SugaredLogger
}

func New(service service.UserService, logger *zap.SugaredLogger) *UserDelivery {
	return &UserDelivery{
		service: service,
		logger:  logger,
	}
}
