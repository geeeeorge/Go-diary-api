package repository

import (
	"context"
	"github.com/geeeeorge/Go-diary-api/model"
	"gorm.io/gorm"
)

// PostgresMemoRepository implements memo repository
type PostgresMemoRepository struct {
	db *gorm.DB
}

// NewPostgresMemoRepository returns new instance of PostgresMemoRepository
func NewPostgresMemoRepository(db *gorm.DB) *PostgresMemoRepository {
	return &PostgresMemoRepository{db: db}
}

// CreateMemo creates a new memo
func (r *PostgresMemoRepository) CreateMemo(ctx context.Context, memo *model.Memo) error {
	return r.db.WithContext(ctx).Create(memo).Error
}

// GetMemoByID retrieves a memo by id
func (r *PostgresMemoRepository) GetMemoByID(ctx context.Context, id int) (*model.Memo, error) {
	var memo model.Memo
	if err := r.db.WithContext(ctx).First(&memo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &memo, nil
}

// GetAllMemo retrieves all memos
func (r *PostgresMemoRepository) GetAllMemo(ctx context.Context) ([]*model.Memo, error) {
	var memos []*model.Memo
	if err := r.db.WithContext(ctx).Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}
