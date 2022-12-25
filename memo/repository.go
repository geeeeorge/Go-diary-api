package memo

import (
	"context"
	"github.com/geeeeorge/Go-diary-api/model"
)

// Repository defines memo repository contract
type Repository interface {
	CreateMemo(ctx context.Context, memo *model.Memo) error
	GetMemoByID(ctx context.Context, id int) (*model.Memo, error)
	GetAllMemo(ctx context.Context) ([]*model.Memo, error)
}
