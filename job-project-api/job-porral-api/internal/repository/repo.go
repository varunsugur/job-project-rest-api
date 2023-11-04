package repository

import (
	"context"
	"errors"
	"golang/internal/models"

	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

type UserRepo interface {
	CreatUser(ctx context.Context, userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string) (models.User, error)
	CreatCompany(ctx context.Context, data models.Company) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	ViewCompanyById(ctx context.Context, cid uint64) (models.Company, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		Db: db,
	}, nil
}
