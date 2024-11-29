package helpercontext

import (
	"context"
	"database/sql"

	"job-application/models"
)

func GetTx(c context.Context) *sql.Tx {
	var ctx models.Ctx = "ctx"
	if tx, ok := c.Value(ctx).(*sql.Tx); ok {
		return tx
	}
	return nil
}

func GetValueRoleFromToken(c context.Context) string {
	var key models.Role = "role_user"
	if role, ok := c.Value(key).(string); ok {
		return role
	}
	return ""
}

func GetValueUserIdFromToken(c context.Context) int64 {
	var key models.ID = "userId"
	if userId, ok := c.Value(key).(int64); ok {
		return userId
	}
	return 0
}
