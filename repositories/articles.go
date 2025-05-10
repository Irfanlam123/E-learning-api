package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type IArticlesRepository interface {
	Create(ctx context.Context, articles *models.Articles) error
	GetAll(ctx context.Context) ([]models.Articles, error)
}

type ArticlesRepository struct {
	db *gorm.DB
}

func NewArticlesRepository(db *gorm.DB) IArticlesRepository {
	return &ArticlesRepository{
		db: db,
	}
}

func (r *ArticlesRepository) Create(ctx context.Context, articles *models.Articles) error {
	return r.db.WithContext(ctx).Create(articles).Error
}
func (r *ArticlesRepository) GetAll(ctx context.Context) ([]models.Articles, error) {
	var articles []models.Articles
	err := r.db.WithContext(ctx).Find(&articles).Error
	return articles, err
}
