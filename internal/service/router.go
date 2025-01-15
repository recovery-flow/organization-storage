package service

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.AuthMdl(service.Config.JWT.AccessToken.SecretKey)
	rateLimiter := httpkit.NewRateLimiter(service.Config.Rate.MaxRequests, service.Config.Rate.TimeWindow, service.Config.Rate.Expiration)

	r.Route("/organization-storage", func(r chi.Router) {

	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
