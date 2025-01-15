package events

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/organization-storage/internal/config"
	"github.com/sirupsen/logrus"
)

const (
	accountQ         = "account"
	accountCreateKey = "account.create"
)

func Listener(ctx context.Context) {
	_, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}
}
