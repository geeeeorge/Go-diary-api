package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/geeeeorge/Go-diary-api/memo"
	"github.com/geeeeorge/Go-diary-api/model"
)

// MemoUsecase implements memo repository
type MemoUsecase struct {
	repository memo.Repository
	validate   *validator.Validate
}

// NewMemoUsecase returns new instance of MemoUsecase
func NewMemoUsecase(repository memo.Repository) *MemoUsecase {
	return &MemoUsecase{
		repository: repository,
		validate:   validator.New(),
	}
}

// CreateMemo creates a new memo
func (u *MemoUsecase) CreateMemo(ctx context.Context, memo *model.Memo) error {
	if err := u.validate.Struct(memo); err != nil {
		return err
	}
	return u.repository.InsertMemo(ctx, memo)
}

// GetMemoByID gets memo by given id
func (u *MemoUsecase) GetMemoByID(ctx context.Context, id int) (*model.Memo, error) {
	return u.repository.SelectMemoByID(ctx, id)
}

// GetAllMemo gets all memos stored in repository
func (u *MemoUsecase) GetAllMemo(ctx context.Context) ([]*model.Memo, error) {
	return u.repository.SelectAllMemo(ctx)
}
