package service

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/recovery-flow/organization-storage/internal/service/handlers"
	"github.com/recovery-flow/roles"
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
	adminGrant := service.TokenManager.RoleGrant(service.Config.JWT.AccessToken.SecretKey, string(roles.RoleUserAdmin))

	r.Route("/org-storage", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)
				r.Route("/organization", func(r chi.Router) {
					r.Post("/", handlers.OrganizationCreate)

					r.Route("/{organization_id}", func(r chi.Router) {
						r.Patch("/", handlers.OrganizationUpdate)
						r.Patch("/logo", handlers.OrganizationUploadLogo)
						r.Patch("/links", handlers.OrganizationLinksUpdate)

						r.Route("/participant", func(r chi.Router) {
							r.Post("/", handlers.ParticipantCreate)

							r.Route("/{user_id}", func(r chi.Router) {
								r.Patch("/", handlers.ParticipantUpdate)
							})
						})
					})
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/organization", func(r chi.Router) {
					r.Route("/{organization_id}", func(r chi.Router) {
						r.Get("/", handlers.OrganizationByID)

						r.Route("/participant", func(r chi.Router) {
							r.Get("/", handlers.ParticipantsByOrganization)
							r.Get("/{user_id}", handlers.ParticipantByUserID)
						})
					})
				})
			})

			r.Route("/admin", func(r chi.Router) {
				r.Use(adminGrant)
				r.Route("/organization", func(r chi.Router) {
					r.Route("/{organization_id}", func(r chi.Router) {
						r.Patch("/{value}", handlers.AdminVerifiedOrganization)

						r.Route("/participant", func(r chi.Router) {
							r.Route("/{user_id}", func(r chi.Router) {
								r.Patch("/{value}", handlers.AdminVerifiedParticipant)
							})
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
