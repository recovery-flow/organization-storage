package service

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.AuthMdl(service.Config.JWT.AccessToken.SecretKey)

	r.Route("/organization-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Route("/organization", func(r chi.Router) {
					r.Post("/", nil) // create org

					r.Route("/{organization_id}", func(r chi.Router) {
						r.Route("/update", func(r chi.Router) {
							r.Patch("/", nil) // update organization
							r.Patch("/photo", nil)
							r.Patch("/link", nil)
						})

						r.Route("/employee", func(r chi.Router) {
							r.Post("/add", nil) // add employee

							r.Route("/{employee_id}", func(r chi.Router) {
								r.Patch("/update", nil) // update employee
							})
						})
					})
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/organization", func(r chi.Router) {
					r.Route("/{organization_id}", func(r chi.Router) {
						r.Patch("/", nil) // info org

						r.Route("/employee", func(r chi.Router) {
							r.Patch("/", nil)              // info employees
							r.Route("/{employee_id}", nil) // info about one employee
						})
					})
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
