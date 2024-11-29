package helpercontext

import (
	"context"
	"database/sql"

	"job-application/models"
)

func SetTx(c context.Context, tx *sql.Tx) context.Context {
	var ctx models.Ctx = "ctx"
	return context.WithValue(c, ctx, tx)
}
