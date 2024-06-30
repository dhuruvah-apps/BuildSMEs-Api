package middleware

import (
	"github.com/dhuruvah-apps/BuildSMEs-Api/config"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/auth"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/session"
	"github.com/dhuruvah-apps/BuildSMEs-Api/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
