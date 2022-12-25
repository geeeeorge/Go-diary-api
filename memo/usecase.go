package memo

import (
	"context"
	"github.com/geeeeorge/Go-diary-api/model"
)

// Usecase defines memo usecase contract
type Usecase interface {
	CreateMemo(ctx context.Context, memo *model.Memo) error
	GetMemoByID(ctx context.Context, id int) (*model.Memo, error)
	GetAllMemo(ctx context.Context) ([]*model.Memo, error)
}
