package memo

import (
	"context"

	"github.com/geeeeorge/Go-diary-api/model"
)

// Repository defines memo repository contract
type Repository interface {
	InsertMemo(ctx context.Context, memo *model.Memo) error
	SelectMemoByID(ctx context.Context, id int) (*model.Memo, error)
	SelectAllMemo(ctx context.Context) ([]*model.Memo, error)
}
