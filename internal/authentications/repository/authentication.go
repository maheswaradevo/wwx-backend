package repository

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/wwx-backend/internal/model"
	"github.com/maheswaradevo/wwx-backend/pkg/common/constants"
	"go.uber.org/zap"
)

type authRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewAuthRepository(db *sql.DB, logger *zap.Logger) *authRepository {
	return &authRepository{
		db:     db,
		logger: logger,
	}
}

func (a authRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	query := constants.CheckUsernameQuery
	rows := a.db.QueryRowContext(ctx, query, username)

	user := &model.User{}

	err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		a.logger.Error("[FindByUsername] failed to scan the data", zap.Error(err))
		return nil, err
	} else if err == sql.ErrNoRows {
		a.logger.Info("[FindByUsername] no data existed")
		return nil, constants.ErrNoUsernameExist
	}
	return user, nil
}
