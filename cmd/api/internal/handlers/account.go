package handlers

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sankarvj/seedgo/internal/account"
	"github.com/sankarvj/seedgo/internal/platform/auth"
	"github.com/sankarvj/seedgo/internal/platform/web"
	"go.opencensus.io/trace"
)

// Account represents the Account API method handler set.
type Account struct {
	db            *sqlx.DB
	authenticator *auth.Authenticator
	// ADD OTHER STATE LIKE THE LOGGER AND CONFIG HERE.
}

// List returns all the existing users in the system.
func (a *Account) List(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	ctx, span := trace.StartSpan(ctx, "handlers.Account.List")
	defer span.End()

	claims, ok := ctx.Value(auth.Key).(auth.Claims)
	if !ok {
		return web.NewShutdownError("claims missing from context")
	}

	accounts, err := account.List(ctx, claims, a.db)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, accounts, http.StatusOK)
}
