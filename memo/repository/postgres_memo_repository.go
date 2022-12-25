package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/geeeeorge/Go-diary-api/model"
)

// PostgresMemoRepository implements memo repository
type PostgresMemoRepository struct {
	db *gorm.DB
}

// NewPostgresMemoRepository returns new instance of PostgresMemoRepository
func NewPostgresMemoRepository(db *gorm.DB) *PostgresMemoRepository {
	return &PostgresMemoRepository{db: db}
}

// InsertMemo inserts a new memo
func (r *PostgresMemoRepository) InsertMemo(ctx context.Context, memo *model.Memo) error {
	return r.db.WithContext(ctx).Create(memo).Error
}

// SelectMemoByID retrieves a memo by id
func (r *PostgresMemoRepository) SelectMemoByID(ctx context.Context, id int) (*model.Memo, error) {
	var memo model.Memo
	if err := r.db.WithContext(ctx).First(&memo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &memo, nil
}

// SelectAllMemo retrieves all memos
func (r *PostgresMemoRepository) SelectAllMemo(ctx context.Context) ([]*model.Memo, error) {
	var memos []*model.Memo
	if err := r.db.WithContext(ctx).Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}
