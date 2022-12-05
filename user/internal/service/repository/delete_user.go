package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	query, args, _ := r.queryBuilder.
		Delete("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	result, err := r.writerDB.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
