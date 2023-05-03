package authentications

import (
	"context"

	"github.com/maheswaradevo/wwx-backend/internal/model"
)

type AuthRepository interface {
	FindByUsername(ctx context.Context, username string) (*model.User, error)
}
